package day1

import "testing"

func TestCalculateDistance(t *testing.T) {
    input := `3   4
4   3
2   5
1   3
3   9
3   3`

    expected := 11
    result, err := CalculateDistance(input)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}