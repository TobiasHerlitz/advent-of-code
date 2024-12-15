package adventhelper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid [][]int
type Coordinate struct {
	X int
	Y int
}

// See if this can be generic. i.e. also return [][]rune
func LoadGrid(filePath string) (Grid, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	grid := [][]int{}
	for _, rowInput := range strings.Split(string(input), "\n") {
		row := []int{}
		for _, cellInput := range rowInput {
			cell, err := strconv.Atoi(string(cellInput))
			if err != nil {
				return nil, err
			}

			row = append(row, cell)
		}

		grid = append(grid, row)
	}

	return grid, nil
}

func (g Grid) ToString() string {
	rowRepresentations := []string{}
	for _, row := range g {
		rowRepresentation := ""
		for _, cell := range row {
			rowRepresentation = fmt.Sprintf("%v%v", rowRepresentation, cell)
		}
		rowRepresentations = append(rowRepresentations, rowRepresentation)
	}

	return strings.Join(rowRepresentations, "\n")
}

func (g Grid) IsWithinBounds(coord Coordinate) bool {
	return coord.Y >= 0 && coord.Y < len(g) && coord.X >= 0 && coord.X < len(g[coord.Y])
}

func (g Grid) ReadCell(coord Coordinate) (int, error) {
	if !g.IsWithinBounds(coord) {
		return 0, fmt.Errorf("coordinate [%d:%d] is out of bounds", coord.X, coord.Y)
	}

	return g[coord.Y][coord.X], nil
}
