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

type RectifyPageNumberTest struct {
	pageNumbers PageNumbers
	expect      PageNumbers
}

func TestParseInput(t *testing.T) {
	expectRules := Rules{
		47: []int{53, 13, 61, 29},
		97: []int{13, 61, 47, 29, 53, 75},
		75: []int{29, 53, 47, 61, 13},
		61: []int{13, 53, 29},
		29: []int{13},
		53: []int{29, 13},
	}
	expectPageNumbers := []PageNumbers{
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

	if !reflect.DeepEqual(gotOrderingRules, expectRules) {
		t.Errorf("got: %v, expect %v", gotOrderingRules, expectRules)
	}

	if !reflect.DeepEqual(gotPageUpdates, expectPageNumbers) {
		t.Errorf("got: %v, expect %v", gotPageUpdates, expectPageNumbers)
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

	testData := []PageNumbersTest{
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

	for _, testCase := range testData {
		isValid := isValid(rules, testCase.pageNumbers)
		if isValid != testCase.isValid {
			t.Errorf("got: %v, expect %v", isValid, testCase.isValid)
		}
	}
}

func TestSumMiddlePageNumbers(t *testing.T) {
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

	gotSumFromValid, _ := sumMiddlePageNumbers(rules, pageUpdates)
	if gotSumFromValid != expect {
		t.Errorf("got: %v, expect %v", gotSumFromValid, expect)
	}
}

func TestToRectifiedPageNumbers(t *testing.T) {
	rules := Rules{
		47: []int{53, 13, 61, 29},
		97: []int{13, 61, 47, 29, 53, 75},
		75: []int{29, 53, 47, 61, 13},
		61: []int{13, 53, 29},
		29: []int{13},
		53: []int{29, 13},
	}

	testData := []RectifyPageNumberTest{
		{
			pageNumbers: PageNumbers{75, 97, 47, 61, 53},
			expect:      PageNumbers{97, 75, 47, 61, 53},
		},
		{
			pageNumbers: PageNumbers{61, 13, 29},
			expect:      PageNumbers{61, 29, 13},
		},
		{
			pageNumbers: PageNumbers{97, 13, 75, 29, 47},
			expect:      PageNumbers{97, 75, 47, 29, 13},
		},
	}

	for _, testCase := range testData {
		got := toRectifiedPageNumbers(rules, testCase.pageNumbers)
		if !reflect.DeepEqual(got, testCase.expect) {
			t.Errorf("Incorrectly rectified. got: %v, expect %v", got, testCase.expect)
		}
	}
}
