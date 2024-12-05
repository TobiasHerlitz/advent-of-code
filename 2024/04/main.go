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

func readCharacter(characterGrid [][]byte, x int, y int) (byte, error) {
	if x < 0 || x > len(characterGrid)-1 || y < 0 || y > len(characterGrid[x])-1 {
		return 0, fmt.Errorf("coordinate [%d:%d] is out of bounds", x, y)
	}

	return characterGrid[x][y], nil
}

func searchDirection(characterGrid [][]byte, term []byte, x int, y int, coordinateShift coordinateShift) bool {
	if len(term) == 0 {
		return true
	}

	gridChar, err := readCharacter(characterGrid, x, y)
	if err != nil {
		return false
	}

	if term[0] == gridChar {
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

func isValidPair(charOne byte, charTwo byte) bool {
	if charOne != 'M' && charOne != 'S' {
		return false
	}

	if charTwo != 'M' && charTwo != 'S' {
		return false
	}

	if charOne == charTwo {
		return false
	}

	return true
}

func findXShapedMas(characterGrid [][]byte, x int, y int) bool {
	if characterGrid[x][y] != 'A' {
		return false
	}

	upRightChar, err := readCharacter(characterGrid, x+1, y+1)
	if err != nil {
		return false
	}

	downRightChar, err := readCharacter(characterGrid, x+1, y-1)
	if err != nil {
		return false
	}

	upLeftChar, err := readCharacter(characterGrid, x-1, y+1)
	if err != nil {
		return false
	}

	downLeftChar, err := readCharacter(characterGrid, x-1, y-1)
	if err != nil {
		return false
	}

	if !isValidPair(upRightChar, downLeftChar) || !isValidPair(upLeftChar, downRightChar) {
		return false
	}

	return true
}

func countXShapedMasOccurences(characterGrid [][]byte) int {
	total := 0

	for x := 0; x < len(characterGrid); x++ {
		for y := 0; y < len(characterGrid[0]); y++ {
			if findXShapedMas(characterGrid, x, y) {
				total++
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

	fmt.Printf("Part 1 - Times 'XMAS' occurs: %d\n", countOccurrences(characterGrid, term))
	fmt.Printf("Part 2 - Times x shaped 'mas' occurs: %d\n", countXShapedMasOccurences(characterGrid))
}
