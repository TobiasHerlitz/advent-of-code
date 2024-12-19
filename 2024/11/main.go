package main

import (
	"fmt"
	"os"
	"strconv"
)

type Stones map[int]int

// blink modifies the stones based on the following rules:
// Replaces 0 with 1
// Splits numbers with even number of digits into two stones
// Else multiplies original by 2024
func (s *Stones) blink() error {
	changes := make(Stones)
	for number, occurences := range *s {
		switch {
		case number == 0:
			changes[1] += occurences
			changes[number] -= occurences
		case getNumberOfDigits(number)%2 == 0:
			stoneStr := strconv.Itoa(number)
			leftStone, err := strconv.Atoi(stoneStr[:len(stoneStr)/2])
			if err != nil {
				return err
			}

			rightStone, err := strconv.Atoi(stoneStr[len(stoneStr)/2:])
			if err != nil {
				return err
			}
			changes[number] -= occurences
			changes[leftStone] += occurences
			changes[rightStone] += occurences
		default:
			changes[number] -= occurences
			changes[number*2024] += occurences
		}
	}

	for number, change := range changes {
		(*s)[number] += change
	}
	return nil
}

func (s *Stones) blinkTimes(times int) error {
	for i := 0; i < times; i++ {
		err := s.blink()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Stones) countStones() int {
	count := 0
	for _, occurences := range *s {
		count += occurences
	}

	return count
}

func getNumberOfDigits(num int) int {
	numberOfDigits := 1
	for num > 9 {
		num /= 10
		numberOfDigits++
	}

	return numberOfDigits
}

func main() {
	stones := Stones{
		0:       1,
		6:       1,
		68:      1,
		852:     1,
		3914:    1,
		50706:   1,
		645371:  1,
		5601550: 1,
	}

	err := stones.blinkTimes(25)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when blinking. Original:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - Total number of stones after blinking 25 times is: %v\n", stones.countStones())

	err = stones.blinkTimes(50)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when blinking. Original:", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Total number of stones after blinking 75 times is: %v\n", stones.countStones())
}
