package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Валідація JWT (цю ф-цію зазвичай не повинні викликати напряму)
func validateTokenGeneric(signedToken string) (userClaims *UserClaim, err error) {
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

	return claims, nil
}

// Валідація токена для всіх випадків (крім реєстрації користувача)
func ValidateToken(signedToken string) (userClaims *UserClaim, err error) {
	claims, err := validateTokenGeneric(signedToken)
	if err != nil {
		return nil, err
	}

	if (claims.Flags & FlagRegistrationConfirmation) != 0 {
		return nil, errors.New("this token can only be used with registration confirmation")
	}

	return claims, nil
}

// Валідація токена для реєстрації користувача
func ValidateTokenForRegistrationConfirmation(signedToken string) (userClaims *UserClaim, err error) {
	userClaims, err = validateTokenGeneric(signedToken)
	if err != nil {
		return
	}

	if (userClaims.Flags & FlagRegistrationConfirmation) == 0 {
		return nil, errors.New("this token is not for registration confirmation")
	}

	return
}
