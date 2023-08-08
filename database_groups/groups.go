package database_groups

import (
	"context"
	"errors"
	"fmt"
	"log"

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

// Validate Groups collection in MongoDB, indexes and others
func ValidateGroupsCollection() {

	ctx := context.TODO()

	groups := Groups()

	err := database.CreateCollectionUniqueIndexOnField(ctx, groups, "unique_name")
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, groups, "name")
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, groups, "description")
	if err != nil {
		panic(err)
	}

	log.Println("GroupsCollection schema validated")
}

func CreateGroup(ctx context.Context, group GroupInfo) (createdGroup GroupInfo, err error) {
	groups := Groups()

	if primitive.ObjectID.IsZero(group.ID) {
		group.ID = primitive.NewObjectID()
	}

	if group.Name == "" {
		err = errors.New("group.Name not specified")
		return
	}

	if group.UniqueName == "" {
		err = errors.New("group.UniqueName not specified")
		return
	}

	result, insertErr := groups.InsertOne(ctx, group)
	if insertErr != nil {
		err = fmt.Errorf("error inserting Group: %s", insertErr.Error())
		return
	}

	group.ID = result.InsertedID.(primitive.ObjectID)

	return group, nil
}
