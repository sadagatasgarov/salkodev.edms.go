package controller_users

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// URL param with special JWT for user registration (email) confirmation
const ConfirmUserRegistrationTokenParam = "token"

// User's email
const ConfirmUserEmail = "email"

func ConfirmRegistration(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	confirmToken := c.Query(ConfirmUserRegistrationTokenParam)
	if confirmToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "confirmation token not specified"})
		return
	}

	claims, errValidate := auth.ValidateTokenForRegistrationConfirmation(confirmToken)
	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidate.Error()})
		return
	}

	email := claims.Email

	//знайти користувача за мейлом
	user, findErr := database_users.FindUser(ctx, email)

	if findErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access denied, user not found"})
		return
	}

	users := database_users.Users()
	filter := bson.M{"_id": user.ID}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "email_confirmed", Value: true}}}}

	//оновити в базі стан користувача що email підтверджено
	_, err := users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
