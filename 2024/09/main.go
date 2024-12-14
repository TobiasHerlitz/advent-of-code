package main

import (
	"fmt"
	"os"
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

func (m memory) findEmptySectionIndex(minSize int) int {
	for i := 0; i < len(m); i++ {
		if m[i] != -1 {
			continue
		}

		emptyCount := 0
		for j := i; j < len(m) && m[j] == -1; j++ {
			emptyCount++
			if emptyCount == minSize {
				return i
			}
		}
	}
	return -1
}

func (m memory) findRightmostFileIndex(from int) (int, int) {
	latestFileId := -1
	length := 0
	for i := from; i >= 0; i-- {
		if i == 0 && latestFileId == -1 {
			return -1, -1
		}

		if i == 0 {
			return i, length + 1
		}

		if latestFileId != m[i] && latestFileId != -1 {
			return i + 1, length
		}

		if m[i] == -1 {
			continue
		}

		length++
		if latestFileId == m[i] || latestFileId == -1 {
			latestFileId = m[i]
			continue
		}
	}

	return -1, -1
}

func (m memory) swapSections(sectionOneIndex int, sectionTwoIndex int, length int) {
	for i := 0; i < length; i++ {
		temp := m[sectionOneIndex+i]
		m[sectionOneIndex+i] = m[sectionTwoIndex+i]
		m[sectionTwoIndex+i] = temp
	}
}

func (m memory) compactWithIntegrity() {
	index := 0
	backwardsIndex := len(m) - 1

	for backwardsIndex > index {
		fileIndex, fileLength := m.findRightmostFileIndex(backwardsIndex)
		if fileIndex == -1 {
			break
		}
		emptyIndex := m.findEmptySectionIndex(fileLength)
		if emptyIndex == -1 {
			backwardsIndex = fileIndex - 1
			continue
		}

		if emptyIndex < fileIndex {
			m.swapSections(emptyIndex, fileIndex, fileLength)
		}

		backwardsIndex = fileIndex - 1
	}
}

func (m memory) findNextEmptyIndex(from int) int {
	for i := from; i < len(m)-1; i++ {
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
			continue
		}

		total += index * value
	}

	return total
}

func main() {
	memoryOne, err := loadMemory()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed loading memory:", err)
		os.Exit(1)
	}

	memoryOne.compact()

	memoryTwo, err := loadMemory()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed loading memory:", err)
		os.Exit(1)
	}

	memoryTwo.compactWithIntegrity()

	fmt.Printf("Part 1 - The checksum is: %v\n", memoryOne.getChecksum())
	fmt.Printf("Part 2 - The checksum is: %v\n", memoryTwo.getChecksum())
}
