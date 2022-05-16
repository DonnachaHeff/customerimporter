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

func TestSortDomainsAlphabetically_WhenNameIsSameButAddressIsDifferent(t *testing.T) {
	recordInfo := map[string]int{
		"bing.com":   7,
		"bing.ie":    2,
		"bing.co.uk": 1,
	}
	expectedOutput := []string{"bing.co.uk", "bing.com", "bing.ie"}

	output := SortDomainsAlphabetically(recordInfo)

	for i, v := range output {
		if v != expectedOutput[i] {
			t.Errorf("TestSortDomainsAlphabetically Failed!")
		} else {
			t.Logf("TestSortDomainsAlphabetically Passed!")
		}
	}
}

func TestRetreiveDomainName(t *testing.T) {
	fullEmailAddress := "www.legolas@gmail.com"

	output, err := RetrieveDomainName(fullEmailAddress)

	if err != nil {
		t.Errorf("TestRetreiveDomainName Failed!")
	} else if output != "gmail.com" {
		t.Errorf("TestRetreiveDomainName Failed!")
	} else {
		t.Logf("TestRetreiveDomainName Passed!")
	}
}

func TestRetreiveDomainName_ShouldOnlyRetrieveTextAfterTheFinalAtSymbol(t *testing.T) {
	fullEmailAddress := "www.legolas@gmail.com@yahoo.ie"

	output, err := RetrieveDomainName(fullEmailAddress)

	if err != nil {
		t.Errorf("TestRetreiveDomainName Failed!")
	} else if output != "yahoo.ie" {
		t.Errorf("TestRetreiveDomainName Failed!")
	} else {
		t.Logf("TestRetreiveDomainName Passed!")
	}
}

func TestRetreiveDomainName_ShouldReturnErrorWhenEmailIsInvalid(t *testing.T) {
	fullEmailAddress := "email"

	output, err := RetrieveDomainName(fullEmailAddress)

	if output != "" {
		t.Errorf("TestRetreiveDomainName Failed!")
	} else if err == nil {
		t.Errorf("TestRetreiveDomainName Failed!")
	} else if err.Error() != "invalid domain name" {
		t.Errorf("TestRetreiveDomainName Failed!")
	} else {
		t.Logf("TestRetreiveDomainName Passed!")
	}
}
