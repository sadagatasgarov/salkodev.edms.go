package controller_folders

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_folders"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FolderPageResponse struct {
	Data       []database_folders.FolderInfo `json:"data"`
	Pagination controller.PaginationInfo     `json:"pagination"`
}

func GetFoldersPage(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var result FolderPageResponse

	folders := database_folders.Folders()

	totalRecords, totalPages, skipCount, page, perPage, err := controller.PaginationPrepare(c, ctx, folders)
	if err != nil {
		return //web response done in PaginationPrepare
	}

	result.Pagination.TotalRecords = totalRecords
	result.Pagination.TotalPages = totalPages

	// Find the documents with the skip and limit options
	cursor, err := folders.Find(ctx, bson.M{}, options.Find().SetSkip(int64(skipCount)).SetLimit(int64(perPage)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "folders.Find " + err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var foldersblock []database_folders.FolderInfo

	// Iterate over the cursor
	for cursor.Next(ctx) {
		var dep database_folders.FolderInfo
		err := cursor.Decode(&dep)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "dep.Find " + err.Error()})
			return
		}
		foldersblock = append(foldersblock, dep) //append to slice
	}

	result.Data = foldersblock
	result.Pagination.CurrentPage = page

	c.JSON(http.StatusOK, result)
}
