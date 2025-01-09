package parser

import (
	"strings"
)

type Direction int

const (
	DirectionRight Direction = iota
	DirectionLeft
	DirectionUp
	DirectionDown
)

type Grid struct {
	data           [][]int
	blockedCells   [][]bool
	guardPosition  [2]int
	guardActive    bool
	direction      Direction
	rows           int
	cols           int
	visitedCells   int
}

func newGrid(rows, cols int) *Grid {
	data := make([][]int, rows)
	blockedCells := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		data[i] = make([]int, cols)
		blockedCells[i] = make([]bool, cols)
	}

	return &Grid{
		data:         data,
		blockedCells: blockedCells,
		rows:         rows,
		cols:         cols,
	}
}

// Checks if the cell that the guard is looking at is still active
// If the guard is not present, it returns true
func (g *Grid) IsNextCellOutsideGrid() bool {
	if !g.guardActive {
		return true
	}

	x, y := g.guardPosition[0], g.guardPosition[1]

	switch g.direction {
	case DirectionRight:
		return y+1 == g.cols
	case DirectionLeft:
		return y-1 == -1
	case DirectionUp:
		return x-1 == -1
	case DirectionDown:
		return x+1 == g.rows
	}

	panic("Unknown direction. Unreachable position.")
}

// Removes the guard from the grid
func (g *Grid) RemoveGuard() {
	g.guardActive = false
}

// Returns if the guard is present or not
func (g *Grid) IsGuardActive() bool {
	return g.guardActive
}

// Returns true if the next cell is blocked
// If the guard is not present, it returns false
func (g *Grid) IsNextCellBlocked() bool {
	if !g.guardActive || g.IsNextCellOutsideGrid() {
		return false
	}

	x, y := g.guardPosition[0], g.guardPosition[1]

	switch g.direction {
	case DirectionRight:
		return g.blockedCells[x][y+1]
	case DirectionLeft:
		return g.blockedCells[x][y-1]
	case DirectionUp:
		return g.blockedCells[x-1][y]
	case DirectionDown:
		return g.blockedCells[x+1][y]
	}

	panic("Unknown direction. Unreachable position.")
}

// Returns a string representation of the grid
func (g *Grid) String() string {
	var sb strings.Builder

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			if g.guardActive && i == g.guardPosition[0] && j == g.guardPosition[1] {
				switch g.direction {
				case DirectionRight:
					sb.WriteRune('>')
				case DirectionLeft:
					sb.WriteRune('<')
				case DirectionUp:
					sb.WriteRune('^')
				case DirectionDown:
					sb.WriteRune('v')
				}
			} else if g.blockedCells[i][j] {
				sb.WriteRune('#')
			} else if g.data[i][j] > 0 {
				sb.WriteRune('X')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

// Moves the guard to the next cell, if possible
func (g *Grid) Move() {
	if g.IsNextCellOutsideGrid() {
		g.RemoveGuard()
		return
	}

	if g.IsNextCellBlocked() {
		return
	}

	x, y := g.guardPosition[0], g.guardPosition[1]
	fx, fy := x, y

	switch g.direction {
	case DirectionRight:
		fy++
		g.guardPosition[1]++
	case DirectionLeft:
		fy--
		g.guardPosition[1]--
	case DirectionUp:
		fx--
		g.guardPosition[0]--
	case DirectionDown:
		fx++
		g.guardPosition[0]++
	}

	g.data[fx][fy]++

	if g.data[fx][fy] == 1 {
		g.visitedCells++
	}
}

// RotateRight rotates the guard to the right
func (g *Grid) RotateRight() {
	switch g.direction {
	case DirectionRight:
		g.direction = DirectionDown
	case DirectionDown:
		g.direction = DirectionLeft
	case DirectionLeft:
		g.direction = DirectionUp
	case DirectionUp:
		g.direction = DirectionRight
	}
}

// Creates a clone of the grid
func (g *Grid) Clone() *Grid {
	data := make([][]int, g.rows)
	blockedCells := make([][]bool, g.rows)

	for i := 0; i < g.rows; i++ {
		data[i] = make([]int, g.cols)
		blockedCells[i] = make([]bool, g.cols)
		for j := 0; j < g.cols; j++ {
			data[i][j] = g.data[i][j]
			blockedCells[i][j] = g.blockedCells[i][j]
		}
	}

	return &Grid{
		data:          data,
		blockedCells:  blockedCells,
		guardPosition: g.guardPosition,
		guardActive:   g.guardActive,
		direction:     g.direction,
		rows:          g.rows,
		cols:          g.cols,
	}
}

// SetBlockedCellAt: Blocks the given coordinates
func (g *Grid) SetBlockedCellAt(x, y int) {
	g.blockedCells[x][y] = true
}

// Rows: Returns the number of rows of the grid
func (g *Grid) Rows() int {
	return g.rows
}

// Cols: Returns the number of columns of the grid
func (g *Grid) Cols() int {
	return g.cols
}

// IsCellBlocked: Returns true if the given coordinates are blocked
func (g *Grid) IsCellBlocked(x, y int) bool {
	return g.blockedCells[x][y]
}

// GetCurrentGuardPosition: Returns the current position of the guard and if it's active
func (g *Grid) GetCurrentGuardPosition() ([2]int, bool) {
	return g.guardPosition, g.guardActive
}

// GetGuardDirection: Returns the direction of the guard
func (g *Grid) GetGuardDirection() Direction {
	return g.direction
}

// CountVisitedCells: Returns the number of visited cells
func (g *Grid) CountVisitedCells() int {
	return g.visitedCells
}

func (g *Grid) GetVisitsAt(x, y int) int {
	return g.data[x][y]
}
