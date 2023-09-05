package database_departments

import (
	"context"
	"log"

	"github.com/AndrewSalko/salkodev.edms.go/database"
)

func ValidateSchema() {

	ctx := context.TODO()

	ValidateDepartmentsCollection(ctx)

	log.Println("Departments schema validated")
}

// Validate Departments collection in MongoDB, indexes and others
func ValidateDepartmentsCollection(ctx context.Context) {

	deps := Departments()

	err := database.CreateCollectionUniqueIndexOnField(ctx, deps, DepartmentInfoFieldUID)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, deps, DepartmentInfoFieldOrgUID)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, deps, DepartmentInfoFieldName)
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, deps, DepartmentInfoFieldDescription)
	if err != nil {
		panic(err)
	}
}
