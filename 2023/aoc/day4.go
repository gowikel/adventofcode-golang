package aoc

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed data/2023_04.txt
var DAY4_DATA string
var CARD_REGEX = regexp.MustCompile(`Card \d+: (.*)|(.*)`)

func Day4() {
	fmt.Printf("- Day 04\n")
	fmt.Printf("  Part 1: %d\n", Day4Part1(DAY4_DATA))
	fmt.Printf("  Part 2: %d\n", Day1Part2(DAY4_DATA))
}

func Day4ParseLine(line string) (map[int]bool, []int, error) {
	w := make(map[int]bool)
	n := make([]int, 0)

	// Example line:
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

	cardSeparator := strings.Index(line, ":")
	numbersSeparator := strings.Index(line, "|")

	if cardSeparator == -1 || numbersSeparator == -1 {
		return nil, nil, fmt.Errorf("invalid line: %q", line)
	}

	winNumbers := line[cardSeparator+1 : numbersSeparator]
	playedNumbers := line[numbersSeparator+1:]

	for _, number := range strings.Fields(winNumbers) {
		var parsedInt int
		_, err := fmt.Sscanf(number, "%d", &parsedInt)

		if err != nil {
			return nil, nil, fmt.Errorf("error procesing line: %q\n"+
				"%v is not a number\n"+"%w", line, number, err)
		}

		w[parsedInt] = true
	}

	for _, number := range strings.Fields(playedNumbers) {
		var parsedInt int
		_, err := fmt.Sscanf(number, "%d", &parsedInt)

		if err != nil {
			return nil, nil, fmt.Errorf("error procesing line: %q\n"+
				"%v is not a number\n"+"%w", line, number, err)
		}

		n = append(n, parsedInt)
	}

	return w, n, nil
}

func Day4Part1(data string) int {
	var result int

	for _, line := range strings.Split(data, "\n") {
		var lineResult int
		w, p, err := Day4ParseLine(line)

		if err != nil {
			panic(err)
		}

		for _, playedNumber := range p {
			if w[playedNumber] && lineResult == 0 {
				lineResult = 1
			} else if w[playedNumber] {
				lineResult *= 2
			}
		}

		result += lineResult
	}

	return result
}

func Day4Part2(data string) int {
	var result int
	// lines := strings.Split(data, "\n")
	// copies := make([]int, len(lines))

	// for idx, line := range strings.Split(data, "\n") {
	// 	var lineResult int
	// 	w, p, err := Day4ParseLine(line)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	for _, playedNumber := range p {
	// 		if w[playedNumber] && lineResult == 0 {
	// 			lineResult = 1
	// 		} else if w[playedNumber] {
	// 			lineResult *= 2
	// 		}
	// 	}

	// 	result += lineResult
	// }

	return result
}
