package auth

import (
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
)

// long secret string with random chars
const EnvNameJWTSecretKey string = "SALKODEV_EDMS_JWT_SECRET_KEY"

var _JWTSecretKeyStr string = os.Getenv(EnvNameJWTSecretKey)

type UserClaim struct {
	jwt.RegisteredClaims
	Email string
}
