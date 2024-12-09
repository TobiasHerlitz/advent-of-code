package main

import (
	"math"
	"strconv"
)

// Old solution without recursion
func (e equation) oldIsValid() bool {
	operatorPermutations := 1 << len(e.operands)
	for i := 0; i < operatorPermutations; i++ {
		var total int64 = int64(e.operands[0])

		for operandIndex, operand := range e.operands[1:] {
			// https://www.algotree.org/algorithms/numeric/subsets_bitwise/
			if (i & (1 << operandIndex)) != 0 {
				total += int64(operand)
			} else {
				total *= int64(operand)
			}

			if total == e.answer {
				return true
			}

			if total > e.answer {
				break
			}
		}
	}

	return false
}

func (e equation) oldIsValidWithConcatenation() (bool, error) {
	operatorPermutations := int(math.Pow(3, float64((len(e.operands) - 1))))
	for i := 0; i < operatorPermutations; i++ {
		var total int64 = int64(e.operands[0])

		num := i
		for _, operand := range e.operands[1:] {
			variation := num % 3
			switch variation {
			case 0:
				total += int64(operand)
			case 1:
				total *= int64(operand)
			case 2:
				total = total*int64(math.Pow10(len(strconv.FormatInt(operand, 10)))) + operand
			}

			if total == e.answer {
				return true, nil
			}

			if total > e.answer {
				break
			}
			num /= 3
		}
	}

	return false, nil
}
