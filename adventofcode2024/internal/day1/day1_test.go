package day1

import "testing"

const testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPartOne(t *testing.T) {
    expected := 11
    result, err := PartOne(testInput)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestPartTwo(t *testing.T) {
    expected := 31
    result, err := PartTwo(testInput)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}
