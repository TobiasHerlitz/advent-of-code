package main

import (
	"reflect"
	"testing"
)

func filterEmpty(stoneMap Stones) Stones {
	filteredStones := make(Stones)
	for number, count := range stoneMap {
		if count == 0 {
			continue
		}

		filteredStones[number] = count
	}

	return filteredStones
}

func TestBlinkOnce(t *testing.T) {
	stones := Stones{125: 1, 17: 1}
	expect := Stones{253000: 1, 1: 1, 7: 1}
	stones.blink()

	if !reflect.DeepEqual(filterEmpty(stones), expect) {
		t.Errorf("Stones are incorrect. got: %v, expect %v - %v", stones, expect, stones.countStones())
	}
}

func TestBlink6Times(t *testing.T) {
	stones := Stones{125: 1, 17: 1}
	expect := Stones{
		0:          2,
		2:          4,
		3:          1,
		4:          1,
		6:          2,
		7:          1,
		8:          1,
		80:         1,
		40:         2,
		48:         2,
		96:         1,
		2024:       1,
		4048:       1,
		14168:      1,
		2097446912: 1,
	}
	stones.blinkTimes(6)

	if !reflect.DeepEqual(filterEmpty(stones), expect) {
		t.Errorf("Stones are incorrect. got: %v, expect %v", stones, expect)
	}
}

func TestBlink25Times(t *testing.T) {
	stones := Stones{17: 1, 125: 1}
	expect := 55312
	stones.blinkTimes(25)

	got := stones.countStones()

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Number of stones are incorrect. got: %v, expect %v", got, expect)
	}
}
