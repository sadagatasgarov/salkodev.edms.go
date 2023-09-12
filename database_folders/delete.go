package database_folders

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// Delete folder. uid field is key
func DeleteFolder(ctx context.Context, uid string) (err error) {

	if uid == "" {
		return errors.New("uid empty")
	}

	dep, err := FindFolderByUID(ctx, uid)
	if err != nil {
		return
	}

	deps := Folders()

	filter := bson.M{"_id": dep.ID}

	_, err = deps.DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	return nil
}
