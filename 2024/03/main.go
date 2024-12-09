package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	entryRegex           = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	termsRegex           = regexp.MustCompile(`[0-9]{1,3}`)
	activeSubStringRegex = regexp.MustCompile(`(do\(\)|^)[\s\S]*?(don't\(\)|$)`)
)

func calculateEntry(entry []byte) (int, error) {
	terms := termsRegex.FindAll(entry, -1)

	if len(terms) != 2 {
		return 0, fmt.Errorf("Expected two terms from entry: %s", entry)
	}

	termOne, err := strconv.Atoi(string(terms[0]))
	if err != nil {
		return 0, err
	}

	termTwo, err := strconv.Atoi(string(terms[1]))
	if err != nil {
		return 0, err
	}

	return termOne * termTwo, nil
}

func sumValidEntries(input []byte) (int, error) {
	validEntries := entryRegex.FindAll(input, -1)

	var total int
	for _, entry := range validEntries {
		product, err := calculateEntry(entry)
		if err != nil {
			return 0, err
		}
		total += product
	}

	return total, nil
}

func sumOnlyActiveEntries(input []byte) (int, error) {
	activeSubStrings := activeSubStringRegex.FindAll(input, -1)

	var total int
	for _, activeSubString := range activeSubStrings {
		sum, err := sumValidEntries(activeSubString)
		if err != nil {
			return 0, err
		}
		total += sum
	}

	return total, nil
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	partOneResponse, err := sumValidEntries(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error:", err)
		os.Exit(1)
	}

	partTwoResponse, err := sumOnlyActiveEntries(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - Total of all valid entries: %d\n", partOneResponse)
	fmt.Printf("Part 2 - Total of all valid and active entries: %d\n", partTwoResponse)
}
