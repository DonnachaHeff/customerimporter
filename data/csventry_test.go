package data

import (
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
