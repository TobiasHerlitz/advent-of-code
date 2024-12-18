// If a program outputs multiple values, they are separated by commas

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

var program = []int{2, 4, 1, 3, 7, 5, 0, 3, 1, 4, 4, 7, 5, 5, 3, 0}

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

func main() {
	output, err := runProgram(program, &registers)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed running program:", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Output of program is: %v\n", output)
}
