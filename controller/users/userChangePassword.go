package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Change password request
type ChangePasswordRequest struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func ChangePassword(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	claim, found := c.Get(auth.AuthUserClaimKey)
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": auth.AuthUserClaimKey + " not found"})
		return
	}

	var changePasswordReq ChangePasswordRequest
	err := c.BindJSON(&changePasswordReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userClaim := claim.(*auth.UserClaim)

	//знайти користувача за email
	user, err := database.FindUser(ctx, userClaim.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: реалізувати валідацію та зміну пароля

	//1) перевірити чи введено коректно поточний пароль
	//2) чи не співпадає поточний пароль з тим що вказано як новий
	//3) чи відповідає політиці пароля

	currentPass := changePasswordReq.Password
	newPass := changePasswordReq.NewPassword

	if currentPass == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied, current password not specified"})
		return
	}

	err = auth.CheckPasswordPolicy(newPass)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//звіряємо поточний пароль (навіть якщо вкрадуть токен, це не допоможе змінити просто так)
	verifyResult := auth.VerifyPassword(currentPass, user.Password)
	if !verifyResult {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied, check current password"})
		return
	}

	passwordHashed := auth.HashPassword(newPass)

	//оновлюємо в базі користувача
	users := database.Users()
	filter := bson.M{"_id": user.ID}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: passwordHashed}}}}

	//оновити в базі стан користувача що email підтверджено
	_, err = users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateToken(user.Email, "") //TODO: зміна пароля - хеш користувача треба розрахувати та записати
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "result GenerateJwtToken: " + err.Error()})
		return
	}

	//повертаємо токен JWT
	c.JSON(http.StatusOK, gin.H{"token": token})

	c.JSON(http.StatusOK, gin.H{"result": "password changed successfully for user " + user.Name})
}
