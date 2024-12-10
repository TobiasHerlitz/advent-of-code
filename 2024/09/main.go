package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

type memory []int

func (m memory) compact() {
	firstEmptyIndex := m.findNextEmptyIndex(0)
	for i := len(m) - 1; i > 0; i-- {
		if m[i] == -1 {
			continue
		}

		if firstEmptyIndex == -1 || firstEmptyIndex > i {
			return
		}

		m[firstEmptyIndex] = m[i]
		m[i] = -1

		firstEmptyIndex = m.findNextEmptyIndex(firstEmptyIndex)
	}
}

func (m memory) findNextEmptyIndex(from int) int {
	for i := from; i < len(m) - 1; i++ {
		if m[i] == -1 {
			return i
		}
	}

	return -1
}

func (m memory) getChecksum() int {
	total := 0
	for index, value := range m {
		if value == -1 {
			break
		}

		total += index * value
	}

	return total
}

func loadMemory() (memory, error) {
	input, err := os.ReadFile("input")
	if err != nil {
		return nil, err
	}

	var memory []int
	for index, value := range input {
		intValue, err := strconv.Atoi(string(value))
		if err != nil {
			return nil, err
		}

		memoryRepresentation := -1
		if index % 2 == 0 {
			memoryRepresentation = index / 2
		}
		
		memory = append(memory, slices.Repeat([]int{memoryRepresentation}, intValue)...)
	}

	return memory, nil
}

func main() {
	memory, err := loadMemory()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed loading memory:", err)
		os.Exit(1)
	}

	memory.compact()

	fmt.Printf("Part 1 - The checksum is: %v\n", memory.getChecksum())
}
