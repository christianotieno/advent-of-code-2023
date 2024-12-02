package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Almanac represents a structure containing a two-dimensional slice of integers.
type Almanac struct {
	TwoDSlices [][]int
}

// SeedFertilizer reads a file, extracts seed and fertilizer data, and applies transformations.
func SeedFertilizer(filepath string) (int, int) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		_ = fmt.Errorf("error reading file: %v", err)
		return 0, 0
	}

	D := strings.TrimSpace(string(data))

	parts := strings.Fields(D)
	seed, others := parts[1], parts[2:]

	seedValues := parseSeed(seed)
	fmt.Printf("File contents: %v", seedValues)

	Fs := make([]*Almanac, len(others))
	for i, s := range others {
		Fs[i] = newAlmanac(s)
	}

	apply := func(almanac *Almanac, i int) int {
		return ApplyOne(almanac, i)
	}

	P1 := applyTransformations(seedValues, Fs, apply)
	P2 := applyTransformations(seedValues, Fs, func(almanac *Almanac, i int) int {
		R := [][]int{{i, i + 1}}
		return minimumValue(ApplyRange(almanac, R)[0])
	})

	return minimumValue(P1), minimumValue(P2)
}

// ApplyOne applies a transformation to a single value.
func ApplyOne(almanac *Almanac, x int) int {
	for _, tuple := range almanac.TwoDSlices {
		dst, src, sz := tuple[0], tuple[1], tuple[2]
		if src <= x && x < src+sz {
			return x + dst - src
		}
	}
	return x
}

// ApplyRange applies a range of transformations to a slice of ranges.
func ApplyRange(almanac *Almanac, R [][]int) [][]int {
	A := make([][]int, 0)

	for _, twoDSlice := range almanac.TwoDSlices {
		dest, src, sz := twoDSlice[0], twoDSlice[1], twoDSlice[2]
		srcEnd := src + sz

		for len(R) > 0 {
			stEd := R[len(R)-1]
			st, ed := stEd[0], stEd[1]

			before := []int{st, minNumber(ed, src)}
			inter := []int{maxNumber(st, src), minNumber(srcEnd, ed)}
			after := []int{maxNumber(srcEnd, st), ed}

			if before[1] > before[0] {
				A = append(A, before)
			}
			if inter[1] > inter[0] {
				A = append(A, []int{inter[0] - src + dest, inter[1] - src + dest})
			}
			if after[1] > after[0] {
				R = append(R, after)
			}
			R = R[:len(R)-1]
		}
	}

	return A
}

// Min finds the minimum value in a slice of integers.
func minimumValue(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

// newAlmanac creates an Almanac from a string representation.
func newAlmanac(s string) *Almanac {
	lines := strings.Split(s, "\n")[1:]
	tuples := make([][]int, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		tuple := make([]int, len(fields))
		for j, field := range fields {
			val, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			tuple[j] = val
		}
		tuples[i] = tuple
	}
	return &Almanac{TwoDSlices: tuples}
}

// parseSeed extracts seed values from a string.
func parseSeed(seed string) []int {
	seedValues := make([]int, 0)
	fields := strings.Fields(seed)[1:]
	for _, val := range fields {
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("Error converting '%s' to integer: %v\n", val, err)
			return nil
		}
		seedValues = append(seedValues, num)
	}
	return seedValues
}

// applyTransformations applies a transformation function to a set of seed values.
func applyTransformations(seedValues []int, Fs []*Almanac, applyFunc func(almanac *Almanac, i int) int) []int {
	result := make([]int, 0)
	for _, x := range seedValues {
		for _, f := range Fs {
			x = applyFunc(f, x)
		}
		result = append(result, x)
	}
	return result
}

// minNumber returns the minimum of two integers.
func minNumber(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// maxNumber returns the maximum of two integers.
func maxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}
