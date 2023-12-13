package aoc

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed data/2023_04.txt
var DAY4_DATA string
var CARD_REGEX = regexp.MustCompile(`Card \d+: (.*)|(.*)`)

func Day4() {
	fmt.Printf("- Day 04\n")
	fmt.Printf("  Part 1: %d\n", Day4Part1(DAY4_DATA))
	fmt.Printf("  Part 2: %d\n", Day4Part2(DAY4_DATA))
}

func Day4ParseLine(line string) (map[int]struct{}, []int, error) {
	w := make(map[int]struct{})
	n := make([]int, 0)

	cardSeparator := strings.Index(line, ":")
	numbersSeparator := strings.Index(line, "|")

	if cardSeparator == -1 || numbersSeparator == -1 {
		return w, n, fmt.Errorf("invalid line: %q", line)
	}

	winNumbers := line[cardSeparator+1 : numbersSeparator]
	playedNumbers := line[numbersSeparator+1:]

	for _, number := range strings.Fields(winNumbers) {
		parsedInt, err := strconv.Atoi(number)
		if err != nil {
			return w, n, fmt.Errorf("error procesing line: %q\n"+
				"%v is not a number\n"+"%w", line, number, err)
		}

		w[parsedInt] = struct{}{}
	}

	for _, number := range strings.Fields(playedNumbers) {
		parsedInt, err := strconv.Atoi(number)
		if err != nil {
			return w, n, fmt.Errorf("error procesing line: %q\n"+
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
			_, ok := w[playedNumber]
			if ok && lineResult == 0 {
				lineResult = 1
			} else if ok {
				lineResult *= 2
			}
		}

		result += lineResult
	}

	return result
}

func CountMatches(w map[int]struct{}, lst []int) (int, error) {
    if w == nil || lst == nil {
        return 0, fmt.Errorf("map or slice is nil")
    }

    var result int

    for _, n := range lst {
        if _, ok := w[n]; ok {
            result += 1
        }
    }

    return result, nil
}

func Day4Part2(data string) int {
	var result int
	lines := strings.Split(data, "\n")
	copies := make([]int, len(lines))

	for i := range copies {
		copies[i] = 1
	}

	for idx, line := range strings.Split(data, "\n") {
		w, p, err := Day4ParseLine(line)

		if err != nil {
			panic(err)
		}

		matches, _ := CountMatches(w, p)

		for i := idx + 1; i < idx + matches+1; i++ {
			copies[i] += copies[idx]
		}
	}

	for _, cpy := range copies {
		result += cpy
	}

	return result
}
