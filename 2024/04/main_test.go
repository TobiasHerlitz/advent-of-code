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
	exampleTerm := []byte("XMAS")
	expected := 18

	got := countOccurrences(exampleData, exampleTerm)

	if got != expected {
		t.Errorf("got: %v, expected %v", got, expected)
	}
}
