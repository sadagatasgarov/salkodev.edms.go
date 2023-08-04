package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var loginReq UserLoginRequest
	err := c.BindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	validationErr := validate.Struct(loginReq)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	//знайти користувача в базі (логін - мейл)
	user, err := database.FindUser(ctx, loginReq.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied, check login and password (1)"})
		return
	}

	//сверим пароль который мы получили в запросе с хешированным в базе...
	//в структуре resultUser.Password уже хеш пароля, а loginReq.Password - открытый пароль
	verifyResult := auth.VerifyPassword(loginReq.Password, user.Password)

	if !verifyResult {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied, check login and password (2)"})
		return
	}

	//перевірити, чи не блоковано облік.запис - тут пароль вже перевірено, тому
	//атакер не знатиме стан облік.запису без пароля
	if user.AccountOptions&database.UserAccountOptionDisabled > 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account disabled"})
		return
	}

	token, err := auth.GenerateToken(user.Email, user.Hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "result GenerateJwtToken: " + err.Error()})
		return
	}

	//повертаємо токен JWT
	c.JSON(http.StatusOK, gin.H{"token": token})
}
