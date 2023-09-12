package database_folders

import (
	"context"
	"fmt"

	"github.com/AndrewSalko/salkodev.edms.go/core"
	"github.com/AndrewSalko/salkodev.edms.go/database_orgs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Creates new Folder
func CreateFolder(ctx context.Context, folder FolderInfo) (createdFolder FolderInfo, err error) {
	folders := Folders()

	_, err = database_orgs.FindOrganizationByUID(ctx, folder.OrganizationUID)
	if err != nil {
		return
	}

	//check how many departments are created in the organization
	// count, err := deps.CountDocuments(ctx, bson.M{DepartmentInfoFieldOrgUID: department.OrganizationUID})
	// if err != nil {
	// 	return
	// }

	// if count > DepartmentsMaxCount {
	// 	err = errors.New("the maximum number of departments already created for this organization")
	// 	return
	// }

	//TODO: check department name (in current organization)

	if primitive.ObjectID.IsZero(folder.ID) {
		folder.ID = primitive.NewObjectID()
	}

	//if folder uid unspecified, generate new uid
	if folder.UID == "" {
		folder.UID = core.GenerateUID()
	}

	result, insertErr := folders.InsertOne(ctx, folder)
	if insertErr != nil {
		err = fmt.Errorf("error inserting Department: %s", insertErr.Error())
		return
	}

	folder.ID = result.InsertedID.(primitive.ObjectID)

	return folder, nil
}
