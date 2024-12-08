package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	orderingRuleRegex = regexp.MustCompile(`^([0-9]{1,2}\|[0-9]{1,2})$`)
	pageNumbersRegex  = regexp.MustCompile(`^([0-9]{1,2},)+([0-9]{1,2})$`)
)

type PageNumbers []int
type Rules map[int][]int

func buildPageNumbers(pageNumbersInput string) PageNumbers {
	var pageUpdate PageNumbers
	for _, pageUpdateItem := range strings.Split(pageNumbersInput, ",") {
		num, err := strconv.Atoi(pageUpdateItem)
		if err != nil {
			continue
		}

		pageUpdate = append(pageUpdate, num)
	}
	return pageUpdate
}

func addRule(rules Rules, ruleInput string) error {
	rulePair := strings.Split(ruleInput, "|")
	firstNumber, err := strconv.Atoi(string(rulePair[0]))
	if err != nil {
		return err
	}

	secondNumber, err := strconv.Atoi(string(rulePair[1]))
	if err != nil {
		return err
	}

	_, found := rules[firstNumber]
	if !found {
		rules[firstNumber] = []int{secondNumber}
		return nil
	}

	rules[firstNumber] = append(rules[firstNumber], secondNumber)
	return nil
}

func parseInput(input string) (Rules, []PageNumbers, error) {
	rows := strings.Split(input, "\n")
	var pageUpdates []PageNumbers
	rules := make(map[int][]int)

	for _, row := range rows {
		ruleInput := orderingRuleRegex.FindString(row)
		if ruleInput != "" {
			err := addRule(rules, ruleInput)
			if err != nil {
				return nil, nil, err
			}
		}

		pageNumbersInput := pageNumbersRegex.FindString(row)
		if pageNumbersInput != "" {
			pageUpdates = append(pageUpdates, buildPageNumbers(pageNumbersInput))
		}
	}

	return rules, pageUpdates, nil
}

func isValid(rules Rules, pageNumbers PageNumbers) bool {
	for index, pageNumber := range pageNumbers {
		_, found := rules[pageNumber]
		if !found {
			continue
		}

		for _, previousPageNumber := range pageNumbers[:index] {
			if slices.Contains(rules[pageNumber], previousPageNumber) {
				return false
			}
		}
	}

	return true
}

// Sums the middle page value from each valid page update
func sumValidPageUpdates(rules Rules, pageUpdates []PageNumbers) int {
	total := 0
	for _, pageNumbers := range pageUpdates {
		if isValid(rules, pageNumbers) {
			total += pageNumbers[len(pageNumbers)/2]
		}
	}
	return total
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rules, pageUpdates, err := parseInput(string(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed parsing input:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - Sum of middle page values from each valid page update: %d\n", sumValidPageUpdates(rules, pageUpdates))
}
