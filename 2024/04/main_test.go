package main

import "testing"

func TestCountOccurences(t *testing.T) {
	exampleData := [][]byte{
		[]byte("MMMSXXMASM"),
		[]byte("MSAMXMSMSA"),
		[]byte("AMXSXMAAMM"),
		[]byte("MSAMASMSMX"),
		[]byte("XMASAMXAMM"),
		[]byte("XXAMMXXAMA"),
		[]byte("SMSMSASXSS"),
		[]byte("SAXAMASAAA"),
		[]byte("MAMMMXMMMM"),
		[]byte("MXMXAXMASX"),
	}
	exampleTerm := "XMAS"
	expected := 18

	got, err := countOccurences(exampleData, exampleTerm)
	if err != nil {
		t.Errorf("Unexpected error in countOccurences(). Original: %v", err)
	}

	if got != expected {
		t.Errorf("got: %v, expected %v", got, expected)
	}
}
