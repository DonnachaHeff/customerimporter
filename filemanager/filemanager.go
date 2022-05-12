package filemanager

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DonnachaHeff/customerimporter/models"
)

func ReadCsvFile(filePath string) map[string]int {
	entries := make(map[string]int)

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		domain, err := models.RetrieveDomainName(record[2])
		if err != nil {
			log.Fatal(err)
		} else {
			if _, ok := entries[domain]; ok {
				entries[domain] += 1
			} else {
				entries[domain] = 1
			}
		}

	}

	return entries
}

func OutputSortedDomainsResultToFile(sortedDomains []string, recordInfo map[string]int) {
	// create file to write results to
	f, err := os.Create("Results")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	for _, key := range sortedDomains {
		// print each domain and it's total number of customers
		_, err = fmt.Fprintln(w, "Domain Name: ", key, ", Total Number of Customers: ", recordInfo[key])
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
}
