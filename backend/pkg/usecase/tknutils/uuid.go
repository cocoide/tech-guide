package tknutils

import (
	"github.com/google/uuid"
	"regexp"
)

func GenerateStrUUID() string {
	return uuid.New().String()
}

func IsValidUUID(u string) bool {
	r := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
	return r.MatchString(u)
}
