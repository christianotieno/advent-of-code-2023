package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"adventofcode2024/internal/day1"
	"adventofcode2024/internal/day2"
)

func fetchInput(day int) (string, error) {
	cacheDir := "inputs"
	cacheFile := filepath.Join(cacheDir, fmt.Sprintf("day%d.txt", day))

	if content, err := os.ReadFile(cacheFile); err == nil {
		return string(content), nil
	}

	sessionCookie := os.Getenv("AOC_SESSION")
	if sessionCookie == "" {
		return "", fmt.Errorf("AOC_SESSION environment variable not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create cache directory: %w", err)
	}

	if err := os.WriteFile(cacheFile, body, 0644); err != nil {
		return "", fmt.Errorf("failed to cache input: %w", err)
	}

	return string(body), nil
}

func runDay(day int) error {
	input, err := fetchInput(day)
	if err != nil {
		return fmt.Errorf("failed to fetch input for day %d: %v", day, err)
	}

	switch day {
	case 1:
		// Part One
		result1, err := day1.PartOne(input)
		if err != nil {
			return fmt.Errorf("day 1 part 1 failed: %v", err)
		}
		fmt.Printf("Day 1 Part 1 Result: %d\n", result1)

		// Part Two
		result2, err := day1.PartTwo(input)
		if err != nil {
			return fmt.Errorf("day 1 part 2 failed: %v", err)
		}
		fmt.Printf("Day 1 Part 2 Result: %d\n", result2)
	case 2:
		result1, err := day2.PartOne(input)
		if err != nil {
			return fmt.Errorf("day 2 part 1 failed: %v", err)
		}
		fmt.Printf("Day 2 Part 1 Result: %d\n", result1)
	default:
		return fmt.Errorf("day %d not implemented yet", day)
	}

	return nil
}

func main() {
	day := 1
	if len(os.Args) > 1 {
		var err error
        day, err = strconv.Atoi(os.Args[1])
        if err != nil {
            log.Fatalf("Invalid day specified: %v", err)
		}
	}

	if err := runDay(day); err != nil {
		log.Fatalf("Error running day %d: %v", day, err)
	}
}