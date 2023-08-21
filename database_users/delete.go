package database_users

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// Delete user. uid field is key
func DeleteUser(ctx context.Context, uid string) (err error) {

	if uid == "" {
		return errors.New("uid empty")
	}

	user, err := FindUserByUID(ctx, uid)
	if err != nil {
		return
	}

	users := Users()

	filter := bson.M{"_id": user.ID}

	_, err = users.DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	return nil
}
