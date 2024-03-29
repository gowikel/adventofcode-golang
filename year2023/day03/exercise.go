package day03

import (
	"fmt"
	"strings"
)

type Exercise struct{}

func (e Exercise) Part1(data string) (int, error) {
	points := GetPoints(data)
	ranges := GetRanges(data, points)
	numberRanges := LocateNumbers(data, ranges)
	lines := strings.Split(data, "\n")

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

func (e Exercise) Part2(data string) (int, error) {
	lines := strings.Split(data, "\n")
	potentialGears := GetGears(data)

	var result int

	for _, potentialGear := range potentialGears {
		ranges := GetRanges(data, [][2]int{potentialGear})
		numberRanges := LocateNumbers(data, ranges)

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
