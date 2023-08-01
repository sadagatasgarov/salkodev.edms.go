package auth

import (
	"os"
)

// long secret string with random chars
const EnvNameJWTSecretKey string = "SALKODEV_EDMS_JWT_SECRET_KEY"

var _JWTSecretKeyStr string = os.Getenv(EnvNameJWTSecretKey)
