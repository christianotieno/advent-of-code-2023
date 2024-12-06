package day5

import (
	"fmt"
	"strings"
)


func PartOne(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	
	splitIndex := -1
	for i, line := range lines {
		if line == "" {
			splitIndex = i
			break
		}
	}
	
	rules := make(map[int][]int)
	for _, line := range lines[:splitIndex] {
		var before, after int
		if _, err := fmt.Sscanf(line, "%d|%d", &before, &after); err != nil {
			return 0, fmt.Errorf("invalid rule format: %s", line)
		}
		rules[before] = append(rules[before], after)
	}

	sum := 0
	for _, line := range lines[splitIndex:] {
		if line == "" {
			continue
		}
		var nums []int
		for _, n := range strings.Split(line, ",") {
			var num int
			if _, err := fmt.Sscanf(n, "%d", &num); err != nil {
				return 0, fmt.Errorf("invalid number in sequence: %s", n)
			}
			nums = append(nums, num)
		}

		if isValidOrder(nums, rules) {
			middleIndex := len(nums) / 2
			sum += nums[middleIndex]
		}
	}

	return sum, nil
}

func isValidOrder(nums []int, rules map[int][]int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if deps, ok := rules[nums[j]]; ok {
				for _, dep := range deps {
					if dep == nums[i] {
						return false
					}
				}
			}
		}
	}
	return true
}

func PartTwo(input string) (int, error) {
    lines := strings.Split(strings.TrimSpace(input), "\n")
    
    splitIndex := -1
    for i, line := range lines {
        if line == "" {
            splitIndex = i
            break
        }
    }
    
    rules := make(map[int][]int)
    for _, line := range lines[:splitIndex] {
        var before, after int
        if _, err := fmt.Sscanf(line, "%d|%d", &before, &after); err != nil {
            return 0, fmt.Errorf("invalid rule format: %s", line)
        }
        rules[before] = append(rules[before], after)
    }

    sum := 0
    for _, line := range lines[splitIndex:] {
        if line == "" {
            continue
        }
        var nums []int
        for _, n := range strings.Split(line, ",") {
            var num int
            if _, err := fmt.Sscanf(n, "%d", &num); err != nil {
                return 0, fmt.Errorf("invalid number in sequence: %s", n)
            }
            nums = append(nums, num)
        }

        if !isValidOrder(nums, rules) {
            sorted := topologicalSort(nums, rules)
            middleIndex := len(sorted) / 2
            sum += sorted[middleIndex]
        }
    }

    return sum, nil
}

func topologicalSort(nums []int, rules map[int][]int) []int {
    graph := make(map[int][]int)
    inDegree := make(map[int]int)
    
    for _, num := range nums {
        graph[num] = []int{}
        inDegree[num] = 0
    }
    
    for _, num := range nums {
        if deps, ok := rules[num]; ok {
            for _, dep := range deps {
                if _, exists := graph[dep]; exists {
                    graph[num] = append(graph[num], dep)
                    inDegree[dep]++
                }
            }
        }
    }
    
    var result []int
    var queue []int
    
    for num := range graph {
        if inDegree[num] == 0 {
            queue = append(queue, num)
        }
    }
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        result = append(result, current)
        
        for _, neighbor := range graph[current] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }
    
    return result
}
