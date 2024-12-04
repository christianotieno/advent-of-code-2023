package day7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type handBid struct {
	hand string
	bid  int
}

func CamelCards(hand string, part2 bool) (int, string) {
	replacements := map[string]string{
		"T": string('9' + 1),
		"J": string('2' - 1),
		"Q": string('9' + 3),
		"K": string('9' + 4),
		"A": string('9' + 5),
	}

	if part2 {
		replacements["J"] = string('9' + 2)
	}

	for from, to := range replacements {
		hand = strings.ReplaceAll(hand, from, to)
	}

	counter := make(map[rune]int)
	for _, c := range hand {
		counter[c]++
	}

	if part2 {
		var target rune = rune(hand[0])
		for k := range counter {
			if k != '1' && (counter[k] > counter[target] || target == '1') {
				target = k
			}
		}

		if '1' == counter['1'] {
			if target != '1' {
				counter[target] += counter['1']
			}
			delete(counter, '1')
		}
	}

	values := make([]int, 0, len(counter))
	for _, v := range counter {
		values = append(values, v)
	}

	sort.Ints(values)

	switch {
	case values[0] == 5:
		return 10, hand
	case values[0] == 1 && values[3] == 4:
		return 9, hand
	case values[0] == 2 && values[2] == 3:
		return 8, hand
	case values[0] == 1 && values[2] == 3:
		return 7, hand
	case values[0] == 1 && values[1] == 2 && values[2] == 2:
		return 6, hand
	case values[0] == 1 && values[1] == 1 && values[2] == 1 && values[3] == 2:
		return 5, hand
	case values[0] == 1 && values[1] == 1 && values[2] == 1 && values[3] == 1 && values[4] == 1:
		return 4, hand
	default:
		fmt.Printf("%v %s %v\n", counter, hand, values)
		return 0, ""
	}
}

func CalculateTotalBid(lines []string, part2 bool) int {
	handBids := make([]handBid, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		handBids[i] = handBid{hand: parts[0], bid: parseInt(parts[1])}
	}

	sort.Slice(handBids, func(i, j int) bool {
		_, valueI := CamelCards(handBids[i].hand, part2)
		_, valueJ := CamelCards(handBids[j].hand, part2)
		return valueI > valueJ
	})

	totalBid := 0
	for i, hb := range handBids {
		totalBid += (i + 1) * hb.bid
	}

	return totalBid
}

func parseInt(s string) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		return 0
	}
	return result
}

func ReadLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
