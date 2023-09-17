package utils

import (
	"fmt"
	"regexp"
)

func ExtractDomainNameFromURL(url string) (string, error) {
	regex := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n]+)`)
	matches := regex.FindStringSubmatch(url)
	if len(matches) >= 2 {
		return matches[1], nil
	}
	return "", fmt.Errorf("failed to extract domain from URL")
}
