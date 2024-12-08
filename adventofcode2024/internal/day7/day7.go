package day7

import (
	"strconv"
	"strings"
)

type equation struct {
	target int
	nums   []int
}

func parseInput(input string) ([]equation, error) {
	var equations []equation
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, ": ")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		var nums []int
		for _, numStr := range strings.Fields(parts[1]) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		equations = append(equations, equation{target: target, nums: nums})
	}
	return equations, nil
}

func evaluatePartOne(nums []int, ops []byte) int {
	result := nums[0]
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case '+':
			result += nums[i+1]
		case '*':
			result *= nums[i+1]
		}
	}
	return result
}

func evaluatePartTwo(nums []int, ops []byte) int {
	result := nums[0]
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case '+':
			result += nums[i+1]
		case '*':
			result *= nums[i+1]
		case '|':
			resultStr := strconv.Itoa(result) + strconv.Itoa(nums[i+1])
			result, _ = strconv.Atoi(resultStr)
		}
	}
	return result
}

func canSolvePartOne(eq equation) bool {
	numOps := len(eq.nums) - 1
	for i := 0; i < (1 << numOps); i++ {
		ops := make([]byte, numOps)
		for j := 0; j < numOps; j++ {
			if (i & (1 << j)) != 0 {
				ops[j] = '+'
			} else {
				ops[j] = '*'
			}
		}
		if evaluatePartOne(eq.nums, ops) == eq.target {
			return true
		}
	}
	return false
}

func canSolvePartTwo(eq equation) bool {
	numOps := len(eq.nums) - 1
	for i := 0; i < (1 << (2 * numOps)); i++ {
		ops := make([]byte, numOps)
		valid := true
		for j := 0; j < numOps; j++ {
			switch (i >> (2 * j)) & 3 {
			case 0:
				ops[j] = '+'
			case 1:
				ops[j] = '*'
			case 2:
				ops[j] = '|'
			default:
				valid = false
			}
		}
		if !valid {
			continue
		}
		if evaluatePartTwo(eq.nums, ops) == eq.target {
			return true
		}
	}
	return false
}

func PartOne(input string) (int, error) {
	equations, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, eq := range equations {
		if canSolvePartOne(eq) {
			sum += eq.target
		}
	}
	return sum, nil
}

func PartTwo(input string) (int, error) {
	equations, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, eq := range equations {
		if canSolvePartTwo(eq) {
			sum += eq.target
		}
	}
	return sum, nil
}

