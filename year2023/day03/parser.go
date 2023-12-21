package day03

import (
	"regexp"
	"strings"
)

func GetPoints(input string) [][2]int {
	result := make([][2]int, 0)

	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			if (char >= '0' && char <= '9') || char == '.' {
				continue
			}

			point := [2]int{i, j}
			result = append(result, point)
		}
	}
	return result
}

func GetRanges(input string, points [][2]int) [][3]int {
	result := make([][3]int, 0, len(points)*3)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, point := range points {
		x := point[0]
		y := point[1]

		leftTarget := y - 1
		rightTarget := y + 2 // [s:e] will not cover e

		if leftTarget < 0 {
			leftTarget = 0
		}

		if rightTarget >= len(lines[x]) {
			rightTarget = len(lines[x]) - 1
		}

		if x > 0 {
			result = append(
				result,
				[3]int{x - 1, leftTarget, rightTarget},
			)
		}

		result = append(result, [3]int{x, leftTarget, rightTarget})

		if x < len(lines)-1 {
			result = append(
				result,
				[3]int{x + 1, leftTarget, rightTarget},
			)
		}
	}

	return result
}

func LocateNumbers(input string, ranges [][3]int) [][3]int {
	number := regexp.MustCompile(`\d+`)

	set := make(map[[3]int]bool)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Calculate ranges
	for _, r := range ranges {
		lineIndex := r[0]
		line := lines[lineIndex]

		start := r[1]
		end := r[2]

		matches := number.FindAllStringIndex(line[start:end], -1)

		if len(matches) > 0 {
			for _, match := range matches {
				startTarget := match[0]
				endTarget := match[1]

				if start > 0 {
					startTarget += start
					endTarget += start
				}

				// Expand to the left
				for startTarget != 0 {
					nextElement := line[startTarget-1]

					if !(nextElement >= '0' && nextElement <= '9') {
						break
					}

					startTarget -= 1
				}

				// Expand to the right
				for endTarget < len(line) {
					nextElement := line[endTarget]

					if !(nextElement >= '0' && nextElement <= '9') {
						break
					}

					endTarget += 1
				}

				// Store the range in a set to avoid repetitions
				set[[3]int{lineIndex, startTarget, endTarget}] = true
			}
		}
	}

	// Convert the set into a [][3]int
	result := make([][3]int, 0, len(set))
	for r := range set {
		result = append(result, r)
	}

	return result
}

func GetGears(input string) [][2]int {
	lines := strings.Split(input, "\n")
	result := make([][2]int, 0)

	for i, line := range lines {
		for j, symbol := range line {
			if symbol == '*' {
				result = append(result, [2]int{i, j})
			}
		}
	}

	return result
}
