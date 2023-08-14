package controller_users

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database_groups"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/gin-gonic/gin"
)

// Remove user from groups (argument - UserGroupsModifyRequest)
func RemoveFromGroup(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	claim, found := c.Get(auth.AuthUserClaimKey)
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": auth.AuthUserClaimKey + " not found"})
		return
	}

	var removeReq UserGroupsModifyRequest
	err := c.BindJSON(&removeReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userClaim := claim.(*auth.UserClaim)

	user, err := database_users.FindUserAndCheckHash(ctx, userClaim.Email, userClaim.UserHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database_groups.RemoveUser(ctx, user, removeReq.UserOID, removeReq.GroupsUniqueNames)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
