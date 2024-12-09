package day04

import (
	"fmt"
	"os"
	"strings"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}
	data_str := string(data)

	var result int

	for _, line := range strings.Split(data_str, "\n") {
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

func (e Exercise) Part2(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}
	data := string(contents)

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
