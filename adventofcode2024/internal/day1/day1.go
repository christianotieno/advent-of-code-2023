package day1

// day1.go: part one https://adventofcode.com/2024/day/1
// day1.go: part two https://adventofcode.com/2024/day/1#part2

import (
	"sort"
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
    lines := strings.Split(strings.TrimSpace(input), "\n")
    
    leftList := make([]int, 0, len(lines))
    rightList := make([]int, 0, len(lines))
    
    for _, line := range lines {
        parts := strings.Fields(line)
        if len(parts) != 2 {
            continue 
        }
        
        left, err := strconv.Atoi(parts[0])
        if err != nil {
            return 0, err
        }
        
        right, err := strconv.Atoi(parts[1])
        if err != nil {
            return 0, err
        }
        
        leftList = append(leftList, left)
        rightList = append(rightList, right)
    }
    
    sort.Ints(leftList)
    sort.Ints(rightList)
    
    totalDistance := 0
    for i := 0; i < len(leftList); i++ {
        distance := abs(leftList[i] - rightList[i])
        totalDistance += distance
    }
    
    return totalDistance, nil
}

func PartTwo(input string) (int, error) {
    lines := strings.Split(strings.TrimSpace(input), "\n")
    
    leftList := make([]int, 0, len(lines))
    rightList := make([]int, 0, len(lines))
    
    for _, line := range lines {
        parts := strings.Fields(line)
        if len(parts) != 2 {
            continue
        }
        
        left, err := strconv.Atoi(parts[0])
        if err != nil {
            return 0, err
        }
        
        right, err := strconv.Atoi(parts[1])
        if err != nil {
            return 0, err
        }
        
        leftList = append(leftList, left)
        rightList = append(rightList, right)
    }
    
    rightCounts := make(map[int]int)
    for _, num := range rightList {
        rightCounts[num]++
    }
    
    similarityScore := 0
    for _, leftNum := range leftList {
        similarityScore += leftNum * rightCounts[leftNum]
    }
    
    return similarityScore, nil
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}