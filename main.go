package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

func main() {
	records := readCsvFile("./customers.csv")
	allDomains := getAllDomains(records)
	recordInfo := countOccurencesOfDomain(allDomains)
	sortedDomains := sortDomainsAlphabetically(recordInfo)
	outputSortedDomainsResultToFile(sortedDomains, recordInfo)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func outputSortedDomainsResultToFile(sortedDomains []string, recordInfo map[string]int) {
	// create file to write results to
	f, err := os.Create("Results")
	checkError(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	for _, key := range sortedDomains {
		// print each domain and it's total number of customers
		_, err = fmt.Fprintln(w, "Domain Name: ", key, ", Total Number of Customers: ", recordInfo[key])
		checkError(err)
	}
	w.Flush()
}

func sortDomainsAlphabetically(recordInfo map[string]int) []string {
	keys := []string{}

	for key := range recordInfo {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func countOccurencesOfDomain(allDomains []string) map[string]int {
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

func getAllDomains(data []CsvEntry) []string {
	var allDomains []string

	for _, records := range data {
		if !reflect.ValueOf(records).IsZero() {
			domain := retrieveDomainName(records.Email)
			allDomains = append(allDomains, domain)
		}
	}

	return allDomains
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

func readCsvFile(filePath string) []CsvEntry {
	var csvEntries []CsvEntry

	f, err := os.Open(filePath)
	checkError(err)
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	checkError(err)

	for _, line := range records {
		csvEntry := CsvEntry{
			FirstName: line[0],
			LastName:  line[1],
			Email:     line[2],
			Gender:    line[3],
			IpAddress: line[4],
		}

		csvEntries = append(csvEntries, csvEntry)
	}

	return csvEntries
}
