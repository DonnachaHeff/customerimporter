package models

import (
	"errors"
	"sort"
	"strings"
)

func SortDomainsAlphabetically(recordInfo map[string]int) []string {
	keys := []string{}

	for key := range recordInfo {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func RetrieveDomainName(fullEmailAddress string) (string, error) {
	// find the final @ symbol within the email address
	at := strings.LastIndex(fullEmailAddress, "@")

	if at >= 0 {
		_, domain := fullEmailAddress[:at], fullEmailAddress[at+1:]

		return domain, nil
	}

	// should skip and return error
	return "", errors.New("empty domain name")
}
