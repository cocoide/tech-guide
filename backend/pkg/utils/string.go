package utils

import (
	"errors"
	"strings"
)

func ExtractIDFromURL(u string) (string, error) {
	if u == "" {
		return "", errors.New("URL is empty")
	}
	parts := strings.Split(u, "/")
	if len(parts) == 0 {
		return "", errors.New("URL is not in expected format")
	}
	return parts[len(parts)-1], nil
}
