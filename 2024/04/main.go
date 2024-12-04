package main

import (
	"bytes"
	"fmt"
	"os"
)

func countOccurences(characterGrid [][]byte, term string) (int, error) {
	return 22, nil
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	characterGrid := bytes.Split(input, []byte("\n"))

	term := "XMAS"
	occurences, err := countOccurences(characterGrid, term)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - Times 'XMAS' occurs: %d\n", occurences)
	// fmt.Printf("Part 2 - Total of all valid and active entries: %d\n", )
}
