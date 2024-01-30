package day09

import (
	"fmt"
)

type Exercise struct{}

func (e Exercise) Part1(data string) (int, error) {
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

func (e Exercise) Part2(data string) (int, error) {
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
