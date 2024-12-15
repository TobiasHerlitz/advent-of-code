package main

import (
	"testing"

	helper "github.com/TobiasHerlitz/advent-of-code/shared/go/adventhelper"
)

func TestFindStartingPoints(t *testing.T) {
	exampleGrid := helper.Grid{
		[]int{8, 9, 0, 1, 0, 1, 2, 3},
		[]int{7, 8, 1, 2, 1, 8, 7, 4},
		[]int{8, 7, 4, 3, 0, 9, 6, 5},
		[]int{9, 6, 5, 4, 9, 8, 7, 4},
		[]int{4, 5, 6, 7, 8, 9, 0, 3},
		[]int{3, 2, 0, 1, 9, 0, 1, 2},
		[]int{0, 1, 3, 2, 9, 8, 0, 1},
		[]int{1, 0, 4, 5, 6, 7, 3, 2},
	}

	expect := 9
	got := len(findStartingPoints(exampleGrid))

	if got != expect {
		t.Errorf("Incorrect number of starting points. got: %v, expect %v. On:\n%v\n", got, expect, exampleGrid.ToString())
	}
}

func TestWalkTrails(t *testing.T) {
	exampleGrid := helper.Grid{
		[]int{8, 9, 0, 1, 0, 1, 2, 3},
		[]int{7, 8, 1, 2, 1, 8, 7, 4},
		[]int{8, 7, 4, 3, 0, 9, 6, 5},
		[]int{9, 6, 5, 4, 9, 8, 7, 4},
		[]int{4, 5, 6, 7, 8, 9, 0, 3},
		[]int{3, 2, 0, 1, 9, 0, 1, 2},
		[]int{0, 1, 3, 2, 9, 8, 0, 1},
		[]int{1, 0, 4, 5, 6, 7, 3, 2},
	}

	tests := []struct {
		startingPoint helper.Coordinate
		expect        int
	}{
		{
			startingPoint: helper.Coordinate{X: 2, Y: 0},
			expect:        5,
		},
		{
			startingPoint: helper.Coordinate{X: 4, Y: 0},
			expect:        6,
		},
		{
			startingPoint: helper.Coordinate{X: 4, Y: 2},
			expect:        5,
		},
		{
			startingPoint: helper.Coordinate{X: 6, Y: 4},
			expect:        3,
		},
		{
			startingPoint: helper.Coordinate{X: 2, Y: 5},
			expect:        1,
		},
		{
			startingPoint: helper.Coordinate{X: 5, Y: 5},
			expect:        3,
		},
		{
			startingPoint: helper.Coordinate{X: 0, Y: 6},
			expect:        5,
		},
		{
			startingPoint: helper.Coordinate{X: 6, Y: 6},
			expect:        3,
		},
		{
			startingPoint: helper.Coordinate{X: 1, Y: 7},
			expect:        5,
		},
	}

	for _, testCase := range tests {
		uniqueEndpoints := make(map[helper.Coordinate]struct{})
		walkTrails(exampleGrid, testCase.startingPoint, uniqueEndpoints)
		got := len(uniqueEndpoints)
		if got != testCase.expect {
			t.Errorf("Incorrect number of trails found. got: %v, expect %v", got, testCase.expect)
		}
	}
}
