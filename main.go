package main

import (
	"github.com/DonnachaHeff/customerimporter/data"
	"github.com/DonnachaHeff/customerimporter/filemanager"
)

func main() {
	domains := filemanager.ReadCsvFile("./customers.csv")
	sortedDomains := data.SortDomainsAlphabetically(domains)
	filemanager.OutputSortedDomainsResultToFile(sortedDomains, domains)
}
