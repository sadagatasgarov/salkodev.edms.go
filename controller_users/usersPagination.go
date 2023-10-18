package controller_users

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsersPageResponse struct {
	Data       []database_users.UserInfo `json:"data"`
	Pagination controller.PaginationInfo `json:"pagination"`
}

func GetUsersPage(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var result UsersPageResponse

	users := database_users.Users()

	totalRecords, totalPages, skipCount, page, perPage, err := controller.PaginationPrepare(c, ctx, users)
	if err != nil {
		return //web response done in PaginationPrepare
	}

	result.Pagination.TotalRecords = totalRecords
	result.Pagination.TotalPages = totalPages

	// Find the documents with the skip and limit options
	cursor, err := users.Find(ctx, bson.M{}, options.Find().SetSkip(int64(skipCount)).SetLimit(int64(perPage)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "users.Find " + err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var usersblock []database_users.UserInfo

	// Iterate over the cursor
	for cursor.Next(ctx) {
		var user database_users.UserInfo
		err := cursor.Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "users.Find " + err.Error()})
			return
		}
		usersblock = append(usersblock, user) //append to slice
	}

	result.Data = usersblock
	result.Pagination.CurrentPage = page

	c.JSON(http.StatusOK, result)
}
