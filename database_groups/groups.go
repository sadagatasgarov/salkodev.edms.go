package database_groups

import (
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const GroupsCollectionName = "Groups"

type GroupInfo struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UniqueName  string             `bson:"unique_name" json:"unique_name" binding:"required"`
	Name        string             `bson:"name" json:"name" binding:"required"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
}

// Отримати колекцію Groups бази даних
func Groups() *mongo.Collection {

	collection := database.DataBase().Collection(GroupsCollectionName)
	return collection
}
