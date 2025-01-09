package day06

import "github.com/gowikel/adventofcode-golang/year2024/day06/parser"

var blockers [][]int = [][]int{
	{6, 3},
	{7, 6},
	{7, 7},
	{8, 1},
	{8, 3},
	{9, 7},
}

// Just move the grid until the guard is not active
func iterateGrid(grid *parser.Grid) {
	for grid.IsGuardActive() {
		if grid.IsNextCellBlocked() {
			grid.RotateRight()
		}
		grid.Move()
	}
}

func initDirectionMap(rows, cols int) [][][4]bool {
	result := make([][][4]bool, rows)

	for i := 0; i < rows; i++ {
		result[i] = make([][4]bool, cols)
	}

	return result
}

// Iterate the grid until a cycle is detected or the guard is outside the grid.
// Returns true if the cycle is detected, false otherwise.
func iterateGridWithCycleDetection(grid *parser.Grid) bool {
	rows := grid.Rows()
	cols := grid.Cols()

	const (
		TOP = 0
		RIGHT = 1
		BOTTOM = 2
		LEFT = 3
	)

	directionMap := initDirectionMap(rows, cols)

	for grid.IsGuardActive() {
		if grid.IsNextCellBlocked() {
			grid.RotateRight()
		}
		grid.Move()

		guardPosition, isActive := grid.GetCurrentGuardPosition()
		guardDirection := grid.GetGuardDirection()

		if !isActive {
			return false
		}

		x, y := guardPosition[0], guardPosition[1]
		var cellEntrance int

		switch guardDirection {
			// Entrance from the left
			case parser.DirectionRight:
				cellEntrance = LEFT
			// Entrance from the right
			case parser.DirectionLeft:
				cellEntrance = RIGHT
			// Entrance from the top
			case parser.DirectionDown:
				cellEntrance = TOP
			// Entrance from the bottom
			case parser.DirectionUp:
				cellEntrance = BOTTOM
		}

		if directionMap[x][y][cellEntrance] {
			return true
		}

		directionMap[x][y][cellEntrance] = true
	}

	return false
}
