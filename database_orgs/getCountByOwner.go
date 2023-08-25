package database_orgs

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// Get number of Organizations where specified user is owner
func GetOrganizationCountByOwner(ctx context.Context, userOwnerUID string) (count int64, err error) {

	orgs := Organizations()

	filter := bson.M{OrganizationInfoFieldOwnerUID: userOwnerUID}
	count, err = orgs.CountDocuments(ctx, filter)

	return
}
