package parser

import (
	"fmt"
	"strings"
)

func Parse(data string) (*Grid, error) {
	data = strings.TrimSpace(data)

	if len(data) == 0 {
		return newGrid(0, 0), nil
	}

	rows := strings.Count(data, "\n") + 1
	cols := strings.Index(data, "\n")

	if cols == -1 {
		cols = len(data)
	}

	g := newGrid(rows, cols)
	lines := strings.Split(data, "\n")

	for i, line := range lines {
		for j, c := range line {
			c := Cell(c)
			if !IsValidCell(c) {
				return g, fmt.Errorf("invalid cell: %c", c)
			}

			g.data[i][j] = c

			if c == GuardTop || c == GuardRight || c == GuardBottom || c == GuardLeft {
				g.isGuardActive = true
				g.guardPosition = [2]int{i, j}
			}
		}
	}

	return g, nil
}
