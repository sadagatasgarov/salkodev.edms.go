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

// For create department request (from API), see full DepartmentInfo. Administrator can make this request and create Organization with UID and OwnderUID
type CreateDepartmentRequest struct {
	UID             string `json:"uid"`                        //uid not required, it will be generated
	OrganizationUID string `json:"org_uid" binding:"required"` //parent organization uid
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description,omitempty"`
}

// Create new department in organization. Administrators group reqiured
func CreateDepartment(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := controller.UserFromGinContextValidateAdministrators(ctx, c)
	if err != nil {
		return
	}

	var depReq CreateDepartmentRequest
	err = c.BindJSON(&depReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	validationErr := validate.Struct(depReq)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	depInfo := database_departments.DepartmentInfo{
		UID:             depReq.UID,
		OrganizationUID: depReq.OrganizationUID,
		Name:            depReq.Name,
		Description:     depReq.Description,
	}

	depCreated, err := database_departments.CreateDepartment(ctx, depInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultData := gin.H{"uid": depCreated.UID}

	c.JSON(http.StatusOK, resultData)
}
