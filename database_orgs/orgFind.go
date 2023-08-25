package database_orgs

import (
	"context"

	"github.com/AndrewSalko/salkodev.edms.go/core"
	"go.mongodb.org/mongo-driver/bson"
)

// Find Organization by uid
func FindOrganizationByUID(ctx context.Context, uid string) (org OrganizationInfo, err error) {

	_, err = core.UIDFromString(uid)
	if err != nil {
		return
	}

	orgs := Organizations()

	filter := bson.M{"uid": uid}
	err = orgs.FindOne(ctx, filter).Decode(&org)

	return org, err
}
