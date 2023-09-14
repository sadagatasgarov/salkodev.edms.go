package controller_folders

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_folders"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Delete Folder API method
func DeleteFolder(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := controller.UserFromGinContextValidateAdministrators(ctx, c)
	if err != nil {
		return
	}

	var folder controller.UIDRequest
	err = c.BindJSON(&folder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	validationErr := validate.Struct(folder)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	//UID is key field, and required to find department

	if folder.UID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uid must be specified"})
		return
	}
	err = database_folders.DeleteFolder(ctx, folder.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	resultData := gin.H{"result": "ok"}
	c.JSON(http.StatusOK, resultData)
}
