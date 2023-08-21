package database_users

import (
	"context"
	"errors"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"go.mongodb.org/mongo-driver/bson"
)

// Modify user. fields specify which property will be changed (see flags: )
func ModifyUser(ctx context.Context, userData UserInfo, userFields int) (err error) {

	//user.UID is required and key field
	if userData.UID == "" {
		return errors.New("uid empty")
	}

	user, err := FindUserByUID(ctx, userData.UID)
	if err != nil {
		return
	}

	users := Users()

	upd := bson.D{}

	if userFields&UserInfoAccountOptions > 0 {
		user.AccountOptions = PurifyAccountOptions(userData.AccountOptions) //for hash regen
		upd = append(upd, bson.E{Key: UserInfoFieldAccountOptions, Value: user.AccountOptions})
	}

	if userFields&UserInfoEmail > 0 {
		upd = append(upd, bson.E{Key: UserInfoFieldEmail, Value: userData.Email})
		user.Email = userData.Email //for hash regen
	}

	if userFields&UserInfoEmailConfirmed > 0 {
		upd = append(upd, bson.E{Key: UserInfoFieldEmailConfirmed, Value: userData.EmailConfirmed})
	}

	if userFields&UserInfoName > 0 {
		upd = append(upd, bson.E{Key: UserInfoFieldName, Value: userData.Name})
		user.Name = userData.Name //for hash regen
	}

	if userFields&UserInfoOrganizationUID > 0 {
		upd = append(upd, bson.E{Key: UserInfoFieldOrganizationUID, Value: userData.OrganizationUID})
		user.OrganizationUID = userData.OrganizationUID //for hash regen
	}

	if userFields&UserInfoPassword > 0 {
		//TODO: check password policy
		user.Password = auth.HashPassword(userData.Password)
		upd = append(upd, bson.E{Key: UserInfoFieldPassword, Value: user.Password})
	}

	//update user hash (UserInfoFieldHash)
	hash := GenerateUserHash(user.UID, user.OrganizationUID, user.Name, user.Email, user.AccountOptions, user.Password)
	upd = append(upd, bson.E{Key: UserInfoFieldHash, Value: hash})

	update := bson.D{{Key: "$set", Value: upd}}

	_, err = users.UpdateByID(ctx, user.ID, update)
	if err != nil {
		return
	}

	return nil
}
