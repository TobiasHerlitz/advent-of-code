package main

import (
	"testing"
)

func intPtr(i int) *int {
	return &i
}

func TestBxl(t *testing.T) {
	registers := Registers{
		A: 0,
		B: 29,
		C: 0,
	}

	expect := 26

	registers.bxl(7)
	if registers.B != expect {
		t.Errorf("Wrong value in register B. got: %v, expect %v", registers.B, expect)
	}
}

func TestBst(t *testing.T) {
	registers := Registers{
		A: 0,
		B: 0,
		C: 9,
	}

	expect := 1

	registers.bst(6)
	if registers.B != expect {
		t.Errorf("Wrong value in register B. got: %v, expect %v", registers.B, expect)
	}
}

func TestBxc(t *testing.T) {
	registers := Registers{
		A: 0,
		B: 2024,
		C: 43690,
	}

	expectRegisterB := 44354

	registers.bxc()
	if registers.B != expectRegisterB {
		t.Errorf("Wrong value in register B. got: %v, expect %v", registers.B, expectRegisterB)
	}
}

func TestRunProgram(t *testing.T) {
	tests := []struct {
		registers       Registers
		program         []int
		expectOutput    string
		expectRegisterA *int
	}{
		{
			registers:       Registers{A: 10, B: 0, C: 0},
			program:         []int{5, 0, 5, 1, 5, 4},
			expectOutput:    "0,1,2",
			expectRegisterA: nil,
		},
		{
			registers:       Registers{A: 2024, B: 0, C: 0},
			program:         []int{0, 1, 5, 4, 3, 0},
			expectOutput:    "4,2,5,6,7,7,7,7,3,1,0",
			expectRegisterA: intPtr(0),
		},
		{
			registers:       Registers{A: 729, B: 0, C: 0},
			program:         []int{0, 1, 5, 4, 3, 0},
			expectOutput:    "4,6,3,5,6,3,5,2,1,0",
			expectRegisterA: nil,
		},
	}

	for _, testCase := range tests {
		got, err := runProgram(testCase.program, &testCase.registers)
		if err != nil {
			t.Errorf("Unexpected error when testing runProgram. original: %v", err)
		}

		if testCase.expectRegisterA != nil && testCase.registers.A != *testCase.expectRegisterA {
			t.Errorf("Wrong value in register A. got: %v, expect %v", testCase.registers.A, *testCase.expectRegisterA)
		}

		if got != testCase.expectOutput {
			t.Errorf("Incorrect output from runProgram. got: %v, expect %v", got, testCase.expectOutput)
		}
	}
}
