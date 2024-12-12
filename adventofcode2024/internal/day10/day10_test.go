package day10

import (
	"testing"
)

const testInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPartOne(t *testing.T) {
	expected := 36
	result, err := PartOne(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}	

func TestPartTwo(t *testing.T) {
	expected := 81
	result, err := PartTwo(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
