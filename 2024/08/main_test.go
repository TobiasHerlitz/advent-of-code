package main

import "testing"

var exampleGrid = [][]rune{
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '0', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '0', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '0', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '0', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', 'A', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', 'A', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'A', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
}

func TestCountAntinodes(t *testing.T) {
	got := countAntinodes(exampleGrid)
	expect := 14

	if got != expect {
		t.Errorf("Wrong number of antinodes. Got: %v. Expect: %v", got, expect)
	}
}