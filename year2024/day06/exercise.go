package day06

import (
	"fmt"
	"os"

	"github.com/gowikel/adventofcode-golang/year2024/day06/parser"
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

	iterateGrid(g)

	return g.CountVisitedCells(), nil
}

func (e Exercise) Part2(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("part2: %w", err)
	}
	
	data := string(contents)
	g, err := parser.Parse(data)
	if err != nil {
		return 0, fmt.Errorf("part2: %w", err)
	}

	cpy := g.Clone()
	result := 0

	iterateGrid(g)

	for i := 0; i < g.Rows(); i++ {
		for j := 0; j < g.Cols(); j++ {
			if g.IsCellBlocked(i, j) || g.GetVisitsAt(i, j) == 0 {
				continue
			}

			gcpy := cpy.Clone()
			gcpy.SetBlockedCellAt(i, j)
		
			if iterateGridWithCycleDetection(gcpy) {
				result++
			}
		}
	}

	return result, nil
}
