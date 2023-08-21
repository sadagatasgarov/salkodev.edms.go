package controller_users

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database_groups"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// For modify user request (from API), see full UserInfo. Administrator can make this request to modify user with all params
type ModifyUserRequest struct {
	UID             string `json:"uid" binding:"required"`
	ModifyFields    int    `json:"modify_fields" binding:"required"`
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	AccountOptions  int    `json:"account_options"`
	OrganizationUID string `json:"org_uid,omitempty"`
	EmailConfirmed  bool   `json:"email_confirmed"`
}

// Create new user. Administrators group reqiured
func ModifyUser(c *gin.Context) {

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

	var user ModifyUserRequest
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

	//user.UID is key field, and required to find user which we want to modify
	//user.ModifyFields is flags(int) which describes which fields need to be changed

	if user.ModifyFields == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "modify_fields must be specified"})
		return
	}

	modifyUser := database_users.UserInfo{
		UID:             user.UID,
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		AccountOptions:  user.AccountOptions,
		OrganizationUID: user.OrganizationUID,
		EmailConfirmed:  user.EmailConfirmed,
	}

	err = database_users.ModifyUser(ctx, modifyUser, user.ModifyFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	resultData := gin.H{"result": "ok"}
	c.JSON(http.StatusOK, resultData)
}
