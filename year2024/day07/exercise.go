package day07

import (
	"fmt"
	"os"

	"go.eryndalor.dev/adventofcode-golang/year2024/day07/parser"
	"go.eryndalor.dev/adventofcode-golang/year2024/day07/search"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	var result int

	equations, err := parseFile(path)
	if err != nil {
		return 0, fmt.Errorf("part1: %w", err)
	}

	for _, eq := range equations {
		if search.Part1(eq) {
			result += eq.Result
		}
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	var result int

	equations, err := parseFile(path)
	if err != nil {
		return 0, fmt.Errorf("part2: %w", err)
	}

	for _, eq := range equations {
		if search.Part2(eq) {
			result += eq.Result
		}
	}
	return result, nil
}

func parseFile(path string) (result []parser.Equation, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			err = fmt.Errorf("error while closing the file: %w", err)
		}
	}()

	equations, err := parser.Parse(file)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}

	return equations, nil
}
