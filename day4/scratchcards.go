package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScratchCards(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		_ = fmt.Errorf("error opening file: %v+", err)
		return 0
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_ = fmt.Errorf("error closing file %v+", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	totalPoints := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		leftPart := strings.Split(parts[0], ":")

		leftNumbers := extractNumbers(leftPart[1])
		rightNumbers := extractNumbers(parts[1])

		points := calculatePoints(leftNumbers, rightNumbers)
		totalPoints += points
	}

	fmt.Printf("Total Points: %d\n", totalPoints)
	return totalPoints
}

func extractNumbers(s string) []int {
	numStr := strings.Fields(strings.TrimSpace(s))
	numbers := make([]int, len(numStr))

	for i, str := range numStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting %s to integer: %v\n", str, err)
			continue
		}
		numbers[i] = num
	}

	return numbers
}

func calculatePoints(leftNumbers, rightNumbers []int) int {
	points := 0
	matched := make(map[int]bool)

	for _, num := range rightNumbers {
		if contains(leftNumbers, num) && !matched[num] {
			points++
			matched[num] = true
		}
	}

	if points > 1 {
		points = 2 << (points - 2)
	}

	return points
}

func contains(numbers []int, target int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}
