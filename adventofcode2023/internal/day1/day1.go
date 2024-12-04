package day1

import (
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0

	for _, line := range lines {
		digits := make([]rune, 0)
		
		for _, c := range line {
			if c >= '0' && c <= '9' {
				digits = append(digits, c)
			}
		}

		if len(digits) > 0 {
			value, _ := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
			sum += value
		}
	}

	return sum, nil
}

func PartTwo(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0

	for _, line := range lines {
		digits := make([]rune, 0)

		for i, c := range line {
			if c >= '0' && c <= '9' {
				digits = append(digits, c)
			}

			wordDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

			
			// _wordDigits := map[string]string {
			// 	"one": "1",
			// 	"two": "2",
			// 	"three": "3",
			// 	"four": "4",
			// 	"five": "5",
			// 	"six": "6",
			// 	"seven": "7",
			// 	"eight": "8",
			// 	"nine": "9",
			// }

			for d, word := range wordDigits {
				if strings.HasPrefix(line[i:], word) {
					digits = append(digits, []rune(strconv.Itoa(d+1))...)
				}
			}
		}

		// runeSlices := []rune(strconv.Itoa(d+1))
		// digits = append(digits, runeSlices[0])
		// digits = append(digits, runeSlices...)

		
		if len(digits) > 0 {
			value, _ := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
			sum += value
		}
	}

	return sum, nil
} 