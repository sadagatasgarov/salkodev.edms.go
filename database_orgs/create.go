package database_orgs

import (
	"context"
	"errors"
	"fmt"

	"github.com/AndrewSalko/salkodev.edms.go/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Creates new Organization in Organizations collection
func CreateOrganization(ctx context.Context, org OrganizationInfo) (createdOrg OrganizationInfo, err error) {
	orgs := Organizations()

	if primitive.ObjectID.IsZero(org.ID) {
		org.ID = primitive.NewObjectID()
	}

	if org.Name == "" {
		err = errors.New("org.Name not specified")
		return
	}

	exists, err := OrganizationWithNameExists(ctx, org.Name)
	if err != nil {
		return
	}

	if exists {
		err = errors.New("organization with such name already exists")
		return
	}

	_, err = core.UIDFromStringWithArg(org.OwnerUID, "org.OwnerUID")
	if err != nil {
		return
	}

	//generate new UID if not specified
	if org.UID == "" {
		org.UID = core.GenerateUID()
	} else {
		_, err = core.UIDFromStringWithArg(org.UID, "org.UID")
		if err != nil {
			return
		}
	}

	result, insertErr := orgs.InsertOne(ctx, org)
	if insertErr != nil {
		err = fmt.Errorf("error inserting Organization: %s", insertErr.Error())
		return
	}

	org.ID = result.InsertedID.(primitive.ObjectID)

	return org, nil
}
