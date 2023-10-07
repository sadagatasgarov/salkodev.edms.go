package controller_orgs

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_orgs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrganizationPageResponse struct {
	Data       []database_orgs.OrganizationInfo `json:"data"`
	Pagination controller.PaginationInfo        `json:"pagination"`
}

func GetOrganizationsPage(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var result OrganizationPageResponse

	orgs := database_orgs.Organizations()

	totalRecords, skipCount, page, perPage, err := controller.PaginationPrepare(c, ctx, orgs)
	if err != nil {
		return //web response done in PaginationPrepare
	}

	result.Pagination.TotalRecords = totalRecords

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
