package day8

import "testing"

const testInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestPartOne(t *testing.T) {
    expected := 14
    result, err := PartOne(testInput)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
} 

func TestPartTwo(t *testing.T) {
    expected := 34
    result, err := PartTwo(testInput)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}