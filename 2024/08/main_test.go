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
	got := countAntinodes(exampleGrid, false)
	expect := 14

	if got != expect {
		t.Errorf("Wrong number of antinodes. Got: %v. Expect: %v", got, expect)
	}
}

func TestCountAntinodesWithEcho(t *testing.T) {
	got := countAntinodes(exampleGrid, true)
	expect := 34

	if got != expect {
		t.Errorf("Wrong number of antinodes with echo. Got: %v. Expect: %v", got, expect)
	}
}
