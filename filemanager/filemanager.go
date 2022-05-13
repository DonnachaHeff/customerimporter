package filemanager

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DonnachaHeff/customerimporter/data"
	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
)

func ReadCsvFile(filePath string) map[string]int {
	entries := make(map[string]int)

	var f *os.File

	// set action for opening file
	action := func(attempt uint) error {
		var err error

		f, err = os.Open(filePath)

		return err
	}

	// retry action 3 times
	err := retry.Retry(action, strategy.Limit(3))
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

		domain, err := data.RetrieveDomainName(record[2])
		if err != nil {
			log.Print(err)
			fmt.Printf("Invalid Name: %s\n\n", record[2])
		} else {
			if _, exists := entries[domain]; exists {
				entries[domain] += 1
			} else {
				entries[domain] = 1
			}
		}

	}

	return entries
}

func OutputSortedDomainsResultToFile(sortedDomains []string, domains map[string]int) {
	// create file to write results to
	var f *os.File

	// action for creating new file
	action := func(attempt uint) error {
		var err error

		f, err = os.Create("Results")

		return err
	}

	// retry action 3 times
	err := retry.Retry(action, strategy.Limit(3))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	for _, key := range sortedDomains {
		// print each domain and it's total number of customers
		_, err = fmt.Fprintln(w, "Domain Name: ", key, ", Total Number of Customers: ", domains[key])
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
}
