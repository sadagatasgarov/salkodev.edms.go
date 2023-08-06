package database

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// Find user by email
func FindUser(ctx context.Context, userEmail string) (user UserInfo, err error) {
	//знайти користувача в базі (логін - мейл)
	users := Users()

	email := strings.ToLower(userEmail)
	err = ValidateValueSanitization(email)
	if err != nil {
		return
	}

	filter := bson.M{"email": email}
	err = users.FindOne(ctx, filter).Decode(&user)

	return user, err
}

// Find user and check user hash with actual hash in db
func FindUserAndCheckHash(ctx context.Context, userEmail string, userHashFromToken string) (user UserInfo, err error) {
	users := Users()

	email := strings.ToLower(userEmail)
	err = ValidateValueSanitization(email)
	if err != nil {
		return
	}

	filter := bson.M{"email": email}
	err = users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return
	}

	//звіряємо хеш користувача з тим, який надійшов в запиті jwt
	hashActual := user.Hash
	if userHashFromToken != hashActual {
		err = errors.New("relogin required")
		return
	}

	return user, err
}
