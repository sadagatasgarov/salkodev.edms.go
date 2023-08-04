package auth

import jwt "github.com/golang-jwt/jwt/v5"

// UserClaim.Flags - флаг означає що це токен для підтвердження реєстрації користувача
// цей токен не можна використовувати ніде в інших функціях
const FlagRegistrationConfirmation = 1

// Змінная для передачі через контекст, яка містить UserClaim (безпеку токена вже перевірено)
const AuthUserClaimKey = "salkodev-jwt-userclaim"

type UserClaim struct {
	jwt.RegisteredClaims
	Email    string `json:"email"`
	Flags    uint   `json:"flags,omitempty"`     //позначки-флаги для окремих типів jwt
	UserHash string `json:"user_hash,omitempty"` //хеш користувача (для виявлення змін)
}
