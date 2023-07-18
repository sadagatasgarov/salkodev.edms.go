package auth

// Згенерувати токен для підтвердження Email (при реєстрації нового користувача)
func GenerateEmailConfirmationToken(userID string, email string) string {

	//TODO: продумати як генерувати токен для підтвердження email (додати "сіль" секретну та хешувати
	//TODO: перевести email - tolower, trim
	token := "email_confirm_token_" + userID + "_" + email

	return token
}
