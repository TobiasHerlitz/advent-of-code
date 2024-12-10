package main

import (
	"bytes"
	"os"
)

func loadGrid() ([][]rune, error) {
	input, err := os.ReadFile("input")
	if err != nil {
		return nil, err
	}

	gridInput := bytes.Split(input, []byte("\n"))

	var grid [][]rune
	for i := 0; i < len(gridInput); i++ {
		row := []rune(string(gridInput[i]))
		grid = append(grid, row)
	}

	return grid, nil
}
