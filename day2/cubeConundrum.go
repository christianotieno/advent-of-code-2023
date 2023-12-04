package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CubeConundrum(filepath string) (int, int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, 0, nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	scanner := bufio.NewScanner(file)
	p1 := 0
	p2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		ok := true

		splitLine := strings.Split(line, ":")
		id := splitLine[0]
		line = splitLine[1]

		V := make(map[string]int)

		for _, event := range strings.Split(line, ";") {
			for _, balls := range strings.Split(event, ",") { // [[3 blue]....]
				fields := strings.Fields(balls) // [[3] [blue].....]
				n, color := fields[0], fields[1]
				nInt, _ := strconv.Atoi(n)

				V[color] = maxNumber(V[color], nInt)

				if nInt > map[string]int{"red": 12, "green": 13, "blue": 14}[color] {
					ok = false
				}
			}
		}

		score := 1
		for _, v := range V {
			score *= v
		}
		p2 += score

		if ok {
			idSplit := strings.Fields(id)
			p1 += parseInt(idSplit[len(idSplit)-1])
		}
	}

	return p1, p2, nil
}

func maxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
