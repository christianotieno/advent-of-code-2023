package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScratchCards(filepath string) (int, int) {
	file, err := os.Open(filepath)
	if err != nil {
		_ = fmt.Errorf("error opening file: %v", err)
		return 0, 0
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_ = fmt.Errorf("error closing file %v+", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	totalPoints := 0
	scratchCards := make(map[int]int)

	i := 0
	for scanner.Scan() {
		scratchCards[i]++
		line := scanner.Text()
		parts := strings.Split(line, "|")
		leftPart, rightPart := parts[0], parts[1]
		idCard := strings.Split(leftPart, ":")
		card := idCard[1]
		cardNumbers := parseIntArray(card)
		rightNumbers := parseIntArray(rightPart)
		points := countIntersection(cardNumbers, rightNumbers)

		if points > 0 {
			totalPoints += 1 << (points - 1)
		}
		for j := 0; j < points; j++ {
			scratchCards[i+1+j] += scratchCards[i]
		}
		i++
	}

	return sumMapValues(scratchCards), totalPoints
}

func parseIntArray(s string) []int {
	var result []int
	nums := strings.Fields(s)
	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func countIntersection(a, b []int) int {
	setA := make(map[int]bool)
	for _, x := range a {
		setA[x] = true
	}

	intersectionCount := 0
	for _, y := range b {
		if setA[y] {
			intersectionCount++
		}
	}

	return intersectionCount
}

func sumMapValues(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}
