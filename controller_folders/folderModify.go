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

// For modify org request (from API)
type ModifyFolderRequest struct {
	UID             string `json:"uid" binding:"required"`
	OrganizationUID string `json:"org_uid"`
	DepartmentUID   string `json:"department_uid"`
	ModifyFields    int    `json:"modify_fields" binding:"required"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
}

// Modify Folder. Administrators group or org-admin user required
func ModifyFolder(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := controller.UserFromGinContextValidateAdministrators(ctx, c)
	if err != nil {
		return
	}

	var folder ModifyFolderRequest
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

	//UID is key field, and required to find user which we want to modify
	//ModifyFields is flags(int) which describes which fields need to be changed

	if folder.ModifyFields == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "modify_fields must be specified"})
		return
	}

	modifyFolder := database_folders.FolderInfo{
		UID:             folder.UID,
		OrganizationUID: folder.OrganizationUID,
		DepartmentUID:   folder.DepartmentUID,
		Name:            folder.Name,
		Description:     folder.Description,
	}

	err = database_folders.ModifyFolder(ctx, modifyFolder, folder.ModifyFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	resultData := gin.H{"result": "ok"}
	c.JSON(http.StatusOK, resultData)
}
