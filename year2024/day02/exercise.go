package day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		levels, err := convertToIntList(fields)
		if err != nil {
			return result, fmt.Errorf("Part1: %w", err)
		}

		if AreLevelsSafe(levels) {
			result++
		}
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := 0

outer:
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		levels, err := convertToIntList(fields)
		if err != nil {
			return result, fmt.Errorf("Part2: %w", err)
		}

		if AreLevelsSafe(levels) ||
			AreLevelsSafe(levels[1:]) ||
			AreLevelsSafe(levels[:len(levels)-1]) {
			result++
			continue
		}

		for i := 1; i < len(levels); i++ {
			t := append([]int(nil), levels[:i]...)
			t = append(t, levels[i+1:]...)

			if AreLevelsSafe(t) {
				result++
				continue outer
			}
		}
	}

	return result, nil
}
