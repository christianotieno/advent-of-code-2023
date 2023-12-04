package day4

import (
	"testing"
)

func TestScratchCards(t *testing.T) {
	totalCards, totalPoints := ScratchCards("scratchcardsTestFile.txt")
	expectedTotalPoints := 13
	expectedTotalCards := 30

	if totalPoints != expectedTotalPoints {
		t.Errorf("Expected %d, got %d", expectedTotalPoints, totalPoints)
	}

	if totalCards != expectedTotalCards {
		t.Errorf("Expected %d, got %d", expectedTotalCards, totalCards)

	}
}
