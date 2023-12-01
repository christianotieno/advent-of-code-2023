package day1

import (
	"testing"
)

func TestTrabuchet(t *testing.T) {
	file, _ := readFile("testFile.txt")

	expect := trabuchet(file)

	got := ""

	if expect != got {
		t.Errorf("Expected %s, got %s", expect, got)
	}
}
