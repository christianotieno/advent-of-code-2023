package day10

import "strings"

type Point struct {
	row, col int
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, ch := range line {
			grid[i][j] = int(ch - '0')
		}
	}
	
	return grid
}

func calculateTrailheadScore(grid [][]int, start Point) int {
	rows, cols := len(grid), len(grid[0])
	visited := make(map[Point]bool)
	reachable9s := make(map[Point]bool)

	var dfs func(p Point, height int)
	dfs = func(p Point, height int) {
		if p.row < 0 || p.row >= rows || p.col < 0 || p.col >= cols {
			return
		}
		if visited[p] || grid[p.row][p.col] != height {
			return
		}
		
		visited[p] = true
		if height == 9 {
			reachable9s[p] = true
			return
		}

		// Try all four directions
		dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range dirs {
			next := Point{p.row + dir.row, p.col + dir.col}
			dfs(next, height+1)
		}
	}

	dfs(start, 0)
	return len(reachable9s)
}

func calculateTrailheadRating(grid [][]int, start Point) int {
	rows, cols := len(grid), len(grid[0])
	paths := 0

	var dfs func(p Point, height int)
	dfs = func(p Point, height int) {
		if p.row < 0 || p.row >= rows || p.col < 0 || p.col >= cols {
			return
		}
		if grid[p.row][p.col] != height {
			return
		}
		
		if height == 9 {
			paths++
			return
		}

		// Try all four directions
		dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range dirs {
			next := Point{p.row + dir.row, p.col + dir.col}
			dfs(next, height+1)
		}
	}

	dfs(start, 0)
	return paths
}

func PartOne(input string) (int, error) {
	grid := parseInput(strings.TrimSpace(input))

	totalScore := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 0 {
				score := calculateTrailheadScore(grid, Point{row, col})
				totalScore += score
			}
		}
	}

	return totalScore, nil
}

func PartTwo(input string) (int, error) {
	grid := parseInput(strings.TrimSpace(input))

	totalRating := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 0 {
				rating := calculateTrailheadRating(grid, Point{row, col})
				totalRating += rating
			}
		}
	}

	return totalRating, nil
}

