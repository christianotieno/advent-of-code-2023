package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Trebuchet(filePath string) (int, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, fmt.Errorf("error opening file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	p1 := 0
	p2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		p1Digits := make([]rune, 0)
		p2Digits := make([]rune, 0)

		for i, c := range line {
			if c >= '0' && c <= '9' {
				p1Digits = append(p1Digits, c)
				p2Digits = append(p2Digits, c)
			}

			digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
			for d, val := range digits {
				if strings.HasPrefix(line[i:], val) {
					p2Digits = append(p2Digits, []rune(strconv.Itoa(d+1))...)
				}
			}
		}

		if len(p1Digits) > 0 {
			p1Int, _ := strconv.Atoi(string(p1Digits[0]) + string(p1Digits[len(p1Digits)-1]))
			p1 += p1Int
		}

		if len(p2Digits) > 0 {
			p2Int, _ := strconv.Atoi(string(p2Digits[0]) + string(p2Digits[len(p2Digits)-1]))
			p2 += p2Int
		}
	}

	return p1, p2, nil
}
