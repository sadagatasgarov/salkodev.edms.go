package database_users

const UserAccountOptionNone = 0

// Користувач повинен змінити пароль при логіні
const UserAccountOptionRequestChangePassword = 1

// Зміна пароля заборонена
const UserAccountOptionChangePasswordNotAllowed = 2

// Пароль ніколи не застаріває
const UserAccountOptionPasswordNeverExpires = 4

// Обліковий запис блоковано
const UserAccountOptionDisabled = 8

// Removes all non-supported bits, leaving only operational account options (user for API)
func PurifyAccountOptions(accountOptions int) int {
	resultOptions := 0

	if accountOptions&UserAccountOptionRequestChangePassword > 0 {
		resultOptions |= UserAccountOptionRequestChangePassword
	}

	if accountOptions&UserAccountOptionChangePasswordNotAllowed > 0 {
		resultOptions |= UserAccountOptionChangePasswordNotAllowed
	}

	if accountOptions&UserAccountOptionPasswordNeverExpires > 0 {
		resultOptions |= UserAccountOptionPasswordNeverExpires
	}

	if accountOptions&UserAccountOptionDisabled > 0 {
		resultOptions |= UserAccountOptionDisabled
	}

	return resultOptions
}
