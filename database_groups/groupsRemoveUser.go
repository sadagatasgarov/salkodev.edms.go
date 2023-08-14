package database_groups

import (
	"context"
	"errors"

	"github.com/AndrewSalko/salkodev.edms.go/core"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"go.mongodb.org/mongo-driver/bson"
)

// Remove user from groups
func RemoveUser(ctx context.Context, actingUser database_users.UserInfo, userUID string, groupsUniqueNames []string) error {

	if len(groupsUniqueNames) == 0 {
		return errors.New("groupsUniqueNames empty")
	}

	//TODO: check if actingUser in Administrators group

	//find user to operate with
	user, err := database_users.FindUserByUID(ctx, userUID)
	if err != nil {
		return err
	}

	currentGroups := core.CreateMapFromStrings(user.Groups)

	needUpdate := false

	for q := 0; q < len(groupsUniqueNames); q++ {
		uniqName := groupsUniqueNames[q]
		_, inGroups := currentGroups[uniqName]
		if inGroups {
			delete(currentGroups, uniqName)
			needUpdate = true
		}
	}

	if needUpdate {
		//now update user's groups in MongoDB
		groupResultUniqNames := core.CreateStringsFromMapKeys(&currentGroups)

		users := database_users.Users()
		filter := bson.M{"_id": user.ID}
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "groups", Value: groupResultUniqNames}}}}

		//оновити в базі стан користувача що email підтверджено
		_, err := users.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return err
		}
	}

	return nil
}
