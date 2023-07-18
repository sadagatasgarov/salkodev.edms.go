package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second) //ctx
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

	//TODO: вернуть токен JWT
	c.JSON(http.StatusOK, gin.H{"result": "OK"})

}
