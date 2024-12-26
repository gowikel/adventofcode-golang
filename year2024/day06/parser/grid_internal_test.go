package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_String_Empty(t *testing.T) {
	g := newGrid(3, 3)
	expected := "...\n...\n..."
	got := g.String()

	assert.Equal(t, expected, got)
}

func TestGrid_NewGrid_Rows(t *testing.T) {
	g := newGrid(4, 5)
	expected := 4
	got := g.Rows()

	assert.Equal(t, expected, got)
}

func TestGrid_NewGrid_Columns(t *testing.T) {
	g := newGrid(4, 5)
	expected := 5
	got := g.Columns()

	assert.Equal(t, expected, got)
}

func TestGrid_NewGrid_IsGuardActive_False(t *testing.T) {
	g := newGrid(3, 3)
	expected := false
	got := g.isGuardActive

	assert.Equal(t, expected, got)
}

func TestGrid_NewGrid_InitialGuardPosition(t *testing.T) {
	g := newGrid(3, 3)
	expected := [2]int{-1, -1}
	got := g.guardPosition

	assert.Equal(t, expected, got)
}

func TestGrid_GuardPosition_InBounds(t *testing.T) {
	g := newGrid(3, 3)
	g.isGuardActive = true
	g.guardPosition = [2]int{1, 2}

	expectedGuardPosition := GuardPositionInsideGrid
	expectedCoordinates := [2]int{1, 2}
	gotGuardPosition, gotCoordinates := g.GuardPosition()

	assert.Equal(t, expectedGuardPosition, gotGuardPosition)
	assert.Equal(t, expectedCoordinates, gotCoordinates)
}

func TestGrid_GuardPosition_OutBounds(t *testing.T) {
	g := newGrid(3, 3)
	g.isGuardActive = false
	g.guardPosition = [2]int{2, 1}

	expectedGuardPosition := GuardPositionOutsideGrid
	expectedCoordinates := [2]int{2, 1}
	gotGuardPosition, gotCoordinates := g.GuardPosition()

	assert.Equal(t, expectedGuardPosition, gotGuardPosition)
	assert.Equal(t, expectedCoordinates, gotCoordinates)
}

func TestGrid_NextCell(t *testing.T) {
	tests := []struct {
		name                string
		setupGrid           func() *Grid
		expectedCell        FutureCell
		expectedGuardActive bool
	}{
		{
			name: "GuardTop_Move",
			setupGrid: func() *Grid {
				// ...
				// #^#
				// .#.
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[1][0] = BlockedCell
				g.data[1][1] = GuardTop
				g.data[1][2] = BlockedCell
				g.data[2][1] = BlockedCell
				return g
			},
			expectedCell:        FutureCellEmpty,
			expectedGuardActive: true,
		},
		{
			name: "GuardTop_Blocked",
			setupGrid: func() *Grid {
				// .#.
				// .^.
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[0][1] = BlockedCell
				g.data[1][1] = GuardTop
				return g
			},
			expectedCell:        FutureCellBlocked,
			expectedGuardActive: true,
		},
		{
			name: "GuardTop_Visited",
			setupGrid: func() *Grid {
				// .X.
				// .^.
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[0][1] = VisitedCell
				g.data[1][1] = GuardTop
				return g
			},
			expectedCell:        FutureCellEmpty,
			expectedGuardActive: true,
		},
		{
			name: "GuardTop_OutBounds",
			setupGrid: func() *Grid {
				// .^.
				// ...
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{0, 1}
				g.data[0][1] = GuardTop
				return g
			},
			expectedCell:        FutureCellOutsideGrid,
			expectedGuardActive: true,
		},
		{
			name: "GuardInactive",
			setupGrid: func() *Grid {
				// ...
				// ...
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = false
				g.guardPosition = [2]int{0, 1}
				return g
			},
			expectedCell:        FutureCellOutsideGrid,
			expectedGuardActive: false,
		},
		{
			name: "GuardRight_Move",
			setupGrid: func() *Grid {
				// .#.
				// #>.
				// .#.
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[0][1] = BlockedCell
				g.data[1][1] = GuardRight
				g.data[2][1] = BlockedCell
				return g
			},
			expectedCell:        FutureCellEmpty,
			expectedGuardActive: true,
		},
		{
			name: "GuardRight_Blocked",
			setupGrid: func() *Grid {
				// ...
				// .>#
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[1][1] = GuardRight
				g.data[1][2] = BlockedCell
				return g
			},
			expectedCell:        FutureCellBlocked,
			expectedGuardActive: true,
		},
		{
			name: "GuardRight_Visited",
			setupGrid: func() *Grid {
				// ...
				// .x>
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[1][1] = GuardRight
				g.data[2][1] = VisitedCell
				return g
			},
			expectedCell:        FutureCellEmpty,
			expectedGuardActive: true,
		},
		{
			name: "GuardRight_OutBounds",
			setupGrid: func() *Grid {
				// ...
				// ..>
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 2}
				g.data[1][2] = GuardRight
				return g
			},
			expectedCell:        FutureCellOutsideGrid,
			expectedGuardActive: true,
		},
		{
			name: "GuardBottom_Move",
			setupGrid: func() *Grid {
				// .#.
				// #v#
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[0][1] = BlockedCell
				g.data[1][0] = BlockedCell
				g.data[1][1] = GuardBottom
				g.data[1][2] = BlockedCell
				return g
			},
			expectedCell:        FutureCellEmpty,
			expectedGuardActive: true,
		},
		{
			name: "GuardBottom_Blocked",
			setupGrid: func() *Grid {
				// ...
				// .v.
				// .#.
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[1][1] = GuardBottom
				g.data[2][1] = BlockedCell
				return g
			},
			expectedCell:        FutureCellBlocked,
			expectedGuardActive: true,
		},
		{
			name: "GuardBottom_Visited",
			setupGrid: func() *Grid {
				// ...
				// .x.
				// .v.
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[1][1] = GuardBottom
				g.data[2][1] = VisitedCell
				return g
			},
			expectedCell:        FutureCellEmpty,
			expectedGuardActive: true,
		},
		{
			name: "GuardBottom_OutBounds",
			setupGrid: func() *Grid {
				// ...
				// ...
				// .v.
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{2, 1}
				g.data[2][1] = GuardBottom
				return g
			},
			expectedCell:        FutureCellOutsideGrid,
			expectedGuardActive: true,
		},
		{
			name: "GuardLeft_Move",
			setupGrid: func() *Grid {
				// .#.
				// .<#
				// .#.
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[0][1] = BlockedCell
				g.data[1][1] = GuardLeft
				g.data[1][2] = BlockedCell
				g.data[2][1] = BlockedCell
				return g
			},
			expectedCell:        FutureCellEmpty,
			expectedGuardActive: true,
		},
		{
			name: "GuardLeft_Blocked",
			setupGrid: func() *Grid {
				// ...
				// #<.
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 1}
				g.data[1][1] = GuardLeft
				g.data[1][0] = BlockedCell
				return g
			},
			expectedCell:        FutureCellBlocked,
			expectedGuardActive: true,
		},
		{
			name: "GuardLeft_OutBounds",
			setupGrid: func() *Grid {
				// ...
				// <..
				// ...
				g := newGrid(3, 3)
				g.isGuardActive = true
				g.guardPosition = [2]int{1, 0}
				g.data[1][0] = GuardLeft
				return g
			},
			expectedCell:        FutureCellOutsideGrid,
			expectedGuardActive: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.setupGrid()
			resultCell, resultGuardActive := g.NextCell()
			assert.Equal(t, tt.expectedCell, resultCell)
			assert.Equal(t, tt.expectedGuardActive, resultGuardActive)
		})
	}
}

