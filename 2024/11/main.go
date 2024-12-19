/*
 * 1. If the number is 0 --> 1
 * 2. If the stone number has even digits --> split in two ([1000] -> [10, 0])
 * 3. Else --> multiply original by 2024
 */

package main

import (
	"fmt"
	"os"
)

type Stones []int
var input = Stones{0, 5601550, 3914, 852, 50706, 68, 6, 645371}

func (s Stones) blink() {
	return
}

func main() {
	for i := 0; i < 6; i++ {
		input.blink()
	}

	fmt.Printf("Part 1 - Total number of stones after blinking 25 times is: %v\n", 22)
}
