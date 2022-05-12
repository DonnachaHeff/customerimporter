package main

import (
	"github.com/DonnachaHeff/customerimporter/filemanager"
	"github.com/DonnachaHeff/customerimporter/models"
)

func main() {
	domains := filemanager.ReadCsvFile("./customers.csv")
	sortedDomains := models.SortDomainsAlphabetically(domains)
	filemanager.OutputSortedDomainsResultToFile(sortedDomains, domains)
}
