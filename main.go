package main

import (
	"adventOfCode2023/day1"
	"adventOfCode2023/day2"
	"adventOfCode2023/day3"
	"adventOfCode2023/day5"
	"fmt"
)

func main() {
	p1, p2, err := day1.Trebuchet("day1/trebuchet.txt")
	if err != nil {
		return
	}
	fmt.Printf("Trebuchet: %d %d \n", p1, p2)

	p1, p2, err = day2.CubeConundrum("day2/cubeConundrum.txt")
	if err != nil {
		return
	}
	fmt.Printf("CubeConundrum: %d %d \n", p1, p2)

	p1, p2 = day3.GearRatios("day3/gearRatios.txt")
	if err != nil {
		return
	}
	fmt.Printf("Gear Ratios: %d %d\n", p1, p2)

	p1, p2 = day5.SeedFertilizer("day5/seedFertilizerTestFile.txt")

	fmt.Printf("SeedFertilizer: %d %d \n", p1, p2)
}
