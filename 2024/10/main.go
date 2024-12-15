/**
 * I could do it with recursion.
 * Translating to nodes could be a solution
 */
package main

import (
	"fmt"
	"os"

	helper "github.com/TobiasHerlitz/advent-of-code/shared/go/adventhelper"
)

var Empty struct{}
var directions = []helper.Coordinate{
	{X: 0, Y: -1}, // Up
	{X: 1, Y: 0},  // Right
	{X: 0, Y: 1},  // Down
	{X: -1, Y: 0}, // Left
}

func findStartingPoints(grid helper.Grid) []helper.Coordinate {
	var startingPoints []helper.Coordinate
	for y, row := range grid {
		for x, cell := range row {
			if cell != 0 {
				continue
			}
			startingPoints = append(startingPoints, helper.Coordinate{X: x, Y: y})
		}
	}
	return startingPoints
}

func walkTrails(grid helper.Grid, position helper.Coordinate, uniqueEndpoints map[helper.Coordinate]struct{}) {
	cell, err := grid.ReadCell(position)
	if err != nil {
		return
	}

	if cell == 9 {
		uniqueEndpoints[position] = Empty
	}

	for _, direction := range directions {
		neighbor := helper.Coordinate{X: position.X + direction.X, Y: position.Y + direction.Y}
		neighborCell, err := grid.ReadCell(neighbor)
		if err == nil && neighborCell == cell+1 {
			walkTrails(grid, neighbor, uniqueEndpoints)
		}
	}
}

func main() {
	grid, err := helper.LoadGrid("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed loading grid:", err)
		os.Exit(1)
	}

	startingPoints := findStartingPoints(grid)
	totalEndpointsReached := 0
	for _, startingPoint := range startingPoints {
		uniqueEndpoints := make(map[helper.Coordinate]struct{})
		walkTrails(grid, startingPoint, uniqueEndpoints)
		totalEndpointsReached += len(uniqueEndpoints)
	}

	fmt.Printf("Part 1 - Total score of all trailheads is: %v\n", totalEndpointsReached)
}
