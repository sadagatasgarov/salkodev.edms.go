package database_folders

import (
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const FoldersCollectionName = "Folders"

// Max folders allowed for one organization
const FoldersMaxCountPerOrganization = 300

type FolderInfo struct {
	ID              primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UID             string             `bson:"uid" json:"uid" binding:"required"`
	CreationTime    time.Time          `bson:"creation_time" json:"creation_time" binding:"required"`
	OrganizationUID string             `bson:"org_uid" json:"org_uid" binding:"required"`
	DepartmentUID   string             `bson:"department_uid" json:"department_uid,omitempty"`
	Name            string             `bson:"name" json:"name" binding:"required"`
	Description     string             `bson:"description,omitempty" json:"description,omitempty"`
}

const FolderInfoFieldUID = "uid"
const FolderInfoFieldOrgUID = "org_uid"
const FolderInfoFieldDepartmentUID = "department_uid"
const FolderInfoFieldName = "name"
const FolderInfoFieldDescription = "description"
const FolderInfoFieldCreationTime = "creation_time"

// flag for modification folder's org_uid
const FolderInfoOrganizationUID = 1

// flag for modification folder's Name
const FolderInfoName = 2

// flag for modification folder's Description
const FolderInfoDescription = 4

// flag for modification folder's DepartmentUID
const FolderInfoDepartmentUID = 8

// Отримати колекцію Folders бази даних
func Folders() *mongo.Collection {

	collection := database.DataBase().Collection(FoldersCollectionName)
	return collection
}
