package main

import (
	"reflect"
	"testing"
)

func TestCompact(t *testing.T) {
	exampleMemory := memory{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9}
	exampleMemory.compact()

	expect := memory{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	if !reflect.DeepEqual(exampleMemory, expect) {
		t.Errorf("Incorrectly compacted. got: %v, expect %v", exampleMemory, expect)
	}
}

func TestCompactWithIntegrity(t *testing.T) {
	exampleMemory := memory{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9}
	exampleMemory.compactWithIntegrity()

	expect := memory{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1}

	if !reflect.DeepEqual(exampleMemory, expect) {
		t.Errorf("Incorrectly compacted. got: %v, expect %v", exampleMemory, expect)
	}
}

func TestGetChecksumForPartOne(t *testing.T) {
	compactedMemory := memory{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	got := compactedMemory.getChecksum()
	expect := 1928

	if got != expect {
		t.Errorf("Incorrect checksum. got: %v, expect %v", got, expect)
	}
}

func TestGetChecksumForPartTwo(t *testing.T) {
	compactedMemory := memory{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1}
	got := compactedMemory.getChecksum()
	expect := 2858

	if got != expect {
		t.Errorf("Incorrect checksum. got: %v, expect %v", got, expect)
	}
}

func TestSwapSections(t *testing.T) {
	tests := []struct {
		memory          memory
		expect          memory
		sectionOneIndex int
		sectionTwoIndex int
		length          int
	}{
		{
			memory:          memory{6, 6, 6, 6, -1, -1, 7, 7, -1, 8, 8, 8, 8, -1, -1},
			expect:          memory{6, 6, 6, 6, 7, 7, -1, -1, -1, 8, 8, 8, 8, -1, -1},
			sectionOneIndex: 4,
			sectionTwoIndex: 6,
			length:          2,
		},
		{
			memory:          memory{-1, -1, -1, -1, 7, 7, 7, 7, -1, -1, 8, 8, 8, 8, -1},
			expect:          memory{8, 8, 8, 8, 7, 7, 7, 7, -1, -1, -1, -1, -1, -1, -1},
			sectionOneIndex: 0,
			sectionTwoIndex: 10,
			length:          4,
		},
	}

	for _, testCase := range tests {
		testCase.memory.swapSections(testCase.sectionOneIndex, testCase.sectionTwoIndex, testCase.length)
		if !reflect.DeepEqual(testCase.memory, testCase.expect) {
			t.Errorf("Incorrectly swapped. got: %v, expect %v", testCase.memory, testCase.expect)
		}
	}
}

func TestFindEmptySectionIndex(t *testing.T) {
	tests := []struct {
		memory  memory
		expect  int
		minSize int
	}{
		{
			memory:  memory{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
			expect:  2,
			minSize: 1,
		},
		{
			memory:  memory{0, 0, 9, 9, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, -1, -1},
			expect:  8,
			minSize: 3,
		},
		{
			memory:  memory{0, 0, 9, 9, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 8, -1, -1},
			expect:  -1,
			minSize: 3,
		},
		{
			memory:  memory{0, 0, 9, 9, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 8, -1, -1},
			expect:  40,
			minSize: 2,
		},
		{
			memory:  memory{0, 0, 9, 9, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, -1, -1},
			expect:  -1,
			minSize: 4,
		},
	}

	for _, testCase := range tests {
		got := testCase.memory.findEmptySectionIndex(testCase.minSize)
		if got != testCase.expect {
			t.Errorf("Failed finding the correct index on: %v. got: %v, expect %v", testCase.memory, got, testCase.expect)
		}
	}
}

func TestFindRightmostFileIndex(t *testing.T) {
	tests := []struct {
		memory      memory
		from        int
		expectSize  int
		expectIndex int
	}{
		{
			memory:      memory{6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, -1, -1},
			from:        14,
			expectIndex: 9,
			expectSize:  4,
		},
		{
			memory:      memory{6, 6, 6, -1, 7, 7, 7, 7, -1, -1, 8, 8, -1, -1, -1},
			from:        14,
			expectIndex: 10,
			expectSize:  2,
		},
		{
			memory:      memory{-1, -1, 6, 6, 6, 6, -1, -1, 7, 7, 7, -1, -1, 9, 9},
			from:        14,
			expectIndex: 13,
			expectSize:  2,
		},
		{
			memory:      memory{-1, -1, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, -1, -1, 9},
			from:        14,
			expectIndex: 14,
			expectSize:  1,
		},
		{
			memory:      memory{6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, -1, -1},
			from:        3,
			expectIndex: 0,
			expectSize:  4,
		},
		{
			memory:      memory{6, 6, 6, -1, 7, 7, 7, 7, -1, -1, 8, 8, -1, -1, -1},
			from:        8,
			expectIndex: 4,
			expectSize:  4,
		},
		{
			memory:      memory{-1, -1, 6, -1, 7, 7, 7, 7, -1, -1, 8, 8, -1, -1, -1},
			from:        1,
			expectIndex: -1,
			expectSize:  -1,
		},
	}

	for _, testCase := range tests {
		gotIndex, gotSize := testCase.memory.findRightmostFileIndex(testCase.from)
		if gotIndex != testCase.expectIndex {
			t.Errorf("Failed finding the correct index on: %v. got: %v, expect %v", testCase.memory, gotIndex, testCase.expectIndex)
		}
		if gotSize != testCase.expectSize {
			t.Errorf("Failed finding the correct size on: %v. got: %v, expect %v", testCase.memory, gotSize, testCase.expectSize)
		}
	}
}
