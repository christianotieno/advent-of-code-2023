package day9

import "testing"

const testInput = `2333133121414131402`

func TestPartOne(t *testing.T) {
    expected := 1928
    result, err := PartOne(testInput)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestPartTwo(t *testing.T) {
    expected := 2858
    result, err := PartTwo(testInput)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}