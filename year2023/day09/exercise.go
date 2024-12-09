package day09

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

	lst, err := Parse(data)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}

	var result int
	for _, l := range lst {
		sr := NewSensorRead(l)
		err := sr.Compute()
		if err != nil {
			return 0, fmt.Errorf("Part1: %w", err)
		}
		result += sr.ExtrapolateForward()
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}
	data := string(contents)

	lst, err := Parse(data)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}

	var result int
	for _, l := range lst {
		sr := NewSensorRead(l)
		err := sr.Compute()
		if err != nil {
			return 0, fmt.Errorf("Part2: %w", err)
		}
		result += sr.ExtrapolateBackward()
	}

	return result, nil
}
