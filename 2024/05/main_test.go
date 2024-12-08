package main

import (
	"reflect"
	"testing"
)

var exampleData = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

type PageNumbersTest struct {
	pageNumbers PageNumbers
	isValid     bool
}

func TestParseInput(t *testing.T) {
	expectedRules := Rules{
		47: []int{53, 13, 61, 29},
		97: []int{13, 61, 47, 29, 53, 75},
		75: []int{29, 53, 47, 61, 13},
		61: []int{13, 53, 29},
		29: []int{13},
		53: []int{29, 13},
	}
	expectedPageNumbers := []PageNumbers{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	gotOrderingRules, gotPageUpdates, err := parseInput(exampleData)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(gotOrderingRules, expectedRules) {
		t.Errorf("got: %v, expected %v", gotOrderingRules, expectedRules)
	}

	if !reflect.DeepEqual(gotPageUpdates, expectedPageNumbers) {
		t.Errorf("got: %v, expected %v", gotPageUpdates, expectedPageNumbers)
	}
}

func TestIsValid(t *testing.T) {
	rules := Rules{
		47: []int{53, 13, 61, 29},
		97: []int{13, 61, 47, 29, 53, 75},
		75: []int{29, 53, 47, 61, 13},
		61: []int{13, 53, 29},
		29: []int{13},
		53: []int{29, 13},
	}

	pageNumbersTests := []PageNumbersTest{
		{
			pageNumbers: PageNumbers{75, 47, 61, 53, 29},
			isValid:     true,
		},
		{
			pageNumbers: PageNumbers{97, 61, 53, 29, 13},
			isValid:     true,
		},
		{
			pageNumbers: PageNumbers{75, 29, 13},
			isValid:     true,
		},
		{
			pageNumbers: PageNumbers{75, 97, 47, 61, 53},
			isValid:     false,
		},
		{
			pageNumbers: PageNumbers{61, 13, 29},
			isValid:     false,
		},
		{
			pageNumbers: PageNumbers{97, 13, 75, 29, 47},
			isValid:     false,
		},
	}

	for _, pageNumbersTest := range pageNumbersTests {
		isValid := isValid(rules, pageNumbersTest.pageNumbers)
		if isValid != pageNumbersTest.isValid {
			t.Errorf("got: %v, expected %v", isValid, pageNumbersTest.isValid)
		}
	}
}

func TestSumValidPageUpdate(t *testing.T) {
	rules := Rules{
		47: []int{53, 13, 61, 29},
		97: []int{13, 61, 47, 29, 53, 75},
		75: []int{29, 53, 47, 61, 13},
		61: []int{13, 53, 29},
		29: []int{13},
		53: []int{29, 13},
	}

	pageUpdates := []PageNumbers{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	expect := 143

	got := sumValidPageUpdates(rules, pageUpdates)
	if got != expect {
		t.Errorf("got: %v, expect %v", got, expect)
	}
}
