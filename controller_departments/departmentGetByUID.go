package controller_departments

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/database_departments"
	"github.com/gin-gonic/gin"
)

// URL part contains uid (last part of get url)
const UIDParam = "uid"

func GetDepartmentByUID(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	uid := c.Param(UIDParam)
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "department uid not specified"})
		return
	}

	dep, err := database_departments.FindDepartmentByUID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dep)
}
