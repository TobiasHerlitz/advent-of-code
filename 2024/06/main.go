package main

import (
	"bytes"
	"errors"
	"fmt"
	"maps"
	"os"
	"slices"
)

type coordinateShift struct {
	x int
	y int
}

type coordinate struct {
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

func cloneGrid(grid [][]rune) [][]rune {
	var newGrid [][]rune
	for _, column := range grid {
		newGrid = append(newGrid, slices.Clone(column))
	}

	return newGrid
}

func readCell(grid [][]rune, x int, y int) (rune, error) {
	if x < 0 || x > len(grid)-1 || y < 0 || y > len(grid[x])-1 {
		return 0, fmt.Errorf("coordinate [%d:%d] is out of bounds", x, y)
	}

	return grid[y][x], nil
}

// isRetreading if the last two breadcrumbs can be found earlier in breadcrumbs
func isRetreading(breadcrumbs *[]coordinate) bool {
	length := len(*breadcrumbs)
	if length < 5 {
		return false
	}

	secondToLastLocation := (*breadcrumbs)[length - 2]
	lastLocation := (*breadcrumbs)[length - 1]
	for index, breadcrumb := range (*breadcrumbs)[:length - 2] {
		if breadcrumb == secondToLastLocation && (*breadcrumbs)[index + 1] == lastLocation {
			return true
		}
	}

	return false
}

func walkRobot(grid [][]rune, x int, y int, breadcrumbs *[]coordinate) {
	robot := grid[y][x]
	robotDirection := robotDirections[robot]
	nextX := x + robotDirection.x
	nextY := y + robotDirection.y

	if isRetreading(breadcrumbs) {
		return
	}

	nextCell, ok := readCell(grid, nextX, nextY)
	if ok != nil {
		grid[y][x] = '.'
		*breadcrumbs = append(*breadcrumbs, coordinate{x, y})
		return
	}

	if nextCell == '#' {
		grid[y][x] = rightTurns[robot]
		walkRobot(grid, x, y, breadcrumbs)
		return
	}

	if nextCell == '.' {
		*breadcrumbs = append(*breadcrumbs, coordinate{x, y})
		grid[y][x] = '.'
		grid[nextY][nextX] = robot
		walkRobot(grid, nextX, nextY, breadcrumbs)
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

	return 0, 0, errors.New("Couldn't find robot")
}

func predictPath(grid [][]rune) ([]coordinate, error) {
	x, y, err := findRobot(grid)
	if err != nil {
		return nil, err
	}

	var breadcrumbs []coordinate
	walkRobot(grid, x, y, &breadcrumbs)
	return breadcrumbs, nil
}

func getDistinctLocations(breadcrumbs []coordinate) []coordinate {
	distinctLocations := make(map[string]coordinate)
	for _, breadcrumb := range breadcrumbs {
		distinctLocations[fmt.Sprintf("%d:%d", breadcrumb.x, breadcrumb.y)] = breadcrumb
	}

	return slices.Collect(maps.Values(distinctLocations))
}

func partOne(grid [][]rune) (int, error) {
	breadcrumbs, err := predictPath(grid)
	if err != nil {
		return 0, err
	}
	return len(getDistinctLocations(breadcrumbs)), nil
}

// Could be more efficient
func partTwo(grid [][]rune) (int, error) {
	breadcrumbs, err := predictPath(cloneGrid(grid))
	if err != nil {
		return 0, err
	}

	robotX, robotY, err := findRobot(grid);
	if err != nil {
		return 0, err
	}

	loopSolutions := 0
	for _, distinctLocation := range getDistinctLocations(breadcrumbs) {
		if distinctLocation.x == robotX && distinctLocation.y == robotY {
			continue
		}
		alteredGrid := cloneGrid(grid)
		alteredGrid[distinctLocation.y][distinctLocation.x] = '#'
		alteredBreadcrumbs, err := predictPath(alteredGrid)
		if err != nil {
			return 0, err
		}

		if isRetreading(&alteredBreadcrumbs) {
			loopSolutions++
		}
	}

	return loopSolutions, nil
}

func main() {
	grid, err := loadGrid()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed loading grid:", err)
		os.Exit(1)		
	}

	distinctLocations, err := partOne(cloneGrid(grid))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed getting distinct locations:", err)
		os.Exit(1)
	}

	possibleObstacleLocations, err := partTwo(cloneGrid(grid))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed finding possible obstacle locations:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - The robot has passed: %d cells\n", distinctLocations) // 4515
	fmt.Printf("Part 2 - Number of valid obstacle placements: %d\n", possibleObstacleLocations) // 1309
}
