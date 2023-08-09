package database_groups

import (
	"context"
	"log"

	"github.com/AndrewSalko/salkodev.edms.go/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const AdministratorsGroupID = "260400060000000000000002"
const AdministratorsGroupName = "Administrators"
const AdministratorsGroupUniqueName = "administrators"
const AdministratorsGroupDescription = "Manage system settings, access administrative tasks"

func ValidateSchema() {

	ctx := context.TODO()

	ValidateGroupsCollection(ctx)
	ValidateGroups(ctx)

	log.Println("Groups schema validated")
}

// Validate Groups collection in MongoDB, indexes and others
func ValidateGroupsCollection(ctx context.Context) {

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
}

// Validate system groups
func ValidateGroups(ctx context.Context) {

	err := validateGroup(ctx, AdministratorsGroupID, AdministratorsGroupName, AdministratorsGroupUniqueName, AdministratorsGroupDescription)
	if err != nil {
		panic(err)
	}

}

func validateGroup(ctx context.Context, id string, name string, uniqueName string, description string) error {

	objId, errObjHex := primitive.ObjectIDFromHex(id)
	if errObjHex != nil {
		return errObjHex
	}

	groups := Groups()

	filter := bson.M{"_id": objId}
	var group GroupInfo
	err := groups.FindOne(ctx, filter).Decode(&group)
	notFound := false
	if err != nil {
		if err == mongo.ErrNoDocuments {
			notFound = true
		} else {
			return err
		}
	}

	if notFound {
		group.Description = description
		group.ID = objId
		group.Name = name
		group.UniqueName = uniqueName

		_, insertErr := groups.InsertOne(ctx, group)
		if insertErr != nil {
			return insertErr
		}
	}

	return nil
}
