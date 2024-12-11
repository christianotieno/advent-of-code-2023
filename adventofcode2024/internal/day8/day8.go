package day8

import (
	"strings"

	"adventofcode2024/utils"
)

type Point struct {
    x, y int
}

func parseInput(input string) ([][]rune, map[rune][]Point) {
    lines := strings.Split(strings.TrimSpace(input), "\n")
    grid := make([][]rune, len(lines))
    antennas := make(map[rune][]Point)
    
    for y, line := range lines {
        grid[y] = []rune(line)
        for x, char := range line {
            if char != '.' {
                antennas[char] = append(antennas[char], Point{x, y})
            }
        }
    }
    
    return grid, antennas
}

func isCollinear(p, a1, a2 Point) bool {
    crossProduct := (a2.y - a1.y) * (p.x - a1.x) - (a2.x - a1.x) * (p.y - a1.y)
    return crossProduct == 0
}

func findAntinodes(grid [][]rune, antennas map[rune][]Point, checkDistances bool) int {
    antinodes := make(map[Point]bool)
    height, width := len(grid), len(grid[0])
    
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            p := Point{x, y}
            
            for _, positions := range antennas {
                if len(positions) < 2 {
                    continue
                }
                
                for i := 0; i < len(positions); i++ {
                    for j := i + 1; j < len(positions); j++ {
                        a1, a2 := positions[i], positions[j]
                        
                        if !isCollinear(p, a1, a2) {
                            continue
                        }

                        if !checkDistances {
                            antinodes[p] = true
                            break
                        }

                        dist1 := utils.Abs(p.x-a1.x) + utils.Abs(p.y-a1.y)
                        dist2 := utils.Abs(p.x-a2.x) + utils.Abs(p.y-a2.y)
                        if dist1 == 2*dist2 || dist2 == 2*dist1 {
                            antinodes[p] = true
                        }
                    }
                }
            }
        }
    }
    
    return len(antinodes)
}

func PartOne(input string) (int, error) {
    grid, antennas := parseInput(input)
    return findAntinodes(grid, antennas, true), nil
}

func PartTwo(input string) (int, error) {
    grid, antennas := parseInput(input)
    return findAntinodes(grid, antennas, false), nil
}