package day2

import "testing"

const testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPartOne(t *testing.T) {
    expected := 2
    result, err := PartOne(testInput)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestPartTwo(t *testing.T) {
    expected := 4
    result, err := PartTwo(testInput)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}