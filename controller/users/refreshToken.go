package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database"
	"github.com/gin-gonic/gin"
)

// Refresh JWT token (obtains new token based on current)
func RefreshToken(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	claim, found := c.Get(auth.AuthUserClaimKey)
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": auth.AuthUserClaimKey + " not found"})
		return
	}

	userClaim := claim.(*auth.UserClaim)
	user, err := database.FindUserAndCheckHash(ctx, userClaim.Email, userClaim.UserHash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateToken(user.Email, user.Hash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	//повертаємо токен JWT
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// func RefreshToken(w http.ResponseWriter, r *http.Request) {
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	tknStr := c.Value
// 	claims := &UserClaim{}
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return _JWTSecretKeyStr, nil
// 	})

// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	// We ensure that a new token is not issued until enough time has elapsed
// 	// In this case, a new token will only be issued if the old token is within
// 	// 30 seconds of expiry. Otherwise, return a bad request status
// 	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// if _JWTSecretKeyStr == "" {
// 	// 	return "", errors.New("JWT secret key not found")
// 	// }

// 	// secretKeyBytes := []byte(_JWTSecretKeyStr)

// 	// expDate := jwt.NewNumericDate(time.Now().Add(48 * time.Hour))

// 	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
// 	// 	RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: expDate},
// 	// 	Email:            email})

// 	// resultJWT, err := token.SignedString(secretKeyBytes)
// 	// if err != nil {
// 	// 	return "", err
// 	// }

// 	// Now, create a new token for the current use, with a renewed expiration time
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	secretKeyBytes := []byte(_JWTSecretKeyStr)

// 	tokenString, err := token.SignedString(secretKeyBytes)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the new token as the users `token` cookie
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	})
// }
