package main

import (
	"fmt"
	"log"

	"adventOfCode2023/internal/days"
)

func main() {
    day8 := days.Day8{}
    
    result, err := day8.PartOne()
    if err != nil {
        log.Fatalf("Day 8 Part 1 failed: %v", err)
    }
    
    fmt.Printf("Day 8 Part 1 Result: %d\n", result)
}