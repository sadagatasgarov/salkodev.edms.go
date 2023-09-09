package database_departments

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// Delete department. uid field is key
func DeleteDepartment(ctx context.Context, uid string) (err error) {

	if uid == "" {
		return errors.New("uid empty")
	}

	dep, err := FindDepartmentByUID(ctx, uid)
	if err != nil {
		return
	}

	//TODO: can't remove department if it has users

	deps := Departments()

	filter := bson.M{"_id": dep.ID}

	_, err = deps.DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	return nil
}
