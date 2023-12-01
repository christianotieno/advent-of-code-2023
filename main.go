package main

import (
	"adventOfCode2023/day1"
	"fmt"
)

func main() {
	p1, p2, err := day1.Trebuchet("day1/calibrationDocument.txt")
	if err != nil {
		return
	}
	fmt.Printf("%d %d", p1, p2)
}
