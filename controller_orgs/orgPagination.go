package controller_orgs

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_orgs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PerPageMin = 10
const PerPageMax = 500

const PageParam = "page"
const PerPage = "per_page"

type OrganizationPageResponse struct {
	Data       []database_orgs.OrganizationInfo `json:"data"`
	Pagination controller.PaginationInfo        `json:"pagination"`
}

func GetOrganizationsPage(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	page := 1

	var result OrganizationPageResponse

	pageNumberStr := c.Query(PageParam)
	if pageNumberStr != "" {
		pageVal, err := strconv.Atoi(pageNumberStr)
		if err == nil && pageVal > 0 {
			page = pageVal
		}
	}

	perPage := PerPageMin
	perPageStr := c.Query(PerPage)

	if perPageStr != "" {
		perPageVal, err := strconv.Atoi(perPageStr)
		if err == nil && perPageVal > 0 {
			perPage = perPageVal

			//check min-max allowed for per_page
			if perPage < PerPageMin || perPage > PerPageMax {
				perPage = PerPageMin
			}
		}
	}

	orgs := database_orgs.Organizations()

	filter := bson.M{}
	orgsCount, err := orgs.CountDocuments(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "orgs.CountDocuments " + err.Error()})
		return
	}

	result.Pagination.TotalRecords = orgsCount

	skipCount := (page - 1) * perPage

	// Find the documents with the skip and limit options
	cursor, err := orgs.Find(ctx, bson.M{}, options.Find().SetSkip(int64(skipCount)).SetLimit(int64(perPage)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "orgs.Find " + err.Error()})
		return
	}

	defer cursor.Close(ctx)

	var orgsBlock []database_orgs.OrganizationInfo

	// Iterate over the cursor
	for cursor.Next(ctx) {
		var org database_orgs.OrganizationInfo
		err := cursor.Decode(&org)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "orgs.Find " + err.Error()})
			return
		}
		orgsBlock = append(orgsBlock, org) //append to slice
	}

	result.Data = orgsBlock
	result.Pagination.CurrentPage = page

	c.JSON(http.StatusOK, result)
}
