package day3

// day3.go: part one https://adventofcode.com/2024/day/3
// day3.go: part two https://adventofcode.com/2024/day/3#part2
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

func PartTwo(input string) (int, error) {
	mulPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	mulRe := regexp.MustCompile(mulPattern)

	doPattern := `(?:do|undo)\(\)`
	dontPattern := `don't\(\)`
	doRe := regexp.MustCompile(doPattern)
	dontRe := regexp.MustCompile(dontPattern)

	doMatches := doRe.FindAllStringIndex(input, -1)
	dontMatches := dontRe.FindAllStringIndex(input, -1)
	
	mulMatches := mulRe.FindAllStringSubmatchIndex(input, -1)
	
	sum := 0
	
	for _, match := range mulMatches {
		pos := match[0]
		
		enabled := true
		lastControlPos := -1
		
		for _, dontMatch := range dontMatches {
			if dontMatch[0] < pos && dontMatch[0] > lastControlPos {
				enabled = false
				lastControlPos = dontMatch[0]
			}
		}
		
		for _, doMatch := range doMatches {
			if doMatch[0] < pos && doMatch[0] > lastControlPos {
				enabled = true
				lastControlPos = doMatch[0]
			}
		}
		
		if enabled {
			num1, _ := strconv.Atoi(input[match[2]:match[3]])
			num2, _ := strconv.Atoi(input[match[4]:match[5]])
			sum += num1 * num2
		}
	}

	return sum, nil
}