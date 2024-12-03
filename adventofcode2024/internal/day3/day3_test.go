package day3

import (
	"testing"
)

const testInput1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const testInput2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPartOne(t *testing.T) {
	expected := 161
	result, err := PartOne(testInput1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 48
	result, err := PartTwo(testInput2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
