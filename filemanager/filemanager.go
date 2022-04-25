package filemanager

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/DonnachaHeff/customerimporter/models"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadCsvFile(filePath string) []models.CsvEntry {
	var csvEntries []models.CsvEntry

	f, err := os.Open(filePath)
	checkError(err)
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	checkError(err)

	for _, line := range records {
		csvEntry := models.CsvEntry{
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

func OutputSortedDomainsResultToFile(sortedDomains []string, recordInfo map[string]int) {
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
