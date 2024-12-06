package day5

import "testing"

const testInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestPartOne(t *testing.T) {
	expected := 143
	result, err := PartOne(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 123
	result, err := PartTwo(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

