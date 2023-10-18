package controller_departments

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_departments"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DepartmentPageResponse struct {
	Data       []database_departments.DepartmentInfo `json:"data"`
	Pagination controller.PaginationInfo             `json:"pagination"`
}

func GetDepartmentsPage(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var result DepartmentPageResponse

	deps := database_departments.Departments()

	totalRecords, totalPages, skipCount, page, perPage, err := controller.PaginationPrepare(c, ctx, deps)
	if err != nil {
		return //web response done in PaginationPrepare
	}

	result.Pagination.TotalRecords = totalRecords
	result.Pagination.TotalPages = totalPages

	// Find the documents with the skip and limit options
	cursor, err := deps.Find(ctx, bson.M{}, options.Find().SetSkip(int64(skipCount)).SetLimit(int64(perPage)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "deps.Find " + err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var depsblock []database_departments.DepartmentInfo

	// Iterate over the cursor
	for cursor.Next(ctx) {
		var dep database_departments.DepartmentInfo
		err := cursor.Decode(&dep)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "dep.Find " + err.Error()})
			return
		}
		depsblock = append(depsblock, dep) //append to slice
	}

	result.Data = depsblock
	result.Pagination.CurrentPage = page

	c.JSON(http.StatusOK, result)
}
