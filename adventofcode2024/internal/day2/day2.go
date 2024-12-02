package day2

// day2.go: part one https://adventofcode.com/2024/day/2
import (
	"strconv"
	"strings"

	"adventofcode2024/utils"
)

func PartOne(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	safeCount := 0

	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) < 2 {
			continue
		}

		levels := make([]int, len(numbers))
		for i, num := range numbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			levels[i] = n
		}

		if isSequenceSafe(levels) {
			safeCount++
		}
	}
	return safeCount, nil
}

func isSequenceSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	isIncreasing := levels[0] < levels[1]

	for i := 1; i < len(levels); i++ {
		diff := utils.Abs(levels[i] - levels[i-1])

		if diff < 1 || diff > 3 {
			return false
		}

		if isIncreasing && levels[i] <= levels[i-1] {
			return false
		}

		if !isIncreasing && levels[i] >= levels[i-1] {
			return false
		}
	}

	return true
}


func PartTwo(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	safeCount := 0

	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) < 2 {
			continue
		}

		levels := make([]int, len(numbers))
		for i, num := range numbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			levels[i] = n
		}

		if isSequenceSafe(levels) || canBeMadeSafe(levels) {
			safeCount++
		}
	}

	return safeCount, nil
}

func canBeMadeSafe(levels []int) bool {
	if len(levels) < 3 {
		return false
	}

	for i := range levels {
		dampened := make([]int, len(levels)-1)
		copy(dampened, levels[:i])
		copy(dampened[i:], levels[i+1:])

		if isSequenceSafe(dampened) {
			return true
		}
	}

	return false
}
