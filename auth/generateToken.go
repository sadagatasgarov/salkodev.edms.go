package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string) (string, error) {

	if _JWTSecretKeyStr == "" {
		return "", errors.New("JWT secret key not found")
	}

	secretKeyBytes := []byte(_JWTSecretKeyStr)

	expDate := jwt.NewNumericDate(time.Now().Add(48 * time.Hour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: expDate},
		Email:            email})

	resultJWT, err := token.SignedString(secretKeyBytes)
	if err != nil {
		return "", err
	}

	return resultJWT, nil
}
