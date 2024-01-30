package day04

import (
	"fmt"
	"strings"
)

type Exercise struct{}

func (e Exercise) Part1(data string) (int, error) {
	var result int

	for _, line := range strings.Split(data, "\n") {
		var lineResult int
		w, p, err := parseLine(line)

		if err != nil {
			return result, fmt.Errorf("Part1: %w", err)
		}

		for _, playedNumber := range p {
			if w.Contains(playedNumber) && lineResult == 0 {
				lineResult = 1
			} else if w.Contains(playedNumber) {
				lineResult *= 2
			}
		}

		result += lineResult
	}

	return result, nil
}

func (e Exercise) Part2(data string) (int, error) {
	var result int
	lines := strings.Split(data, "\n")
	copies := make([]int, len(lines))

	for i := range copies {
		copies[i] = 1
	}

	for idx, line := range strings.Split(data, "\n") {
		w, p, err := parseLine(line)

		if err != nil {
			return result, fmt.Errorf("Part2: %w", err)
		}

		matches, _ := countMatches(w, p)

		for i := idx + 1; i < idx+matches+1; i++ {
			copies[i] += copies[idx]
		}
	}

	for _, cpy := range copies {
		result += cpy
	}

	return result, nil
}
