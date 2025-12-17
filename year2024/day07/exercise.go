package day07

import (
	"fmt"
	"os"

	"github.com/gowikel/adventofcode-golang/year2024/day07/parser"
	"github.com/gowikel/adventofcode-golang/year2024/day07/search"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (result int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("unable to open file: %w", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			err = fmt.Errorf("error while closing the file: %w", err)
		}
	}()

	equations, err := parser.Parse(file)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}

	for _, eq := range equations {
		validate, err := search.Part1(eq)
		if err != nil {
			return 0, fmt.Errorf("search failed: %w", err)
		}
		if validate {
			result += eq.Result
		}
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("unable to open file: %w", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			err = fmt.Errorf("error while closing the file: %w", err)
		}
	}()

	equations, err := parser.Parse(file)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}

	var result int

	for _, eq := range equations {
		validate, err := search.Part2(eq)
		if err != nil {
			return 0, fmt.Errorf("search failed: %w", err)
		}
		if validate {
			result += eq.Result
		}
	}
	return result, nil
}
