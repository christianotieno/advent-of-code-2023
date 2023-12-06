package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WaitForIt(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		_ = fmt.Errorf("error opening file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var times, dist []int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) > 1 {
			values := convertToIntSlice(fields[1:])
			switch strings.ToLower(fields[0]) {
			case "time:":
				times = values
			case "distance:":
				dist = values
			}
		}
	}

	if err := scanner.Err(); err != nil {
		_ = fmt.Errorf("error reading file: %v", err)
	}

	ans := 1
	for i := 0; i < len(times); i++ {
		ans *= countValidDistances(times[i], dist[i])
	}

	return ans
}

func convertToIntSlice(strSlice []string) []int {
	intSlice := make([]int, len(strSlice))
	for i, v := range strSlice {
		num, err := strconv.Atoi(v)
		if err != nil {
			_ = fmt.Errorf("error converting string to int: %v", err)
		}
		intSlice[i] = num
	}
	return intSlice
}

func convertToInt(values []int) int {
	result, multiplier := 0, 1
	for i := len(values) - 1; i >= 0; i-- {
		result += values[i] * multiplier
		multiplier *= 10
	}
	return result
}

func countValidDistances(time, distance int) int {
	validDistances := 0
	for x := 0; x <= time; x++ {
		dx := x * (time - x)
		if dx >= distance {
			validDistances++
		}
	}
	return validDistances
}
