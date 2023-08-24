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

// Отримати колекцію Organizations бази даних
func Organizations() *mongo.Collection {

	collection := database.DataBase().Collection(OrganizationsCollectionName)
	return collection
}
