package controller_departments

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_departments"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Delete Department API method
func DeleteDepartment(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := controller.UserFromGinContextValidateAdministrators(ctx, c)
	if err != nil {
		return
	}

	var dep controller.UIDRequest
	err = c.BindJSON(&dep)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	validationErr := validate.Struct(dep)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	//UID is key field, and required to find department

	if dep.UID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uid must be specified"})
		return
	}
	err = database_departments.DeleteDepartment(ctx, dep.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	resultData := gin.H{"result": "ok"}
	c.JSON(http.StatusOK, resultData)
}
