package core

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// Generate new UID
func GenerateUID() string {
	uidStr := strings.ToUpper(uuid.New().String())
	return uidStr
}

// Validate uid input
func UIDFromString(uidInputStr string) (uid uuid.UUID, err error) {

	uid, err = uuid.Parse(uidInputStr)
	if err != nil {
		return
	}

	return
}

func UIDFromStringWithArg(uidInputStr string, argName string) (uid uuid.UUID, err error) {

	if uidInputStr == "" {
		err = errors.New(argName + " unspecified")
		return
	}

	uid, err = uuid.Parse(uidInputStr)
	if err != nil {
		err = errors.New(argName + " " + err.Error())
	}

	return
}
