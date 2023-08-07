package core

import (
	"strings"

	"github.com/google/uuid"
)

// Generate new UID
func GenerateUID() string {
	uidStr := strings.ToUpper(uuid.New().String())
	return uidStr
}
