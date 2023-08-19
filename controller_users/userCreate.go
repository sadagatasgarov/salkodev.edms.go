package controller_users

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database_groups"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

// For create user request (from API), see full UserInfo. Administrator can make this request and create user with all params
type CreateUserRequest struct {
	Name            string   `json:"name" binding:"required"`
	Email           string   `json:"email" binding:"required"`
	Password        string   `json:"password" binding:"required"`
	AccountOptions  int      `json:"account_options" binding:"required"`
	OrganizationUID string   `json:"org_uid,omitempty"`
	EmailConfirmed  bool     `json:"email_confirmed"`
	Groups          []string `json:"groups"`
}

// Create new user. Administrators group reqiured
func CreateUser(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	claim, found := c.Get(auth.AuthUserClaimKey)
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": auth.AuthUserClaimKey + " not found"})
		return
	}

	userClaim := claim.(*auth.UserClaim)

	userActing, err := database_users.FindUserAndCheckHash(ctx, userClaim.Email, userClaim.UserHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database_groups.CheckAdministratorsGroup(userActing.Groups)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user CreateUserRequest
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	validationErr := validate.Struct(user)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	//перевести email до lower-case
	emailNormalized := strings.ToLower(user.Email)

	users := database_users.Users()

	count, err := users.CountDocuments(ctx, bson.M{"email": emailNormalized})

	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error detected while fetching the email"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user with email already exists"})
		return
	}

	passwordHashed := auth.HashPassword(user.Password)

	accOps := database_users.PurifyAccountOptions(user.AccountOptions)

	userInfo := database_users.UserInfo{
		Name:            user.Name,
		Email:           user.Email,
		Password:        passwordHashed,
		AccountOptions:  accOps,
		OrganizationUID: user.OrganizationUID,
		EmailConfirmed:  user.EmailConfirmed,
		Groups:          user.Groups}

	createdUser, err := database_users.CreateUser(ctx, userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	//return created user uid
	resultData := gin.H{"uid": createdUser.UID}

	c.JSON(http.StatusOK, resultData)
}
