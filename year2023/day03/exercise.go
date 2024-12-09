package day03

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

	points := GetPoints(data_str)
	ranges := GetRanges(data_str, points)
	numberRanges := LocateNumbers(data_str, ranges)
	lines := strings.Split(data_str, "\n")

	var result int

	for _, numberRange := range numberRanges {
		lineIdx := numberRange[0]
		start := numberRange[1]
		end := numberRange[2]

		number := lines[lineIdx][start:end]
		var parsedInt int

		_, err := fmt.Sscanf(number, "%d", &parsedInt)

		if err != nil {
			return result, fmt.Errorf("Part1: %w", err)
		}

		result += parsedInt
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}

	data_str := string(data)

	lines := strings.Split(data_str, "\n")
	potentialGears := GetGears(data_str)

	var result int

	for _, potentialGear := range potentialGears {
		ranges := GetRanges(data_str, [][2]int{potentialGear})
		numberRanges := LocateNumbers(data_str, ranges)

		// To be a gear, you must have two different numbers around
		if len(numberRanges) == 2 {
			var sb strings.Builder
			var g1 int
			var g2 int

			for _, numberRange := range numberRanges {
				lineIdx := numberRange[0]
				start := numberRange[1]
				end := numberRange[2]

				n := lines[lineIdx][start:end]

				sb.WriteString(n)
				sb.WriteRune(' ')
			}

			numbers := sb.String()
			_, err := fmt.Sscanf(numbers, "%d %d ", &g1, &g2)

			if err != nil {
				return result, fmt.Errorf("Part2: %w", err)
			}

			result += g1 * g2
		}
	}

	return result, nil
}
