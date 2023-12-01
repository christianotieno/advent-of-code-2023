package main

import (
	"adventOfCode2023/day1"
	"fmt"
)

func main() {
	sum := day1.Trebuchet("day1/calibrationDocument.txt")
	fmt.Println("Total sum of digit sums:", sum)
}
