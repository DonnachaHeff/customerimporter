package models

import (
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

func RetrieveDomainName(fullEmailAddress string) string {
	// find the final @ symbol within the email address
	at := strings.LastIndex(fullEmailAddress, "@")

	if at >= 0 {
		_, domain := fullEmailAddress[:at], fullEmailAddress[at+1:]

		return domain
	}

	return "invalidDomain"
}
