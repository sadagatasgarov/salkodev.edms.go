package database_orgs

import (
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const OrganizationsCollectionName = "Organizations"

type OrganizationInfo struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UID         string             `bson:"uid" json:"uid" binding:"required"`
	OwnerUID    string             `bson:"owner_uid" json:"owner_uid" binding:"required"`
	Name        string             `bson:"name" json:"name" binding:"required"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
}

const OrganizationInfoFieldUID = "uid"
const OrganizationInfoFieldOwnerUID = "owner_uid"
const OrganizationInfoFieldName = "name"
const OrganizationInfoFieldDescription = "description"

// flag for modification Org Name
const OrganizationInfoName = 1

// flag for modification Org Description
const OrganizationInfoDescription = 2

// flag for modification Org OwnerUID (change owner user)
const OrganizationInfoOwnerUID = 4

// Отримати колекцію Organizations бази даних
func Organizations() *mongo.Collection {

	collection := database.DataBase().Collection(OrganizationsCollectionName)
	return collection
}
