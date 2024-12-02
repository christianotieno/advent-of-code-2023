package day7

import "testing"

func TestCamelCards(t *testing.T) {
	got, _ := CamelCards("1", false)
	expected := 0

	if expected != got {
		t.Errorf("Expected %v got %v", expected, got)
	}
}

func TestCamelCardsPart2(t *testing.T) {
	got, _ := CamelCards("1", true)
	expected := 0

	if expected != got {
		t.Errorf("Expected %v got %v", expected, got)
	}
}
