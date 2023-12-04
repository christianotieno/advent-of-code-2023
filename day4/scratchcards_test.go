package day4

import (
	"testing"
)

func TestScratchCards(t *testing.T) {
	got := ScratchCards("scratchcards.txt")
	expect := 13

	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}
