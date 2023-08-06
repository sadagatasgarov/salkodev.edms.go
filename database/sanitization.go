package database

import (
	"errors"
	"strings"
)

//const MongoForbiddenCharacters="${}"

func ValidateValueSanitization(val string) error {

	valTrimmed := strings.TrimSpace(val)

	if valTrimmed == "" {
		return errors.New("empty or whitespace argument not allowed")
	}

	if strings.HasPrefix(valTrimmed, "{") || strings.HasPrefix(valTrimmed, "$") {
		return errors.New("characters '{', '$' not allowed at start of string")
	}

	return nil
}
