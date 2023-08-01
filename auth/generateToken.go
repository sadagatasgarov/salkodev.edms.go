package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Common-used JWT token
func GenerateToken(email string) (string, error) {
	token, err := _GenerateToken(email, 0)
	return token, err
}

// JWT token for user registration confirmation. It can't be used anywhere instead of one confirmation function
func GenerateTokenForUserRegistration(email string) (string, error) {
	token, err := _GenerateToken(email, FlagRegistrationConfirmation)
	return token, err
}

func _GenerateToken(email string, flags uint) (string, error) {

	if _JWTSecretKeyStr == "" {
		return "", errors.New("JWT secret key not found")
	}

	secretKeyBytes := []byte(_JWTSecretKeyStr)

	var token *jwt.Token

	expireTime := time.Now().Add(1 * time.Hour) //normal token lives 1 hour

	if flags&FlagRegistrationConfirmation != 0 {
		expireTime = time.Now().Add(24 * time.Hour) //confirm reg.token lives 24 hours
	}

	expDate := jwt.NewNumericDate(expireTime)
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: expDate},
		Email:            email,
		Flags:            flags})

	resultJWT, err := token.SignedString(secretKeyBytes)
	if err != nil {
		return "", err
	}

	return resultJWT, nil
}
