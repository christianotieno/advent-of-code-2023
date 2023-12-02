package day2

import (
	"testing"
)

func TestCubeConundrum(t *testing.T) {
	testFilePath := "puzzleInputTestFile.txt"

	p1, p2, err := CubeConundrum(testFilePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if p1 != expectedP1Value {
		t.Errorf("Expected p1 value %d, got %d", expectedP1Value, p1)
	}
	if p2 != expectedP2Value {
		t.Errorf("Expected p2 value %d, got %d", expectedP2Value, p2)
	}

}

const expectedP1Value = 8
const expectedP2Value = 2286
