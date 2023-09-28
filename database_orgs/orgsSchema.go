package database_orgs

import (
	"context"
	"log"

	"github.com/AndrewSalko/salkodev.edms.go/database"
)

func ValidateSchema() {

	ctx := context.TODO()

	ValidateOrgsCollection(ctx)

	log.Println("Organizations schema validated")
}

// Validate Organizations collection in MongoDB, indexes and others
func ValidateOrgsCollection(ctx context.Context) {

	orgs := Organizations()

	err := database.CreateCollectionUniqueIndexOnField(ctx, orgs, OrganizationInfoFieldUID)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, orgs, OrganizationInfoFieldOwnerUID)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, orgs, OrganizationInfoFieldName)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, orgs, OrganizationInfoFieldDescription)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, orgs, OrganizationInfoFieldCreationTime)
	if err != nil {
		panic(err)
	}

}
