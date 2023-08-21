package database_users

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Generates hash on critical user data, for controlling changes
func GenerateUserHash(uid string, orgUid string, name string, email string, accountOptions int, passwordHash string) string {

	dataStr := fmt.Sprintf("uid:%s orgUid:%s name:%s email:%s accountOptions:%x passwordHash:%s", uid, orgUid, name, email, accountOptions, passwordHash)
	data := []byte(dataStr)

	//hashing SHA256
	sha256Hash := sha256.Sum256(data)
	sha256HashString := hex.EncodeToString(sha256Hash[:])

	return sha256HashString
}
