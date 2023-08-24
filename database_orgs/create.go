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

	if org.OwnerUID == "" {
		err = errors.New("org.OwnerUID not specified")
		return
	}

	//generate new UID if not specified
	if org.UID == "" {
		org.UID = core.GenerateUID()
	}

	result, insertErr := orgs.InsertOne(ctx, org)
	if insertErr != nil {
		err = fmt.Errorf("error inserting Organization: %s", insertErr.Error())
		return
	}

	org.ID = result.InsertedID.(primitive.ObjectID)

	return org, nil
}
