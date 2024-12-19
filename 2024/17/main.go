package main

import (
	"fmt"
	"os"
	"strings"
)

type Registers struct {
	A int
	B int
	C int
}

type RegisterID string

const (
	registerA RegisterID = "A"
	registerB RegisterID = "B"
	registerC RegisterID = "C"
)

var registers = Registers{
	A: 50230824,
	B: 0,
	C: 0,
}

// var program = []int{2, 4, 1, 3, 7, 5, 0, 3, 1, 4, 4, 7, 5, 5, 3, 0}
var program = []int{2, 4, 1, 3, 7, 5, 0, 3, 1, 4, 4, 7, 5, 5, 3, 0}

// 1. Store registerA value in registerB
// 2. Bitwise XOR on RegisterB and 3, store in RegisterB
// 3. Store registerB value in registerC
// 4. Divide registerA by 8
// 5. Store registerA value in registerB
// 6. Bitwise XOR on registerB and registerC, store in registerB
// 7. Output mod 8 of registerB
// 8. Restart until registerA is 0

// 2. Bitwise XOR on RegisterA and 3, store in RegisterB
// 4. Divide registerA by 8
// 6. Bitwise XOR on registerA and registerB, store in registerB
// 7. Output mod 8 of registerB
// 8. Restart until registerA is 0

// Opcodes 0, 6, and 7.
func (r *Registers) division(writeTarget RegisterID, operand int) {
	quotient := r.A >> r.getComboOperand(operand)

	switch writeTarget {
	case registerA:
		r.A = quotient
	case registerB:
		r.B = quotient
	case registerC:
		r.C = quotient
	}
}

// Opcode 1
func (r *Registers) bxl(operand int) {
	r.B = r.B ^ operand
}

// Opcode 2
func (r *Registers) bst(operand int) {
	r.B = r.getComboOperand(operand) % 8
}

// Opcode 3
func (r *Registers) jnz(operand int) (bool, int) {
	if r.A == 0 {
		return false, 0
	}

	return true, operand
}

// Opcode 4
func (r *Registers) bxc() {
	r.B = r.B ^ r.C
}

// Opcode 5
func (r *Registers) out(operand int) string {
	return fmt.Sprintf("%v", r.getComboOperand(operand)%8)
}

func (r *Registers) getComboOperand(operandValue int) int {
	if operandValue == 4 {
		return r.A
	}

	if operandValue == 5 {
		return r.B
	}

	if operandValue == 6 {
		return r.C
	}

	return operandValue
}

func runProgram(program []int, registers *Registers) (string, error) {
	var outputs []string
	for programPointer := 0; programPointer < len(program); {
		operand := program[programPointer+1]
		switch program[programPointer] {
		case 0:
			registers.division(registerA, operand)
		case 1:
			registers.bxl(operand)
		case 2:
			registers.bst(operand)
		case 3:
			jump, newProgramPointer := registers.jnz(operand)
			if jump {
				programPointer = newProgramPointer
				continue
			}
		case 4:
			registers.bxc()
		case 5:
			outputs = append(outputs, registers.out(operand))
		case 6:
			registers.division(registerB, operand)
		case 7:
			registers.division(registerC, operand)
		default:
			return "", fmt.Errorf("Failed matching instruction to handler. Got: %v", program[programPointer])
		}

		programPointer += 2
	}

	return strings.Join(outputs, ","), nil
}

// Thoughts:
// Out should execute 16 times
// The program has to end with [3, x]. Right?
// The program is going to "bounce" on the final [3, x] until register a is 0
// A is only written to on [0, x]
// The only time opcode 0 is used is [0, 3]. i.e. a divided by 8 (2^3)
// 
// Guess: Given that A is divided by 8, 16 times to produce a value that lets the program "escape"
// So: A / 8^16 < 1
// A < 8^16
// 
// And given that it needs to "bounce" 16 times.
// So: A / 8^15 > 1
// A > 8^15
// 
// 35 184 372 088 832 < a < 281 474 976 710 656
// Yeeee, thats not reasonable
func searchForLowestA(program []int) (int, error) {

	// return 0, fmt.Errorf("Failed finding lowest A")
	return 0, nil
}

func main() {
	output, err := runProgram(program, &registers)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed running program:", err)
		os.Exit(1)
	}

	lowestA, err := searchForLowestA(program)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed searching for lowest A:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - Output of program is: %v\n", output)
	fmt.Printf("Part 2 - Lowest A that outputs an identical program is: %v\n", lowestA)
}
