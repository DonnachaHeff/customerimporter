package main

import (
	"github.com/DonnachaHeff/customerimporter/filemanager"
	"github.com/DonnachaHeff/customerimporter/models"
)

func main() {
	records := filemanager.ReadCsvFile("./customers.csv")
	allDomains := models.GetAllDomains(records)
	recordInfo := models.CountOccurencesOfDomain(allDomains)
	sortedDomains := models.SortDomainsAlphabetically(recordInfo)
	filemanager.OutputSortedDomainsResultToFile(sortedDomains, recordInfo)
}
