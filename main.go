package main

import (
	"adventOfCode2023/day1"
	"adventOfCode2023/day2"
	"fmt"
)

func main() {
	p1, p2, err := day1.Trebuchet("day1/calibrationDocument.txt")
	if err != nil {
		return
	}
	fmt.Printf("Trebuchet: %d %d ", p1, p2)

	p1, p2, err = day2.CubeConundrum("day2/puzzleInput.txt")
	if err != nil {
		return
	}
	fmt.Printf("CubeConundrum: %d %d", p1, p2)
}
