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

// For create folder request (from API), see full FolderInfo. Administrator can make this request and create Organization with UID and OwnderUID
type CreateFolderRequest struct {
	UID             string `json:"uid"`                               //uid not required, it will be generated
	OrganizationUID string `json:"org_uid" binding:"required"`        //parent organization uid
	DepartmentUID   string `json:"department_uid" binding:"required"` //parent department uid
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description,omitempty"`
}

// Create new folder in department. Administrators group reqiured
func CreateFolder(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := controller.UserFromGinContextValidateAdministrators(ctx, c)
	if err != nil {
		return
	}

	var folderReq CreateFolderRequest
	err = c.BindJSON(&folderReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	validationErr := validate.Struct(folderReq)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	folderInfo := database_folders.FolderInfo{
		UID:             folderReq.UID,
		OrganizationUID: folderReq.OrganizationUID,
		DepartmentUID:   folderReq.DepartmentUID,
		Name:            folderReq.Name,
		Description:     folderReq.Description,
	}

	folderCreated, err := database_folders.CreateFolder(ctx, folderInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultData := gin.H{"uid": folderCreated.UID}

	c.JSON(http.StatusOK, resultData)
}
