package auth

import (
	"context"
	"strings"

	"github.com/AndrewSalko/salkodev.edms.go/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name           string             `json:"name" binding:"required"`
	Email          string             `json:"email" binding:"required"`
	Password       string             `json:"password" binding:"required"`
	EmailConfirmed bool               `json:"email_confirmed"`
}

// Find user by email
func FindUser(ctx context.Context, userEmail string) (UserInfo, error) {
	//знайти користувача в базі (логін - мейл)
	users := database.Users()

	email := strings.ToLower(userEmail)

	filter := bson.M{"email": email}
	var resultUser UserInfo
	errFindUser := users.FindOne(ctx, filter).Decode(&resultUser)

	return resultUser, errFindUser
}
