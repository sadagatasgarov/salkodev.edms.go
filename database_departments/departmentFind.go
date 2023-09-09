package database_departments

import (
	"context"

	"github.com/AndrewSalko/salkodev.edms.go/core"
	"go.mongodb.org/mongo-driver/bson"
)

// Find Department by uid
func FindDepartmentByUID(ctx context.Context, uid string) (org DepartmentInfo, err error) {

	_, err = core.UIDFromString(uid)
	if err != nil {
		return
	}

	deps := Departments()

	filter := bson.M{"uid": uid}
	err = deps.FindOne(ctx, filter).Decode(&org)

	return org, err
}
