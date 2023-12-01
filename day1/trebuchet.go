package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Trebuchet(file string) int {
	numbers, _ := extractCombinedFirstAndLastDigitFromFile(file)

	totalSum := sumOfDigits(numbers)

	return totalSum
}

func extractCombinedFirstAndLastDigit(input string) (int, error) {
	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(input, -1)

	if len(matches) == 0 {
		return 0, fmt.Errorf("no digits found in line: %s", input)
	}

	if len(matches) == 1 {
		digit, err := strconv.Atoi(matches[0])
		if err != nil {
			return 0, fmt.Errorf("error converting to integer: %v", err)
		}
		return digit*10 + digit, nil
	}

	firstDigit, err1 := strconv.Atoi(matches[0])
	lastDigit, err2 := strconv.Atoi(matches[len(matches)-1])

	if err1 != nil || err2 != nil {
		return 0, fmt.Errorf("error converting to integer: %v %v", err1, err2)
	}

	combined := firstDigit*10 + lastDigit
	return combined, nil
}

func extractCombinedFirstAndLastDigitFromFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var allPairs []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		combined, err := extractCombinedFirstAndLastDigit(line)

		if err != nil {
			fmt.Println("Error extracting digits:", err)
			continue
		}

		allPairs = append(allPairs, combined)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allPairs, nil
}

func sumOfDigits(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}
