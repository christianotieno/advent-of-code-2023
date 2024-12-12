package day11

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type CacheKey struct {
	stone      string
	blinksLeft int
}

func processStone(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}
	if len(stone)%2 == 0 {
		mid := len(stone) / 2
		left := stone[:mid]
		right := stone[mid:]
		leftNum, _ := strconv.Atoi(left)
		rightNum, _ := strconv.Atoi(right)
		return []string{strconv.Itoa(leftNum), strconv.Itoa(rightNum)}
	}
	num := new(big.Int)
	num.SetString(stone, 10)
	multiplier := big.NewInt(2024)
	num.Mul(num, multiplier)
	return []string{num.String()}
}

func simulateNBlinks(stones []string, n int) int {
	cache := make(map[CacheKey]int) 

	var countStones func(stone string, blinksLeft int) int
	countStones = func(stone string, blinksLeft int) int {
		key := CacheKey{stone, blinksLeft}
		if result, exists := cache[key]; exists {
			return result
		}

		if blinksLeft == 0 {
			cache[key] = 1 
			return 1
		}

		nextStones := processStone(stone)
		total := 0
		for _, nextStone := range nextStones {
			total += countStones(nextStone, blinksLeft-1)
		}

		cache[key] = total
		return total
	}

	totalStones := 0
	for _, stone := range stones {
		totalStones += countStones(stone, n)
	}
	return totalStones
}

func parseInput(input string) ([]string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("invalid input format")
	}

	if len(lines) == 1 {
		return strings.Fields(lines[0]), nil
	}
	return strings.Fields(lines[1]), nil
}

func PartOne(input string) (int, error) {
	stones, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	return simulateNBlinks(stones, 25), nil
}

func PartTwo(input string) (int, error) {
	stones, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	return simulateNBlinks(stones, 75), nil
}