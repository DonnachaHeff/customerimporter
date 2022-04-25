package models

import (
	"reflect"
	"sort"
	"strings"
)

type CsvEntry struct {
	FirstName string
	LastName  string
	Email     string
	Gender    string
	IpAddress string
}

func GetAllDomains(data []CsvEntry) []string {
	var allDomains []string

	for _, records := range data {
		if !reflect.ValueOf(records).IsZero() {
			domain := retrieveDomainName(records.Email)
			allDomains = append(allDomains, domain)
		}
	}

	return allDomains
}

func CountOccurencesOfDomain(allDomains []string) map[string]int {
	frequency_records := make(map[string]int)

	for _, domain := range allDomains {
		// don't want to count empty occurences
		if len(domain) > 0 {
			// check if the domain exists within the map
			_, exist := frequency_records[domain]

			if exist {
				frequency_records[domain] += 1
			} else {
				frequency_records[domain] = 1
			}
		}
	}

	return frequency_records
}

func SortDomainsAlphabetically(recordInfo map[string]int) []string {
	keys := []string{}

	for key := range recordInfo {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func retrieveDomainName(fullEmailAddress string) string {
	// find the final @ symbol within the email address
	at := strings.LastIndex(fullEmailAddress, "@")
	if at >= 0 {
		// seperate the domain from the full address
		_, domain := fullEmailAddress[:at], fullEmailAddress[at+1:]

		return domain
	}

	return ""
}
