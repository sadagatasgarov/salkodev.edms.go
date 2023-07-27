package auth

import (
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Refresh JWT token (obtains new token based on current)
// TODO:  RefreshToken not implemented!
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value
	claims := &UserClaim{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return _JWTSecretKeyStr, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if _JWTSecretKeyStr == "" {
	// 	return "", errors.New("JWT secret key not found")
	// }

	// secretKeyBytes := []byte(_JWTSecretKeyStr)

	// expDate := jwt.NewNumericDate(time.Now().Add(48 * time.Hour))

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
	// 	RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: expDate},
	// 	Email:            email})

	// resultJWT, err := token.SignedString(secretKeyBytes)
	// if err != nil {
	// 	return "", err
	// }

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKeyBytes := []byte(_JWTSecretKeyStr)

	tokenString, err := token.SignedString(secretKeyBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
