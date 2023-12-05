package day5

import "testing"

func TestSeedFertilizer(t *testing.T) {
	p1, p2 := SeedFertilizer("seedFertilizerTestFile.txt")
	expectedP1 := 35
	expectedP2 := 0

	if expectedP1 != p1 {
		t.Errorf("Expected %d, got %d", expectedP1, p1)
	}

	if expectedP2 != p2 {
		t.Errorf("Expected %d, got %d", expectedP2, p2)
	}
}
