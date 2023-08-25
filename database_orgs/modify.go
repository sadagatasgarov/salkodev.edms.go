package database_orgs

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// Modify organization. fields specify which property will be changed (see flags)
func ModifyOrganization(ctx context.Context, orgData OrganizationInfo, orgFields int) (err error) {

	//UID is required and key field
	if orgData.UID == "" {
		return errors.New("uid empty")
	}

	org, err := FindOrganizationByUID(ctx, orgData.UID)
	if err != nil {
		return
	}

	orgs := Organizations()
	upd := bson.D{}

	if orgFields&OrganizationInfoName > 0 {
		upd = append(upd, bson.E{Key: OrganizationInfoFieldName, Value: orgData.Name})
		org.Name = orgData.Name //TODO: check if org name exists!
	}

	if orgFields&OrganizationInfoDescription > 0 {
		upd = append(upd, bson.E{Key: OrganizationInfoFieldDescription, Value: orgData.Description})
		org.Description = orgData.Description
	}

	if orgFields&OrganizationInfoOwnerUID > 0 {
		//TODO: check new UID valid, check admins group or current user must be owner
		upd = append(upd, bson.E{Key: OrganizationInfoFieldOwnerUID, Value: orgData.OwnerUID})
		org.OwnerUID = orgData.OwnerUID
	}

	update := bson.D{{Key: "$set", Value: upd}}

	_, err = orgs.UpdateByID(ctx, org.ID, update)
	if err != nil {
		return
	}

	return nil
}
