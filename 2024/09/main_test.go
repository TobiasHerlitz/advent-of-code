package main

import (
	"reflect"
	"testing"
)

var exampleMemory = memory{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9}
var compactedMemory = memory{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

func TestDefragment(t *testing.T) {
	exampleMemory.compact()

	if !reflect.DeepEqual(exampleMemory, compactedMemory) {
		t.Errorf("Incorrectly compacted. got: %v, expect %v", exampleMemory, compactedMemory)
	}
}

func TestGetChecksum(t *testing.T) {
	got := compactedMemory.getChecksum()
	expect := 1928

	if got != expect {
		t.Errorf("Incorrect checksum. got: %v, expect %v", got, expect)
	}
}
