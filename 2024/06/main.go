package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

type coordinateShift struct {
	x int
	y int
}

var robotDirections = map[rune]coordinateShift{
	'^': {x: 0, y: -1},
	'>': {x: 1, y: 0},
	'v': {x: 0, y: 1},
	'<': {x: -1, y: 0},
}

var rightTurns = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

func readCell(grid [][]rune, x int, y int) (rune, error) {
	if x < 0 || x > len(grid)-1 || y < 0 || y > len(grid[x])-1 {
		return 0, fmt.Errorf("coordinate [%d:%d] is out of bounds", x, y)
	}

	return grid[y][x], nil
}

func parseInput(input []byte) [][]rune {
	gridInput := bytes.Split(input, []byte("\n"))

	var grid [][]rune
	for i := 0; i < len(gridInput); i++ {
		row := []rune(string(gridInput[i]))
		grid = append(grid, row)
	}

	return grid
}

func walkRobot(grid [][]rune, x int, y int) {
	robot := grid[y][x]
	robotDirection := robotDirections[robot]
	nextX := x + robotDirection.x
	nextY := y + robotDirection.y

	nextCell, ok := readCell(grid, nextX, nextY)
	if ok != nil {
		grid[y][x] = 'X'
		return
	}

	if nextCell == '#' {
		grid[y][x] = rightTurns[robot]
		walkRobot(grid, x, y)
		return
	}

	if nextCell == '.' || nextCell == 'X' {
		grid[y][x] = 'X'
		grid[nextY][nextX] = robot
		walkRobot(grid, nextX, nextY)
		return
	}
}

func findRobot(grid [][]rune) (int, int, error) {
	for y, column := range grid {
		for x, cell := range column {
			_, found := robotDirections[cell]
			if found {
				return x, y, nil
			}
		}
	}

	return 0, 0, errors.New("Couldnt find robot")
}

func predictPath(grid [][]rune) ([][]rune, error) {
	x, y, err := findRobot(grid)
	if err != nil {
		return nil, err
	}

	walkRobot(grid, x, y)
	return grid, nil
}

func countPassedCells(grid [][]rune) int {
	total := 0
	for _, column := range grid {
		for _, cell := range column {
			if cell == 'X' {
				total++
			}
		}
	}

	return total
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed reading input:", err)
		os.Exit(1)
	}

	grid := parseInput(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed parsing input:", err)
		os.Exit(1)
	}

	paintedGrid, err := predictPath(grid)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed painting grid:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - The robot has passed: %d cells\n", countPassedCells(paintedGrid))
}
