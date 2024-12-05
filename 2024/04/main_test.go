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

func TestCountXShapedMasOccurences(t *testing.T) {
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

	expected := 9

	got := countXShapedMasOccurences(exampleData)

	if got != expected {
		t.Errorf("got: %v, expected %v", got, expected)
	}
}

func TestIsValidPairWithValidInput(t *testing.T) {
	expected := true

	got := isValidPair('M', 'S')

	if got != expected {
		t.Errorf("got: %v, expected %v", got, expected)
	}
}

func TestIsValidPairWithSameChars(t *testing.T) {
	expected := false

	got := isValidPair('M', 'M')

	if got != expected {
		t.Errorf("got: %v, expected %v", got, expected)
	}
}

func TestIsValidPairWithWrongChar(t *testing.T) {
	expected := false

	got := isValidPair('X', 'M')

	if got != expected {
		t.Errorf("got: %v, expected %v", got, expected)
	}
}
