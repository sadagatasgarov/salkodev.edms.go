package auth

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// long secret string with random chars
const EnvNameJWTSecretKey string = "SALKODEV_EDMS_JWT_SECRET_KEY"

var _JWTSecretKeyStr string = os.Getenv(EnvNameJWTSecretKey)

type UserClaim struct {
	jwt.RegisteredClaims
	Email string
}

func GenerateJwtToken(email string) (string, error) {

	if _JWTSecretKeyStr == "" {
		return "", errors.New("JWT secret key not found")
	}

	secretKeyBytes := []byte(_JWTSecretKeyStr)

	expDate := jwt.NewNumericDate(time.Now().Add(48 * time.Hour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: expDate},
		Email:            email})

	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(48 * time.Hour)
	// claims["user"] = email

	resultJWT, err := token.SignedString(secretKeyBytes)
	if err != nil {
		return "", err
	}

	return resultJWT, nil
}
