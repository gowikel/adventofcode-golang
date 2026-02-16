package day01

import (
	"fmt"
	"os"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	// #nosec G304 -- Path is controlled by the user running the solution, not external input
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}

	numbers, err := ParseInput(string(data))
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum, nil
}

func (e Exercise) Part2(path string) (int, error) {
	// #nosec G304 -- Path is controlled by the user running the solution, not external input
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}

	numbers, err := EnhancedParseInput(string(data))
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum, nil
}
