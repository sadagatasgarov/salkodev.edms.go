package database_groups

import (
	"context"
	"errors"

	"github.com/AndrewSalko/salkodev.edms.go/core"
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
)

// Check user in groups
func UserInGroups(ctx context.Context, userUID string, groupsUniqueNames []string) (member []bool, err error) {

	length := len(groupsUniqueNames)
	if length == 0 {
		err = errors.New("groupsUniqueNames empty")
		return
	}

	//find user to operate with
	user, err := database_users.FindUserByUID(ctx, userUID)
	if err != nil {
		return
	}

	currentGroups := core.CreateMapFromStrings(user.Groups)
	member = make([]bool, length)

	for i := 0; i < length; i++ {
		grName := groupsUniqueNames[i]
		_, inGroups := currentGroups[grName]
		if inGroups {
			member[i] = true
		}
	}

	return
}

// Check user in group
func UserInGroup(ctx context.Context, userUID string, groupUniqueName string) (member bool, err error) {

	uniqNames := []string{groupUniqueName}
	result, err := UserInGroups(ctx, userUID, uniqNames)
	if err != nil {
		return
	}

	member = result[0]
	return
}

// Check groups uniq names for Administrators
func CheckAdministratorsGroup(userGroups []string) error {

	for _, gr := range userGroups {
		if gr == database.AdministratorsGroupUniqueName {
			return nil
		}
	}

	return errors.New("administrators group membership required")
}
