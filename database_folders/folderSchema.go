package database_folders

import (
	"context"
	"log"

	"github.com/AndrewSalko/salkodev.edms.go/database"
)

func ValidateSchema() {

	ctx := context.TODO()

	ValidateFoldersCollection(ctx)

	log.Println("Folders schema validated")
}

// Validate Folders collection in MongoDB, indexes and others
func ValidateFoldersCollection(ctx context.Context) {

	deps := Folders()

	err := database.CreateCollectionUniqueIndexOnField(ctx, deps, FolderInfoFieldUID)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, deps, FolderInfoFieldOrgUID)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, deps, FolderInfoFieldDepartmentUID)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, deps, FolderInfoFieldName)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, deps, FolderInfoFieldDescription)
	if err != nil {
		panic(err)
	}
}
