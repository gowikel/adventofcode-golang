package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}
	data := string(contents)

	matches := regex.FindAllStringSubmatch(data, -1)
	result := 0

	for _, match := range matches {
		// Those two cannot fail, because of the regex
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])

		result += a * b
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}
	data := string(contents)

	matches := regex.FindAllStringSubmatch(data, -1)
	active := true
	result := 0

	for _, match := range matches {
		switch match[0] {
		case "do()":
			active = true
		case "don't()":
			active = false
		}

		if active {
			// Those two cannot fail because of the regex
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			result += a * b
		}
	}
	return result, nil
}
