package main

import (
	"os"
	"slices"
	"strconv"
)

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
		if index%2 == 0 {
			memoryRepresentation = index / 2
		}

		memory = append(memory, slices.Repeat([]int{memoryRepresentation}, intValue)...)
	}

	return memory, nil
}
