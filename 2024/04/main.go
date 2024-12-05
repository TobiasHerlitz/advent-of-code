package main

import (
	"bytes"
	"fmt"
	"os"
)

type coordinateShift struct {
	x int
	y int
}

var directions = map[string]coordinateShift{
	"north":     {x: 0, y: 1},
	"northeast": {x: 1, y: 1},
	"east":      {x: 1, y: 0},
	"southeast": {x: 1, y: -1},
	"south":     {x: 0, y: -1},
	"southwest": {x: -1, y: -1},
	"west":      {x: -1, y: 0},
	"northwest": {x: -1, y: 1},
}

func inBounds(characterGrid [][]byte, x int, y int) bool {
	return x >= 0 && x < len(characterGrid) && y >= 0 && y < len(characterGrid[x])
}

func searchDirection(characterGrid [][]byte, term []byte, x int, y int, coordinateShift coordinateShift) bool {
	if len(term) == 0 {
		return true
	}

	if !inBounds(characterGrid, x, y) {
		return false
	}

	if term[0] == characterGrid[x][y] {
		return searchDirection(characterGrid, term[1:], x+coordinateShift.x, y+coordinateShift.y, coordinateShift)
	}

	return false
}

func searchAllDirections(characterGrid [][]byte, term []byte, x int, y int) int {
	hits := 0
	for _, direction := range directions {
		hit := searchDirection(characterGrid, term, x, y, direction)
		if hit == true {
			hits++
		}
	}

	return hits
}

func countOccurrences(characterGrid [][]byte, term []byte) int {
	total := 0

	for x := 0; x < len(characterGrid); x++ {
		for y := 0; y < len(characterGrid[0]); y++ {
			if characterGrid[x][y] == term[0] {
				total += searchAllDirections(characterGrid, term, x, y)
			}
		}
	}
	return total
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	characterGrid := bytes.Split(input, []byte("\n"))

	term := []byte("XMAS")
	occurences := countOccurrences(characterGrid, term)

	fmt.Printf("Part 1 - Times 'XMAS' occurs: %d\n", occurences)
	// fmt.Printf("Part 2 - Total of all valid and active entries: %d\n", )
}
