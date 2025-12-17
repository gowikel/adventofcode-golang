package day07

import (
	"fmt"
	"os"

	"github.com/gowikel/adventofcode-golang/internal/runner"
	"github.com/gowikel/adventofcode-golang/year2024/day07/parser"
	"github.com/gowikel/adventofcode-golang/year2024/day07/search"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (result int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer func() {
		err = file.Close()
	}()

	equations, err := parser.Parse(file)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}

	for _, eq := range equations {
		if search.Part1(eq) {
			result += eq.Result
		}
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	return 0, runner.ErrPartNotImplemented
}
