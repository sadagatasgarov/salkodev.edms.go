package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ChangePasswordRequest struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func ChangePassword(c *gin.Context) {

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second) //ctx
	defer cancel()

	var changePasswordReq ChangePasswordRequest
	err := c.BindJSON(&changePasswordReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate := validator.New()
	// validationErr := validate.Struct(loginReq)

	// if validationErr != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
	// 	return
	// }

	//TODO: реализовать смену пароля
	c.JSON(http.StatusOK, gin.H{"result": "OK"})

}
