package day8

import (
	"fmt"
	"os"
	"strings"
)

type Day8 struct{}

func (d Day8) PartOne() (int, error) {
	input, err := os.ReadFile("inputs/day8.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read input: %w", err)
	}
	
	lines := strings.TrimSpace(string(input))
	
	fmt.Printf("Input: %v", lines)
	return 0, nil
}

func (d Day8) PartTwo() (int, error) {
	return 0, nil
}
