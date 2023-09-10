package database_users

import (
	"context"
	"errors"
	"fmt"

	"github.com/AndrewSalko/salkodev.edms.go/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Creates new User in Users collection
func CreateUser(ctx context.Context, user UserInfo) (createdUser UserInfo, err error) {
	users := Users()

	if primitive.ObjectID.IsZero(user.ID) {
		user.ID = primitive.NewObjectID()
	}

	//TODO: Password must be hashed here - validate it

	//Name, Email, Password required
	if user.Name == "" {
		err = errors.New("user.Name not specified")
		return
	}

	if user.Email == "" {
		err = errors.New("user.Email not specified")
		return
	}

	if user.Password == "" {
		err = errors.New("user.Password (hash) not specified")
		return
	}

	//generate new UID if not specified
	if user.UID == "" {
		user.UID = core.GenerateUID()
	}

	//Org UID not required

	//розрахувати хеш важливих даних користувача
	user.Hash = GenerateUserHash(user.UID, user.OrganizationUID, user.DepartmentUID, user.Name, user.Email, user.AccountOptions, user.Password)

	result, insertErr := users.InsertOne(ctx, user)
	if insertErr != nil {
		err = fmt.Errorf("error inserting User: %s", insertErr.Error())
		return
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}
