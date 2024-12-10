package adventhelper

import "fmt"

func PrintStuff() {
	fmt.Println("Stuff")
}

// func loadEquations() ([]equation, error) {
// 	input, err := os.ReadFile("input")
// 	if err != nil {
// 		return nil, err
// 	}

// 	var equations []equation
// 	for _, equationRow := range strings.Split(string(input), "\n") {
// 		equationParts := strings.Split(equationRow, ": ")

// 		answer, err := strconv.ParseInt(equationParts[0], 10, 64)
// 		if err != nil {
// 			return nil, err
// 		}

// 		var operands []int64
// 		for _, operandInput := range strings.Split(equationParts[1], " ") {
// 			operand, err := strconv.Atoi(operandInput)
// 			if err != nil {
// 				return nil, err
// 			}
// 			operands = append(operands, int64(operand))
// 		}

// 		equations = append(equations, equation{
// 			answer,
// 			operands,
// 		})
// 	}

// 	return equations, nil
// }
