package day4

import (
	"strings"
)

func PartOne(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	count := 0
	target := "XMAS"
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]

				if row+(dx*3) >= 0 && row+(dx*3) < rows &&
				   col+(dy*3) >= 0 && col+(dy*3) < cols {
					match := true
					for i := 0; i < 4; i++ {
						if grid[row+(dx*i)][col+(dy*i)] != rune(target[i]) {
							match = false
							break
						}
					}
					if match {
						count++
					}
				}
			}
		}
	}

	return count, nil
}

func PartTwo(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if grid[row][col] != 'A' {
				continue
			}

			ul_lr := string([]rune{
				grid[row-1][col-1],
				grid[row][col],
				grid[row+1][col+1],
			})

			ur_ll := string([]rune{
				grid[row-1][col+1],
				grid[row][col],
				grid[row+1][col-1],
			})

			isValidDiagonal := func(s string) bool {
				return s == "MAS" || s == "SAM"
			}

			if isValidDiagonal(ul_lr) && isValidDiagonal(ur_ll) {
				count++
			}
		}
	}

	return count, nil
}
