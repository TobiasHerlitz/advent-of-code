package main

import (
	"testing"
)

var exampleEquations = []equation{
	{
		answer:   int64(190),
		operands: []int64{10, 19},
	},
	{
		answer:   int64(3267),
		operands: []int64{81, 40, 27},
	},
	{
		answer:   int64(83),
		operands: []int64{17, 5},
	},
	{
		answer:   int64(156),
		operands: []int64{15, 6},
	},
	{
		answer:   int64(7290),
		operands: []int64{6, 8, 6, 15},
	},
	{
		answer:   int64(161011),
		operands: []int64{16, 10, 13},
	},
	{
		answer:   int64(192),
		operands: []int64{17, 8, 14},
	},
	{
		answer:   int64(21037),
		operands: []int64{9, 7, 18, 13},
	},
	{
		answer:   int64(292),
		operands: []int64{11, 6, 16, 20},
	},
}

func TestPartOne(t *testing.T) {
	got, err := partOne(exampleEquations)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expect int64 = 3749

	if got != expect {
		t.Errorf("got: %v, expect %v", got, expect)
	}
}

func TestPartTwo(t *testing.T) {
	got, err := partTwo(exampleEquations)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expect int64 = 11387

	if got != expect {
		t.Errorf("got: %v, expect %v", got, expect)
	}
}
