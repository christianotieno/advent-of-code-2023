package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunBinarySearch(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		_ = fmt.Errorf("error opening file: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			return
		}
	}()

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
		ans *= binarySearch(times[i], dist[i])
	}

	return ans
}

func binarySearch(t, d int) int {
	g := func(x int) int {
		return x * (t - x)
	}

	lo := 0
	hi := t / 2

	if hi*(t-hi) < d {
		return 0
	}

	for lo+1 < hi {
		m := (lo + hi) / 2
		if g(m) >= d {
			hi = m
		} else {
			lo = m
		}
	}

	first := hi

	last := int((float64(t) / 2) + (float64(t)/2 - float64(first)))

	return last - first + 1
}
