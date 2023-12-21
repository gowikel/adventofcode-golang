package day04

import (
	"fmt"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func parseLine(line string) (mapset.Set[int], []int, error) {
	w := mapset.NewSet[int]()
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

		w.Add(parsedInt)
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