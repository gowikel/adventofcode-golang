package parser

import (
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
			if c == '#' {
				g.blockedCells[i][j] = true
			} else if c == '^' {
				g.guardPosition = [2]int{i, j}
				g.direction = DirectionUp
				g.guardActive = true
			} else if c == 'v' {
				g.guardPosition = [2]int{i, j}
				g.direction = DirectionDown
				g.guardActive = true
			} else if c == '>' {
				g.guardPosition = [2]int{i, j}
				g.direction = DirectionRight
				g.guardActive = true
			} else if c == '<' {
				g.guardPosition = [2]int{i, j}
				g.direction = DirectionLeft
				g.guardActive = true
			}
		}
	}

	return g, nil
}
