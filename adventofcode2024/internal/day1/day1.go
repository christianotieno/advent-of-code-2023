package day1

import (
	"sort"
	"strconv"
	"strings"
)

func CalculateDistance(input string) (int, error) {
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

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}