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
			switch c {
			case '#':
				g.blockedCells[i][j] = true
			case '^', 'v', '>', '<':
				g.guardPosition = [2]int{i, j}
				g.guardActive = true
			}

			switch c {
			case '^':
				g.direction = DirectionUp
			case 'v':
				g.direction = DirectionDown
			case '>':
				g.direction = DirectionRight
			case '<':
				g.direction = DirectionLeft
			}
		}
	}

	return g, nil
}
