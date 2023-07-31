package auth

import "strings"

// Згенерувати токен для підтвердження Email (при реєстрації нового користувача)
func GenerateEmailConfirmationToken(userID string, email string) string {

	//TODO: продумати як генерувати токен для підтвердження email (додати "сіль" секретну та хешувати
	//TODO: перевести email - tolower, trim

	emailLowCase := strings.ToLower(email)

	token := "email_confirm_token_" + userID + "_" + emailLowCase

	return token
}
