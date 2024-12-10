package main

import (
	"fmt"
	"os"
)

func withinBounds(grid [][]rune, y int, x int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func antinodesFromAntenna(grid [][]rune, y int, x int) map[string]struct{} {
	antenna := grid[y][x]
	antinodes := make(map[string]struct{})
	for opposingY, column := range grid {
		for opposingX, cell := range column {
			if antenna != cell || (y == opposingY && x == opposingX) {
				continue
			}

			antinodeY := y - (opposingY - y)
			antinodeX := x - (opposingX - x)

			if withinBounds(grid, antinodeY, antinodeX) {
				antinodes[fmt.Sprintf("%v:%v", antinodeX, antinodeY)] = struct{}{}
			}
		}
	}

	return antinodes
}

func countAntinodes(grid [][]rune) int {
	antinodes := make(map[string]struct{})
	for y, column := range grid {
		for x, cell := range column {
			if cell != '.' {
				newAntinodes := antinodesFromAntenna(grid, y, x)
				for key := range newAntinodes {
					antinodes[key] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	grid, err := loadGrid()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed loading equations:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - Number of antinodes: %v\n", countAntinodes(grid))
	// fmt.Printf("Part 2 - : %v\n", )
}
