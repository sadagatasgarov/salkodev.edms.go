package database

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	ID              primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UID             string             `bson:"uid" json:"uid" binding:"required"`
	OrganizationUID string             `bson:"org_uid" json:"org_uid"`
	Name            string             `bson:"name" json:"name" binding:"required"`
	Email           string             `bson:"email" json:"email" binding:"required"`
	AccountOptions  int                `bson:"account_options" json:"account_options" binding:"required"`
	Password        string             `bson:"password" json:"password" binding:"required"`
	EmailConfirmed  bool               `bson:"email_confirmed" json:"email_confirmed"`
	Hash            string             `bson:"hash" json:"hash"` //хеш користувача (для виявлення змін)
}

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
		user.UID = generateUID()
	}

	//Org UID not required

	//розрахувати хеш важливих даних користувача
	user.Hash = generateUserHash(user.UID, user.OrganizationUID, user.Name, user.Email, user.AccountOptions, user.Password)

	result, insertErr := users.InsertOne(ctx, user)
	if insertErr != nil {
		err = fmt.Errorf("error inserting User: %s", insertErr.Error())
		return
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

// Generate new UID
func generateUID() string {
	uidStr := uuid.New().String()
	return uidStr
}

// Generates hash on critical user data, for controlling changes
func generateUserHash(uid string, orgUid string, name string, email string, accountOptions int, passwordHash string) string {

	dataStr := fmt.Sprintf("uid:%s orgUid:%s name:%s email:%s accountOptions:%x passwordHash:%s", uid, orgUid, name, email, accountOptions, passwordHash)
	data := []byte(dataStr)

	//hashing SHA256
	sha256Hash := sha256.Sum256(data)
	sha256HashString := hex.EncodeToString(sha256Hash[:])

	return sha256HashString
}
