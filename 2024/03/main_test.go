package main

import "testing"

func TestMultiplyNumbers(t *testing.T) {
	exampleString := []byte("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	expected := 161

	got, err := multiplyNumbers(exampleString)
	if err != nil {
		t.Errorf("Unexpected error in multipluNumbers(). Original: %v", err)
	}

	if got != expected {
		t.Errorf("got: %v, expected %v", got, expected)
	}
}
