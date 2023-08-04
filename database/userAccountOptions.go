package database

const UserAccountOptionNone = 0

// Користувач повинен змінити пароль при логіні
const UserAccountOptionRequestChangePassword = 1

// Зміна пароля заборонена
const UserAccountOptionChangePasswordNotAllowed = 2

// Пароль ніколи не застаріває
const UserAccountOptionPasswordNeverExpires = 4

// Обліковий запис блоковано
const UserAccountOptionDisabled = 8
