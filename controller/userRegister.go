package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"github.com/AndrewSalko/salkodev.edms.go/email"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRegistrationRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user UserRegistrationRequest
	err := c.BindJSON(&user)
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

	users := database.Users()

	count, err := users.CountDocuments(ctx, bson.M{"email": emailNormalized})

	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error detected while fetching the email"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User with email already exists"})
		return
	}

	passwordHashed := auth.HashPassword(user.Password)
	user.Password = passwordHashed

	resultInsertionNumber, insertErr := users.InsertOne(ctx, user)
	if insertErr != nil {
		msg := fmt.Sprintf("Error inserting User: %s", insertErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	userIDStr := resultInsertionNumber.InsertedID.(primitive.ObjectID).Hex()
	//потрібно надіслати Email з особливим посиланням-підтвердженням (токен для підтвердження)

	emailConfirmToken := auth.GenerateEmailConfirmationToken(userIDStr, emailNormalized)

	//TODO: зробити шаблон email для підтв.реєстрації
	emailBody := "Click on link to finish registration. Code: " + emailConfirmToken
	email.SendMail(emailNormalized, "SalkoDev EDMS registration", emailBody)

	jwtToken, err := auth.GenerateToken(emailNormalized)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resultData := gin.H{
		"userID":             userIDStr,
		"token":              jwtToken,
		"confirmation_token": emailConfirmToken}

	c.JSON(http.StatusOK, resultData)
}
