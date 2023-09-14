package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database_groups"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/gin-gonic/gin"
)

// Helper - get UserInfo from gin.Context and check administrators group, also returns json if error
func UserFromGinContext(ctx context.Context, c *gin.Context) (userActing database_users.UserInfo, administratorsMember bool, err error) {

	administratorsMember = false
	claim, found := c.Get(auth.AuthUserClaimKey)
	if !found {
		msg := auth.AuthUserClaimKey + " not found"
		err = errors.New(msg)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	userClaim := claim.(*auth.UserClaim)

	userActing, err = database_users.FindUserAndCheckHash(ctx, userClaim.Email, userClaim.UserHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: func CheckAdministratorsGroup should return bool and err
	errAdmins := database_groups.CheckAdministratorsGroup(userActing.Groups)
	if errAdmins == nil {
		administratorsMember = true
	}

	return userActing, administratorsMember, err
}

// Get User from context, and validate Administrators group membership
func UserFromGinContextValidateAdministrators(ctx context.Context, c *gin.Context) (userActing database_users.UserInfo, err error) {

	userActing, admins, err := UserFromGinContext(ctx, c)
	if err != nil {
		return
	}

	if !admins {
		msg := "administrators membership required"
		err = errors.New(msg)
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	return userActing, err
}
