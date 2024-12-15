package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type equation struct {
	answer   int64
	operands []int64
}

func calculateOperands(target int64, total int64, operands []int64, allowConcatenation bool) bool {
	if len(operands) == 0 {
		return target == total
	}

	if total > target {
		return false
	}

	if calculateOperands(target, total+operands[0], operands[1:], allowConcatenation) {
		return true
	}

	if calculateOperands(target, total*operands[0], operands[1:], allowConcatenation) {
		return true
	}

	if !allowConcatenation {
		return false
	}

	concat := total*int64(math.Pow10(len(strconv.FormatInt(operands[0], 10)))) + operands[0]
	if calculateOperands(target, concat, operands[1:], true) {
		return true
	}

	return false
}

func (e equation) isValid(allowConcatenation bool) (bool, error) {
	return calculateOperands(e.answer, int64(e.operands[0]), e.operands[1:], allowConcatenation), nil
}

func partOne(equations []equation) (int64, error) {
	var sumOfValidAnswers int64 = 0
	for _, equation := range equations {
		isValid, err := equation.isValid(false)
		if err != nil {
			return 0, err
		}

		if isValid {
			sumOfValidAnswers += equation.answer
		}
	}
	return sumOfValidAnswers, nil
}

func partTwo(equations []equation) (int64, error) {
	var sumOfValidAnswers int64 = 0
	for _, equation := range equations {
		isValid, err := equation.isValid(true)
		if err != nil {
			return 0, err
		}

		if isValid {
			sumOfValidAnswers += equation.answer
		}
	}
	return sumOfValidAnswers, nil
}

func main() {
	equations, err := loadEquations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed loading equations:", err)
		os.Exit(1)
	}

	sumOfValidAnswers, err := partOne(equations)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed solving part one:", err)
		os.Exit(1)
	}

	sumOfValidAnswersWithConcatenation, err := partTwo(equations)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed solving part one:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - Sum of valid answers: %v\n", sumOfValidAnswers)
	fmt.Printf("Part 2 - Sum of valid answers: %v\n", sumOfValidAnswersWithConcatenation)
}
