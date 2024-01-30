package day01

import "fmt"

type Exercise struct{}

func (e Exercise) Part1(data string) (int, error) {
	numbers, err := ParseInput(data)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum, nil
}

func (e Exercise) Part2(data string) (int, error) {
	numbers, err := EnhancedParseInput(data)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum, nil
}
