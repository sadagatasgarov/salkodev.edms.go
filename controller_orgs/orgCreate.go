package controller_orgs

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database_groups"
	"github.com/AndrewSalko/salkodev.edms.go/database_orgs"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// For create org request (from API), see full OrganizationInfo. Administrator can make this request and create Organization with UID and OwnderUID
type CreateOrganizationRequest struct {
	UID         string `json:"uid"` //uid not required, it will be generated
	Name        string `json:"name" binding:"required"`
	OwnerUID    string `json:"owner_uid"` //owner_uid can only be set by Admins, for everyone else it uses current user uid
	Description string `json:"description,omitempty"`
}

// Create new organization. Administrators group reqiured
func CreateOrganization(c *gin.Context) {

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

	//if in Admins group - can create any number of orgs, and specify any user as owner.
	//If not - can create only one organization, and owner must be empty (applied as current user)

	admins := false

	err = database_groups.CheckAdministratorsGroup(userActing.Groups)
	if err == nil {
		admins = true
	}

	var org CreateOrganizationRequest
	err = c.BindJSON(&org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	validationErr := validate.Struct(org)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if !admins {
		//Assume we not admins, so we can create only ONE organization (for self)
		org.OwnerUID = userActing.UID
		//check if current user alredy has Organization
		count, err := database_orgs.GetOrganizationCountByOwner(ctx, userActing.UID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"you can create only one organization": err.Error()})
			return
		}
	} else {

		if org.OwnerUID == "" { //admin can specify any user as owner
			org.OwnerUID = userActing.UID
		}
	}

	orgData := database_orgs.OrganizationInfo{
		UID:         org.UID,
		OwnerUID:    org.OwnerUID,
		Name:        org.Name,
		Description: org.Description,
	}

	createdOrg, err := database_orgs.CreateOrganization(ctx, orgData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resultData := gin.H{"uid": createdOrg.UID}

	c.JSON(http.StatusOK, resultData)
}
