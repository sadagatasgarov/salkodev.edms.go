package controller_departments

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database_departments"
	"github.com/AndrewSalko/salkodev.edms.go/database_groups"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// For modify org request (from API)
type ModifyDepartmentRequest struct {
	UID             string `json:"uid" binding:"required"`
	OrganizationUID string `json:"org_uid"`
	ModifyFields    int    `json:"modify_fields" binding:"required"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
}

// Modify department. Administrators group or org-admin user required
func ModifyDepartment(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	claim, found := c.Get(auth.AuthUserClaimKey)
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": auth.AuthUserClaimKey + " not found"})
		return
	}

	userClaim := claim.(*auth.UserClaim)

	userActing, err := database_users.FindUserAndCheckHash(ctx, userClaim.Email, userClaim.UserHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database_groups.CheckAdministratorsGroup(userActing.Groups)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dep ModifyDepartmentRequest
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

	//UID is key field, and required to find user which we want to modify
	//ModifyFields is flags(int) which describes which fields need to be changed

	if dep.ModifyFields == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "modify_fields must be specified"})
		return
	}

	modifyDep := database_departments.DepartmentInfo{
		UID:             dep.UID,
		OrganizationUID: dep.OrganizationUID,
		Name:            dep.Name,
		Description:     dep.Description,
	}

	err = database_departments.ModifyDepartment(ctx, modifyDep, dep.ModifyFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	resultData := gin.H{"result": "ok"}
	c.JSON(http.StatusOK, resultData)
}
