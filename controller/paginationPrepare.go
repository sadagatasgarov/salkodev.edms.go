package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Min per_page value
const PaginationPerPageMin = 10

// Max per_page value
const PaginationPerPageMax = 500

const PaginationPageParam = "page"
const PaginationPerPageParam = "per_page"

// Helper function for pagination init process
func PaginationPrepare(c *gin.Context, ctx context.Context, collection *mongo.Collection) (totalRecords int64, totalPages int64, skipCount int64, page int64, perPage int, err error) {

	page = 1

	pageNumberStr := c.Query(PaginationPageParam)
	if pageNumberStr != "" {
		pageVal, err := strconv.Atoi(pageNumberStr)
		if err == nil && pageVal > 0 {
			page = int64(pageVal)
		}
	}

	perPage = PaginationPerPageMin
	perPageStr := c.Query(PaginationPerPageParam)

	if perPageStr != "" {
		perPageVal, err := strconv.Atoi(perPageStr)
		if err == nil && perPageVal > 0 {
			perPage = perPageVal

			//check min-max allowed for per_page
			if perPage < PaginationPerPageMin || perPage > PaginationPerPageMax {
				perPage = PaginationPerPageMin
			}
		}
	}

	filter := bson.M{}
	docsCount, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "collection.CountDocuments " + err.Error()})
		return
	}

	totalRecords = docsCount
	totalPages = docsCount / int64(perPage)
	remainder := docsCount % int64(perPage)
	if remainder > 0 {
		totalPages++
	}

	//check non-existend page
	if page > totalPages {
		errMsg := fmt.Sprintf("page %d not found", page)
		err = errors.New(errMsg)
		c.JSON(http.StatusNotFound, gin.H{"error": errMsg})
		return
	}

	skipCount = (page - 1) * int64(perPage)
	err = nil

	return
}
