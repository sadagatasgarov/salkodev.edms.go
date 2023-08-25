package database_orgs

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// Check if Org with specified name exists
func OrganizationWithNameExists(ctx context.Context, orgName string) (exists bool, err error) {

	orgs := Organizations()

	filter := bson.M{OrganizationInfoFieldName: orgName}
	count, err := orgs.CountDocuments(ctx, filter)
	if err != nil {
		return
	}

	if count > 0 {
		exists = true
		return
	}

	return false, nil
}
