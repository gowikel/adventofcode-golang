package day06

import (
	"fmt"
	"os"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}
	data := string(contents)

	result := 1
	races, err := Parse(data)

	if err != nil {
		return result, fmt.Errorf("Part1: %w", err)
	}

	for _, race := range races {
		waysToWin := CountWaysToWin(race)
		result *= waysToWin
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}
	data := string(contents)

	result := 1
	race, err := ParsePart2(data)

	if err != nil {
		return result, fmt.Errorf("Part2: %w", err)
	}

	// Surprised it was so easy, compared
	// to the previous day
	waysToWin := CountWaysToWin(race)
	result *= waysToWin

	return result, nil
}
