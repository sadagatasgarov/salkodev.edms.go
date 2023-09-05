package database_departments

import (
	"context"
	"errors"
	"fmt"

	"github.com/AndrewSalko/salkodev.edms.go/core"
	"github.com/AndrewSalko/salkodev.edms.go/database_orgs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Creates new Department
func CreateDepartment(ctx context.Context, department DepartmentInfo) (createdDepartment DepartmentInfo, err error) {
	deps := Departments()

	_, err = database_orgs.FindOrganizationByUID(ctx, department.OrganizationUID)
	if err != nil {
		return
	}

	//check how many departments are created in the organization
	count, err := deps.CountDocuments(ctx, bson.M{DepartmentInfoFieldOrgUID: department.OrganizationUID})
	if err != nil {
		return
	}

	if count > DepartmentsMaxCount {
		err = errors.New("the maximum number of departments already created for this organization")
		return
	}

	//TODO: check department name (in current organization)

	if primitive.ObjectID.IsZero(department.ID) {
		department.ID = primitive.NewObjectID()
	}

	//if department uid unspecified, generate new uid
	if department.UID == "" {
		department.UID = core.GenerateUID()
	}

	result, insertErr := deps.InsertOne(ctx, department)
	if insertErr != nil {
		err = fmt.Errorf("error inserting Department: %s", insertErr.Error())
		return
	}

	department.ID = result.InsertedID.(primitive.ObjectID)

	return department, nil
}
