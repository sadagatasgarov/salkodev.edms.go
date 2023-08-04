package database

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// Find user by email
func FindUser(ctx context.Context, userEmail string) (user UserInfo, err error) {
	//знайти користувача в базі (логін - мейл)
	users := Users()

	email := strings.ToLower(userEmail)

	filter := bson.M{"email": email}
	err = users.FindOne(ctx, filter).Decode(&user)

	return user, err
}
