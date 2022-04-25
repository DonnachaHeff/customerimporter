package models

import (
	"reflect"
	"testing"
)

func TestSortDomainsAlphabetically(t *testing.T) {
	recordInfo := map[string]int{
		"gmail.com":  7,
		"yahoo.ie":   2,
		"bing.co.uk": 1,
	}
	expectedOutput := []string{"bing.co.uk", "gmail.com", "yahoo.ie"}

	output := SortDomainsAlphabetically(recordInfo)

	for i, v := range output {
		if v != expectedOutput[i] {
			t.Errorf("TestSortDomainsAlphabetically Failed!")
		} else {
			t.Logf("TestSortDomainsAlphabetically Passed!")
		}
	}
}

func TestCountOccurencesOfDomain(t *testing.T) {
	allDomains := []string{"gmail.com", "gmail.com", "yahoo.ie"}
	expectedOutput := map[string]int{
		"gmail.com": 2,
		"yahoo.ie":  1,
	}

	output := CountOccurencesOfDomain(allDomains)

	if reflect.DeepEqual(expectedOutput, output) != true {
		t.Errorf("TestCountOccurencesOfDomain Failed!")
	} else {
		t.Logf("TestCountOccurencesOfDomain Passed!")
	}
}

func TestGetAllDomains(t *testing.T) {
	gmailCsvEntry := CsvEntry{FirstName: "stephen", LastName: "king", Email: "www.stephenking@gmail.com"}
	yahooCsvEntry := CsvEntry{FirstName: "tom", LastName: "robbins", Email: "www.tomrobbins@yahoo.ie"}
	data := []CsvEntry{gmailCsvEntry, yahooCsvEntry}
	expectedOutput := []string{"gmail.com", "yahoo.ie"}

	output := GetAllDomains(data)

	for i, v := range output {
		if v != expectedOutput[i] {
			t.Errorf("TestGetAllDomains Failed!")
		} else {
			t.Logf("TestGetAllDomains Passed!")
		}
	}
}
