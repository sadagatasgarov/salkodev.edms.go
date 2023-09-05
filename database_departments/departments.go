package database_departments

import (
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Field in Organization
const DepartmentsFieldName = "departments"

const DepartmentsCollectionName = "Departments"

// Max departments allowed for one organization
const DepartmentsMaxCount = 100

type DepartmentInfo struct {
	ID              primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UID             string             `bson:"uid" json:"uid" binding:"required"`
	OrganizationUID string             `bson:"org_uid" json:"org_uid" binding:"required"`
	Name            string             `bson:"name" json:"name" binding:"required"`
	Description     string             `bson:"description,omitempty" json:"description,omitempty"`
}

const DepartmentInfoFieldUID = "uid"
const DepartmentInfoFieldOrgUID = "org_uid"
const DepartmentInfoFieldName = "name"
const DepartmentInfoFieldDescription = "description"

// flag for modification Department org_uid
const DepartmentInfoOrganizationUID = 1

// flag for modification Department Name
const DepartmentInfoName = 2

// flag for modification Department Description
const DepartmentInfoDescription = 4

// Отримати колекцію Departments бази даних
func Departments() *mongo.Collection {

	collection := database.DataBase().Collection(DepartmentsCollectionName)
	return collection
}
