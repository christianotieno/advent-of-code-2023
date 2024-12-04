package day1

import "testing"

const testInput1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const testInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestPartOne(t *testing.T) {
    expected := 142
    result, err := PartOne(testInput1)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestPartTwo(t *testing.T) {
    expected := 281
    result, err := PartTwo(testInput2)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
} 