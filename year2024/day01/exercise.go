package day01

import (
	"fmt"
	"math"
	"sort"
)

type Exercise struct{}

func (e Exercise) Part1(data string) (int, error) {
	l1, l2, err := Parse(data)
	if err != nil {
		return 0, fmt.Errorf("part1: %w", err)
	}

	if len(l1) != len(l2) {
		return 0, fmt.Errorf("part1: len(l1) != len(l2)")
	}

	sort.Slice(l1, func(i, j int) bool {
		return l1[i] < l1[j]
	})

	sort.Slice(l2, func(i, j int) bool {
		return l2[i] < l2[j]
	})

	result := 0
	for i := 0; i < len(l1); i++ {
		a := l1[i]
		b := l2[i]

		diff := int(math.Abs(float64(a - b)))
		result += diff
	}

	return result, nil
}

func (e Exercise) Part2(data string) (int, error) {
	l1, l2, err := Parse(data)
	if err != nil {
		return 0, fmt.Errorf("part2: %w", err)
	}

	counter := map[int]int{}
	result := 0

	for _, n := range l2 {
		counter[n]++
	}

	for _, n := range l1 {
		times := counter[n]
		score := n * times

		result += score
	}

	return result, nil
}
