package day3

import (
	"regexp"
	"strconv"
)

func PartOne(input string) (int, error) {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	
	matches := re.FindAllStringSubmatch(input, -1)
	sum := 0
	
	for _, match := range matches {
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}
		
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			continue
		}
		
		sum += num1 * num2
	}
	
	return sum, nil
}
