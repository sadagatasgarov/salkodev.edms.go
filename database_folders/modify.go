package database_folders

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// Modify folder. fields specify which property will be changed (see flags)
func ModifyFolder(ctx context.Context, folderData FolderInfo, folderFields int) (err error) {

	//UID is required and key field
	if folderData.UID == "" {
		return errors.New("uid empty")
	}

	dep, err := FindFolderByUID(ctx, folderData.UID)
	if err != nil {
		return
	}

	deps := Folders()
	upd := bson.D{}

	if folderFields&FolderInfoOrganizationUID > 0 {
		upd = append(upd, bson.E{Key: FolderInfoFieldOrgUID, Value: folderData.OrganizationUID})
		dep.OrganizationUID = folderData.OrganizationUID
	}

	if folderFields&FolderInfoName > 0 {
		upd = append(upd, bson.E{Key: FolderInfoFieldName, Value: folderData.Name})
		dep.Name = folderData.Name //TODO: check if department name exists (in one org)
	}

	if folderFields&FolderInfoDescription > 0 {
		upd = append(upd, bson.E{Key: FolderInfoFieldDescription, Value: folderData.Description})
		dep.Description = folderData.Description
	}

	update := bson.D{{Key: "$set", Value: upd}}

	_, err = deps.UpdateByID(ctx, dep.ID, update)
	if err != nil {
		return
	}

	return nil
}
