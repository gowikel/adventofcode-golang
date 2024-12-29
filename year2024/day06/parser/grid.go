package parser

import (
	"strings"
)

type GuardPosition int

const (
	GuardPositionOutsideGrid GuardPosition = iota
	GuardPositionInsideGrid
)

type Cell rune

const (
	EmptyCell   Cell = '.'
	BlockedCell Cell = '#'
	VisitedCell Cell = 'x'
	GuardTop    Cell = '^'
	GuardRight  Cell = '>'
	GuardBottom Cell = 'v'
	GuardLeft   Cell = '<'
)

var registeredCells = map[Cell]struct{}{
	EmptyCell:   {},
	BlockedCell: {},
	VisitedCell: {},
	GuardTop:    {},
	GuardRight:  {},
	GuardBottom: {},
	GuardLeft:   {},
}

func (c Cell) String() string {
	return string(c)
}

func IsValidCell(c Cell) bool {
	_, ok := registeredCells[c]
	return ok
}

type FutureCell int

const (
	FutureCellOutsideGrid FutureCell = iota
	FutureCellEmpty
	FutureCellBlocked
)

// Grid Represents a 2D grid, where are the action happens
type Grid struct {
	rows          int
	cols          int
	data          [][]Cell
	isGuardActive bool
	guardPosition [2]int
}

func newGrid(rows, cols int) *Grid {
	data := make([][]Cell, rows)

	for i := range data {
		data[i] = make([]Cell, cols)

		for j := range data[i] {
			data[i][j] = EmptyCell
		}
	}

	return &Grid{
		rows:          rows,
		cols:          cols,
		data:          data,
		guardPosition: [2]int{-1, -1},
	}
}

// String returns a string representation of the grid
func (g *Grid) String() string {
	var sb strings.Builder

	for i, row := range g.data {
		for _, c := range row {
			sb.WriteRune(rune(c))
		}

		if i < len(g.data)-1 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

func (g *Grid) Rows() int {
	return g.rows
}

func (g *Grid) Columns() int {
	return g.cols
}

// GuardPosition returns if the guard is inside or outside the grid
// along with the last known position of the guard.
// If there is no known guard position, [2]int{-1, -1} should be
// returned
func (g *Grid) GuardPosition() (GuardPosition, [2]int) {
	if g.isGuardActive {
		return GuardPositionInsideGrid, g.guardPosition
	}

	return GuardPositionOutsideGrid, g.guardPosition
}

// nextCellPosition calculates where the next cell will be
func (g *Grid) nextCellPosition() (int, int) {
	x, y := g.guardPosition[0], g.guardPosition[1]
	guard := g.data[x][y]

	var nextCellX, nextCellY int

	switch guard {
	case GuardTop:
		nextCellX, nextCellY = x-1, y
	case GuardRight:
		nextCellX, nextCellY = x, y+1
	case GuardBottom:
		nextCellX, nextCellY = x+1, y
	case GuardLeft:
		nextCellX, nextCellY = x, y-1
	}

	return nextCellX, nextCellY
}

type NextCellResult struct {
	X, Y int
	Cell Cell
	IsGuardActive bool
	IsNextCellInBounds bool
}

// NextCell reveals the next cell that is in front of the guard
// with important metadata around it.
func (g *Grid) NextCell() NextCellResult {
	if !g.isGuardActive {
		return NextCellResult{
			IsGuardActive: false,
			IsNextCellInBounds: false,
		}
	}

	x, y := g.nextCellPosition()

	if x < 0 || y < 0 || x >= g.rows || y >= g.cols {
		return NextCellResult{
			IsGuardActive: true,
			IsNextCellInBounds: false,
		}
	}

	cell := g.data[x][y]

	return NextCellResult{
		X:             x,
		Y:             y,
		Cell:          cell,
		IsGuardActive: true,
		IsNextCellInBounds: true,
	}
}

// Move moves the guard to the next cell. It returns true if the guard moved
// false otherwise.
func (g *Grid) Move() bool {
	if !g.isGuardActive {
		return false
	}

	x, y := g.guardPosition[0], g.guardPosition[1]
	nx, ny := g.nextCellPosition()

	guard := g.data[x][y]
	g.data[x][y] = VisitedCell

	if nx < 0 || ny < 0 || nx >= g.rows || ny >= g.cols {
		g.isGuardActive = false
		return true
	}

	nc := g.data[nx][ny]

	if nc == BlockedCell {
		g.data[x][y] = guard
		return false
	}

	g.data[nx][ny] = guard
	g.guardPosition = [2]int{nx, ny}

	return true
}

// RotateRight rotates the guard if it is active. It does
// nothing otherwise. It will return true if the guard has
// been rotated, false otherwise.
func (g *Grid) RotateRight() bool {
	if !g.isGuardActive {
		return false
	}

	x, y := g.guardPosition[0], g.guardPosition[1]
	guard := g.data[x][y]

	switch guard {
	case GuardLeft:
		g.data[x][y] = GuardTop
	case GuardTop:
		g.data[x][y] = GuardRight
	case GuardRight:
		g.data[x][y] = GuardBottom
	case GuardBottom:
		g.data[x][y] = GuardLeft
	}

	return true
}

// VisitedCells returns the number of visited cells
func (g *Grid) VisitedCells() int {
	var result int

	for _, row := range g.data {
		for _, cell := range row {
			if cell == VisitedCell {
				result++
			}
		}
	}

	return result
}
