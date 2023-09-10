package database_users

import (
	"context"
	"log"
	"os"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const AdminAccountName = "admin"
const AdminAccountMail = "admin@system"

// if db is new and Admin account not found, pass must be set as env-var (only for first start)
const AdminAccountPassEnv = "SALKODEV_EDMS_ADMIN_PSW"
const AdminAccountUID = "26040000-0000-0000-0000-0000000000AD"

// _id for Admin user (12 byte hex)
const AdminAccountIDStr = "2604000000000000000000AD"

func ValidateSchema() {

	ctx := context.TODO()

	ValidateUsersCollection(ctx)
	validateAdminAccount(ctx)

	log.Println("Users schema validated")
}

// Validate Users collection in MongoDB, indexes and others
func ValidateUsersCollection(ctx context.Context) {

	users := Users()

	err := database.CreateCollectionUniqueIndexOnField(ctx, users, "email")
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionUniqueIndexOnField(ctx, users, "uid")
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, users, "org_uid")
	if err != nil {
		panic(err)
	}

	err = database.CreateCollectionIndexOnField(ctx, users, "name")
	if err != nil {
		panic(err)
	}
}

func validateAdminAccount(ctx context.Context) {

	users := Users()

	//find admin account
	filter := bson.M{"email": AdminAccountMail}
	var user UserInfo
	err := users.FindOne(ctx, filter).Decode(&user)
	notFound := false
	if err != nil {
		if err == mongo.ErrNoDocuments {
			notFound = true
		} else {
			panic(err)
		}
	}

	if notFound {
		//create admin account
		objID, errObjHex := primitive.ObjectIDFromHex(AdminAccountIDStr)
		if errObjHex != nil {
			panic(errObjHex)
		}

		admPsw := os.Getenv(AdminAccountPassEnv)
		if admPsw == "" {
			panic("For new database you must specify Admin password as env.variable: " + AdminAccountPassEnv)
		}

		user.ID = objID
		user.Name = AdminAccountName
		user.Email = AdminAccountMail //stub mail used as login
		user.AccountOptions = UserAccountOptionPasswordNeverExpires
		user.Password = auth.HashPassword(admPsw)
		user.EmailConfirmed = true
		user.UID = AdminAccountUID
		user.Groups = []string{database.AdministratorsGroupUniqueName}
		user.Hash = GenerateUserHash(user.UID, user.OrganizationUID, user.DepartmentUID, user.Name, user.Email, user.AccountOptions, user.Password)

		_, insertErr := users.InsertOne(ctx, user)
		if insertErr != nil {
			panic(insertErr)
		}
	}

	//TODO: validate user account if found
}
