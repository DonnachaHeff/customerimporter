package main

import (
	"github.com/DonnachaHeff/customerimporter/filemanager"
	"github.com/DonnachaHeff/customerimporter/models"
)

func main() {
	records := filemanager.ReadCsvFile("./customers.csv")
	sortedDomains := models.SortDomainsAlphabetically(records)
	filemanager.OutputSortedDomainsResultToFile(sortedDomains, records)
}
