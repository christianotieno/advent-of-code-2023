package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GearRatios(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, 0
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}

	R := len(lines)
	C := len(lines[0])

	p1 := 0
	nums := make(map[[2]int][]int)

	for r := 0; r < R; r++ {
		gears := make(map[[2]int]struct{})
		n := 0
		hasPart := false

		for c := 0; c < C; c++ {
			if isDigit(lines[r][c]) {
				num, _ := strconv.Atoi(string(lines[r][c]))
				n = n*10 + num

				for rr := -1; rr <= 1; rr++ {
					for cc := -1; cc <= 1; cc++ {
						nr, nc := r+rr, c+cc
						if nr >= 0 && nr < R && nc >= 0 && nc < C {
							ch := lines[nr][nc]
							if !isDigit(ch) && ch != '.' {
								hasPart = true
							}
							if ch == '*' {
								gears[[2]int{nr, nc}] = struct{}{}
							}
						}
					}
				}
			} else if n > 0 {
				for gear := range gears {
					nums[gear] = append(nums[gear], n)
				}
				if hasPart {
					p1 += n
				}
				n = 0
				hasPart = false
				gears = make(map[[2]int]struct{})
			}
		}

		// Add the gear ratios for the last part number if needed
		if n > 0 {
			for gear := range gears {
				nums[gear] = append(nums[gear], n)
			}
			if hasPart {
				p1 += n
			}
		}
	}

	p2 := 0
	for _, v := range nums {
		if len(v) == 2 {
			p2 += v[0] * v[1]
		}
	}

	return p1, p2
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
