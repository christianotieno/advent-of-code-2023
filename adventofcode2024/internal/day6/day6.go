package day6

import (
	"strings"
)

type Position struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

func PartOne(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	
	var start Position
	directions := map[string]Direction{
		"^": {0, -1},
		">": {1, 0}, 
		"v": {0, 1},
		"<": {-1, 0},
	}
	
	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = []rune(line)
		if idx := strings.IndexAny(line, "^>v<"); idx != -1 {
			start = Position{idx, y}
		}
	}
	
	visited := make(map[Position]bool)
	pos := start
	dir := "^"
	visited[pos] = true
	
	for {
		nextPos := Position{
			x: pos.x + directions[dir].dx,
			y: pos.y + directions[dir].dy,
		}
		
		if nextPos.y < 0 || nextPos.y >= len(grid) || 
		   nextPos.x < 0 || nextPos.x >= len(grid[0]) {
			break
		}
		
		if grid[nextPos.y][nextPos.x] == '#' {
			dir = map[string]string{
				"^": ">",
				">": "v", 
				"v": "<",
				"<": "^",
			}[dir]
		} else {
			pos = nextPos
			visited[pos] = true
		}
	}
	
	return len(visited), nil
}

func PartTwo(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	
	var start Position
	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = []rune(line)
		if idx := strings.IndexAny(line, "^>v<"); idx != -1 {
			start = Position{idx, y}
		}
	}
	
	loopCount := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != '.' || (x == start.x && y == start.y) {
				continue
			}
			if createsLoop(grid, start, Position{x, y}) {
				loopCount++
			}
		}
	}
	
	return loopCount, nil
}

func createsLoop(originalGrid [][]rune, start, obstacle Position) bool {
	grid := make([][]rune, len(originalGrid))
	for i := range originalGrid {
		grid[i] = make([]rune, len(originalGrid[i]))
		copy(grid[i], originalGrid[i])
	}
	grid[obstacle.y][obstacle.x] = '#'
	
	directions := map[string]Direction{
		"^": {0, -1},
		">": {1, 0},
		"v": {0, 1},
		"<": {-1, 0},
	}
	
	type visitState struct {
		pos Position
		dir string
	}
	visited := make(map[visitState]int)
	
	pos := start
	dir := "^"
	steps := 0
	maxSteps := len(grid) * len(grid[0]) * 4
	
	for steps < maxSteps {
		state := visitState{pos, dir}
		if count, exists := visited[state]; exists && steps-count > 1 {
			return true
		}
		visited[state] = steps
		
		nextPos := Position{
			x: pos.x + directions[dir].dx,
			y: pos.y + directions[dir].dy,
		}
		
		if nextPos.y < 0 || nextPos.y >= len(grid) || 
		   nextPos.x < 0 || nextPos.x >= len(grid[0]) {
			return false
		}
		
		if grid[nextPos.y][nextPos.x] == '#' {
			dir = map[string]string{
				"^": ">",
				">": "v",
				"v": "<",
				"<": "^",
			}[dir]
		} else {
			pos = nextPos
		}
		
		steps++
	}
	
	return false
}
