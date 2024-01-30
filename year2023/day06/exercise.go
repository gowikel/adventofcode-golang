package day06

import (
	"fmt"
)

type Exercise struct{}

func (e Exercise) Part1(data string) (int, error) {
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

func (e Exercise) Part2(data string) (int, error) {

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
