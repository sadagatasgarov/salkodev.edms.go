package auth

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Check password policy
func CheckPasswordPolicy(password string) error {

	if password == "" {
		return errors.New("password is empty")
	}

	//TODO: розробити політику паролів
	if len(password) < 3 {
		return errors.New("password too short")
	}

	return nil
}
