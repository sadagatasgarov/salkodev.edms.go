package database_users

import (
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Ім'я колекції Users (користувачі системи)
const UsersCollectionName = "Users"

type UserInfo struct {
	ID              primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UID             string             `bson:"uid" json:"uid" binding:"required"`
	OrganizationUID string             `bson:"org_uid" json:"org_uid,omitempty"`
	Name            string             `bson:"name" json:"name" binding:"required"`
	Email           string             `bson:"email" json:"email" binding:"required"`
	AccountOptions  int                `bson:"account_options" json:"account_options" binding:"required"`
	Password        string             `bson:"password" json:"password" binding:"required"` //password hash
	EmailConfirmed  bool               `bson:"email_confirmed" json:"email_confirmed"`
	Hash            string             `bson:"hash" json:"hash"` //хеш користувача (для виявлення змін)
	Groups          []string           `bson:"groups" json:"groups"`
}

const UserInfoFieldUID = "uid"
const UserInfoFieldOrganizationUID = "org_uid"
const UserInfoFieldName = "name"
const UserInfoFieldEmail = "email"
const UserInfoFieldAccountOptions = "account_options"
const UserInfoFieldPassword = "password"
const UserInfoFieldEmailConfirmed = "email_confirmed"
const UserInfoFieldHash = "hash"
const UserInfoFieldGroups = "groups"

// Flag for UserModify - modify OrganizationUID
const UserInfoOrganizationUID = 1

// Flag for UserModify - modify Name
const UserInfoName = 2

// Flag for UserModify - modify Email
const UserInfoEmail = 4

// Flag for UserModify - modify AccountOptions
const UserInfoAccountOptions = 8

// Flag for UserModify - modify Password
const UserInfoPassword = 16

// Flag for UserModify - modify EmailConfirmed
const UserInfoEmailConfirmed = 32

// Отримати колекцію Users бази даних
func Users() *mongo.Collection {
	collection := database.DataBase().Collection(UsersCollectionName)
	return collection
}
