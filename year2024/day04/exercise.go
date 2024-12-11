package day04

import (
	"fmt"
	"os"
	"strings"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}
	data := string(contents)
	lines := strings.Split(data, "\n")
	result := 0

	for i, line := range lines {
		for j, c := range line {
			if c == 'X' {
				// Can we write XMAS horizontally?
				if j < len(line)-3 && line[j+1] == 'M' && line[j+2] == 'A' && line[j+3] == 'S' {
					result++
				}

				// Can we write XMAS backwards?
				if j-3 >= 0 && line[j-1] == 'M' && line[j-2] == 'A' && line[j-3] == 'S' {
					result++
				}

				// Can we write XMAS vertically?
				if i < len(lines)-3 && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' &&
					lines[i+3][j] == 'S' {
					result++
				}

				// Can we write XMAS vertically backwards?
				if i-3 >= 0 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' &&
					lines[i-3][j] == 'S' {
					result++
				}

				// Can we write XMAS upper-right diagonal?
				if i-3 >= 0 && j < len(line)-3 && lines[i-1][j+1] == 'M' &&
					lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
					result++
				}

				// Can we write XMAS down-right diagonal?
				if i < len(lines)-3 && j < len(line)-3 && lines[i+1][j+1] == 'M' &&
					lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
					result++
				}

				// Can we write XMAS down-left diagonal?
				if i < len(lines)-3 && j-3 >= 0 && lines[i+1][j-1] == 'M' &&
					lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
					result++
				}

				// Can we write XMAS upper-left diagonal?
				if i-3 >= 0 && j-3 >= 0 && lines[i-1][j-1] == 'M' &&
					lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
					result++
				}
			}
		}
	}

	return result, nil
}

func hasUpperBottomMSCrossAt(lines []string, i, j int) bool {
	return i-1 >= 0 && i+1 < len(lines) && j-1 >= 0 && j+1 < len(lines[i+1]) &&
		((lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') ||
			(lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M'))
}

func hasBottomUpperMSCrossAt(lines []string, i, j int) bool {
	return i-1 >= 0 && i+1 < len(lines) && j-1 >= 0 && j+1 < len(lines[i+1]) &&
		((lines[i+1][j-1] == 'M' && lines[i-1][j+1] == 'S') ||
			(lines[i+1][j-1] == 'S' && lines[i-1][j+1] == 'M'))
}

func (e Exercise) Part2(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}
	data := string(contents)

	lines := strings.Split(data, "\n")
	result := 0

	for i, line := range lines {
		for j, c := range line {
			if c == 'A' && hasBottomUpperMSCrossAt(lines, i, j) &&
				hasUpperBottomMSCrossAt(lines, i, j) {
				result++
			}
		}
	}

	return result, nil
}
