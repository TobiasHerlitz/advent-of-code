package main

import (
	"reflect"
	"testing"
)

func TestBlinkOnce(t *testing.T) {
	stones := Stones{125, 17}
	expect := Stones{253000, 1, 7}
	stones.blink()

	if !reflect.DeepEqual(stones, expect) {
		t.Errorf("Stones are incorrect. got: %v, expect %v", stones, expect)
	}
}

func TestBlink6Times(t *testing.T) {
	stones := Stones{125, 17}
	expect := Stones{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}
	stones.blinkTimes(6)

	if !reflect.DeepEqual(stones, expect) {
		t.Errorf("Stones are incorrect. got: %v, expect %v", stones, expect)
	}
}

func TestBlink25Times(t *testing.T) {
	stones := Stones{125, 17}
	expect := 55312
	stones.blinkTimes(25)

	got := len(stones)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Number of stones are incorrect. got: %v, expect %v", got, expect)
	}
}
