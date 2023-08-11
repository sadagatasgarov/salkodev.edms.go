package database_groups

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateGroup(ctx context.Context, group GroupInfo) (createdGroup GroupInfo, err error) {
	groups := Groups()

	if primitive.ObjectID.IsZero(group.ID) {
		group.ID = primitive.NewObjectID()
	}

	if group.Name == "" {
		err = errors.New("group.Name not specified")
		return
	}

	if group.UniqueName == "" {
		err = errors.New("group.UniqueName not specified")
		return
	}

	result, insertErr := groups.InsertOne(ctx, group)
	if insertErr != nil {
		err = fmt.Errorf("error inserting Group: %s", insertErr.Error())
		return
	}

	group.ID = result.InsertedID.(primitive.ObjectID)

	return group, nil
}