func Test_Move_GuardTop(t *testing.T) {
	// ...
	// .^.
	// ...
	g := newGrid(3, 3)
	g.isGuardActive = true
	g.guardPosition = [2]int{1, 1}
	g.data[1][1] = GuardTop

	r := g.Move()

	assert.Equal(t, [2]int{0, 1}, g.guardPosition)
	assert.Equal(t, true, g.isGuardActive)
	assert.Equal(t, GuardTop, g.data[0][1])
	assert.Equal(t, VisitedCell, g.data[1][1])
	assert.Equal(t, true, r)
}

func Test_Move_GuardTop_Blocked(t *testing.T) {
	// .#.
	// .^.
	// ...
	g := newGrid(3, 3)
	g.isGuardActive = true
	g.guardPosition = [2]int{1, 1}
	g.data[0][1] = BlockedCell
	g.data[1][1] = GuardTop

	r := g.Move()

	assert.Equal(t, [2]int{1, 1}, g.guardPosition)
	assert.Equal(t, true, g.isGuardActive)
	assert.Equal(t, GuardTop, g.data[1][1])
	assert.Equal(t, false, r)
}

func Test_Move_GuardTop_OutBounds(t *testing.T) {
	// .^.
	// ...
	// ...
	g := newGrid(3, 3)
	g.isGuardActive = true
	g.guardPosition = [2]int{0, 1}
	g.data[0][1] = GuardTop

	r := g.Move()

	assert.Equal(t, [2]int{0, 1}, g.guardPosition)
	assert.Equal(t, false, g.isGuardActive)
	assert.Equal(t, VisitedCell, g.data[0][1])
	assert.Equal(t, true, r)
}

func Test_Move_NoGuard(t *testing.T) {
	g := newGrid(3, 3)
	g.isGuardActive = false
	g.guardPosition = [2]int{0, 0}

	r := g.Move()

	assert.Equal(t, [2]int{0, 0}, g.guardPosition)
	assert.Equal(t, false, g.isGuardActive)
	assert.Equal(t, false, r)
}

func Test_Move_NoGuard_InvalidPosition(t *testing.T) {
	g := newGrid(3, 3)
	g.isGuardActive = false
	g.guardPosition = [2]int{-1, -1}

	r := g.Move()

	assert.Equal(t, [2]int{-1, -1}, g.guardPosition)
	assert.Equal(t, false, g.isGuardActive)
	assert.Equal(t, false, r)
}
