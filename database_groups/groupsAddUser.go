package database_groups

import (
	"context"
	"errors"

	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"go.mongodb.org/mongo-driver/bson"
)

// Add user to groups
func AddUser(ctx context.Context, actingUser database_users.UserInfo, userUID string, groupsUniqueNames []string) error {

	if groupsUniqueNames == nil {
		return errors.New("groupsUniqueNames == nil")
	}

	if len(groupsUniqueNames) == 0 {
		return errors.New("groupsUniqueNames empty")
	}

	//TODO: check if actingUser in Administrators group

	//find user to operate with
	user, err := database_users.FindUserByUID(ctx, userUID)
	if err != nil {
		return err
	}

	currentGroups := make(map[string]string)

	//Read user current groups
	groups := user.Groups
	if groups != nil {
		for i := 0; i < len(groups); i++ {
			gr := groups[i]
			if gr != "" {
				currentGroups[gr] = ""
			}
		}
	}

	//var groupsToAdd []string = make([]string, len(groupsUniqueNames))
	needUpdate := false

	for q := 0; q < len(groupsUniqueNames); q++ {
		uniqName := groupsUniqueNames[q]
		_, inGroups := currentGroups[uniqName]
		if !inGroups {
			currentGroups[uniqName] = ""
			needUpdate = true
		}
	}

	if needUpdate {

		//now update user's groups in MongoDB
		groupResultUniqNames := make([]string, 0, len(currentGroups))
		for k := range currentGroups {
			groupResultUniqNames = append(groupResultUniqNames, k)
		}

		users := database_users.Users()
		filter := bson.M{"_id": user.ID}
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "groups", Value: groupResultUniqNames}}}}

		//оновити в базі стан користувача що email підтверджено
		_, err := users.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
