package database_folders

import (
	"context"

	"github.com/AndrewSalko/salkodev.edms.go/core"
	"go.mongodb.org/mongo-driver/bson"
)

// Find Folder by uid
func FindFolderByUID(ctx context.Context, uid string) (folder FolderInfo, err error) {

	_, err = core.UIDFromString(uid)
	if err != nil {
		return
	}

	folders := Folders()

	filter := bson.M{"uid": uid}
	err = folders.FindOne(ctx, filter).Decode(&folder)

	return folder, err
}
