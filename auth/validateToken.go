package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Валідація JWT
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&UserClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(_JWTSecretKeyStr), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*UserClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now().Local()) {
		err = errors.New("token expired")
		return
	}
	return
}
