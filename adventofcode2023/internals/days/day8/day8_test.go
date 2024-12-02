package day8

import "testing"

func TestDay8(t *testing.T) {
	day8 := Day8{}
	got, err := day8.PartOne()
	if err != nil {
		t.Fatalf("Failed to run PartOne: %v", err)
	}
	expected := 0

	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
