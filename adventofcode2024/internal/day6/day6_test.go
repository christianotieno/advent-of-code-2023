package day6

import (
	"testing"
)

const testInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPartOne(t *testing.T) {
	expected := 41
	result, err := PartOne(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 6
	result, err := PartTwo(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
