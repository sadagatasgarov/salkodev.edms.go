package database_orgs

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// Delete organization. uid field is key
func DeleteOrganization(ctx context.Context, uid string) (err error) {

	if uid == "" {
		return errors.New("uid empty")
	}

	org, err := FindOrganizationByUID(ctx, uid)
	if err != nil {
		return
	}

	orgs := Organizations()

	filter := bson.M{"_id": org.ID}

	_, err = orgs.DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	return nil
}
