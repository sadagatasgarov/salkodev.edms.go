package routes

import (
	"strings"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/gin-gonic/gin"
)

// Перевіряє токен JWT в запиті
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		//строка tokenString містить Bearer <token>, розділимо на дві частини пробілом
		parts := strings.Split(tokenString, " ")
		if len(parts) < 2 {
			context.JSON(401, gin.H{"error": "authorization header invalid"})
			context.Abort()
			return
		}

		token := parts[1] //second part

		_, err := auth.ValidateToken(token)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
