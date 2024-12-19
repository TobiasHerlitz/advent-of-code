package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Stones []int

func hasEvenDigits(num int) bool {
	if num == 0 {
		return true
	}

	return len(strconv.Itoa(num))%2 == 0
}

// blink modifies the stones based on the following rules
// Replaces 0 with 1
// Splits numbers with even number of digits into two stones
// Else multiplies original by 2024
func (s *Stones) blink() error {
	for i := 0; i < len(*s); i++ {
		switch {
		case (*s)[i] == 0:
			(*s)[i] = 1
		case hasEvenDigits((*s)[i]):
			stoneStr := strconv.Itoa((*s)[i])
			leftStone, err := strconv.Atoi(stoneStr[:len(stoneStr)/2])
			if err != nil {
				return err
			}

			rightStone, err := strconv.Atoi(stoneStr[len(stoneStr)/2:])
			if err != nil {
				return err
			}

			(*s)[i] = leftStone
			*s = slices.Insert(*s, i+1, rightStone)
			i++ // Skip newly inserted stone
		default:
			(*s)[i] *= 2024
		}
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

func main() {
	stones := Stones{0, 5601550, 3914, 852, 50706, 68, 6, 645371}
	err := stones.blinkTimes(25)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when blinking. Original:", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 - Total number of stones after blinking 25 times is: %v\n", len(stones))
}
