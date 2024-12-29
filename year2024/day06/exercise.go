package day06

import (
	"fmt"
	"os"

	"github.com/gowikel/adventofcode-golang/year2024/day06/parser"

	"github.com/gowikel/adventofcode-golang/internal/runner"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("part1: %w", err)
	}

	data := string(contents)
	g, err := parser.Parse(data)
	if err != nil {
		return 0, fmt.Errorf("part1: %w", err)
	}

	nc := g.NextCell()
	for nc.IsGuardActive {
		if nc.Cell == parser.BlockedCell {
			g.RotateRight()
		}
		g.Move()
		nc = g.NextCell()
	}

	return g.VisitedCells(), nil
}

func (e Exercise) Part2(path string) (int, error) {
	return 0, runner.ErrPartNotImplemented
}
