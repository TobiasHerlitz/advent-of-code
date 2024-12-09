package main

import (
	"testing"
)

func TestFindRobot(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	}

	expectX := 4
	expectY := 6 // Note that origin is in top left corner
	gotX, gotY, err := findRobot(grid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if gotX != expectX || gotY != expectY {
		t.Errorf("Expected to find robot at [%v:%v]. Got: [%v:%v]", expectX, expectY, gotX, gotY)
	}
}

func TestPartOne(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	}

	expect := 41

	got, err := partOne(grid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if got != expect {
		t.Errorf("got: %v, expect %v", got, expect)
	}
}

func TestPartTwo(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	}

	expect := 6

	got, err := partTwo(grid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if got != expect {
		t.Errorf("got: %v, expect %v", got, expect)
	}
}

func TestIsRetreadingReturnsFalse(t *testing.T) {
	breadcrumbs := []coordinate{
		{x: 1, y: 1},
		{x: 1, y: 2},
		{x: 1, y: 3},
		{x: 1, y: 4},
	}

	got := isRetreading(&breadcrumbs)
	expect := false

	if got != expect {
		t.Errorf("got: %v, expect %v", got, expect)
	}
}

func TestIsRetreadingReturnsTrue(t *testing.T) {
	breadcrumbs := []coordinate{
		{x: 1, y: 1},
		{x: 1, y: 2},
		{x: 1, y: 3},
		{x: 1, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 3},
		{x: 1, y: 4},
	}

	got := isRetreading(&breadcrumbs)
	expect := true

	if got != expect {
		t.Errorf("got: %v, expect %v", got, expect)
	}
}

func TestPredictPathBreaksOutOfLoop(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'#', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	}

	_, err := predictPath(grid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
