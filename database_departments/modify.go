package database_departments

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// Modify department. fields specify which property will be changed (see flags)
func ModifyDepartment(ctx context.Context, depData DepartmentInfo, depFields int) (err error) {

	//UID is required and key field
	if depData.UID == "" {
		return errors.New("uid empty")
	}

	dep, err := FindDepartmentByUID(ctx, depData.UID)
	if err != nil {
		return
	}

	deps := Departments()
	upd := bson.D{}

	if depFields&DepartmentInfoOrganizationUID > 0 {
		upd = append(upd, bson.E{Key: DepartmentInfoFieldOrgUID, Value: depData.OrganizationUID})
		dep.OrganizationUID = depData.OrganizationUID
	}

	if depFields&DepartmentInfoName > 0 {
		upd = append(upd, bson.E{Key: DepartmentInfoFieldName, Value: depData.Name})
		dep.Name = depData.Name //TODO: check if department name exists (in one org)
	}

	if depFields&DepartmentInfoDescription > 0 {
		upd = append(upd, bson.E{Key: DepartmentInfoFieldDescription, Value: depData.Description})
		dep.Description = depData.Description
	}

	update := bson.D{{Key: "$set", Value: upd}}

	_, err = deps.UpdateByID(ctx, dep.ID, update)
	if err != nil {
		return
	}

	return nil
}
