package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_IsNextCellOutsideGrid(t *testing.T) {
	tests := []struct {
		name string
		guardPosition [2]int
		direction Direction
		guardActive bool
		want bool
	}{
		{
			name: "OutsideGridTest",
			guardPosition: [2]int{0, 0},
			direction: DirectionUp,
			guardActive: true,
			want: true,
		},
		{
			name: "InsideGridTest",
			guardPosition: [2]int{0, 0},
			direction: DirectionRight,
			guardActive: true,
			want: false,
		},
		{
			name: "NoGuardTest",
			guardPosition: [2]int{0, 0},
			direction: DirectionRight,
			guardActive: false,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGrid(3, 3)
			g.guardPosition = tt.guardPosition
			g.direction = tt.direction
			g.guardActive = tt.guardActive

			assert.Equal(t, tt.want, g.IsNextCellOutsideGrid())
		})
	}
}

func Test_IsNextCellOutsideGrid_OutsideGridTest(t *testing.T) {
	g := newGrid(3, 3)
	g.guardPosition = [2]int{0, 0}
	g.direction = DirectionUp
	g.guardActive = true

	assert.True(t, g.IsNextCellOutsideGrid())
}

func Test_RemoveGuard(t *testing.T) {
	g := newGrid(3, 3)
	g.guardPosition = [2]int{0, 0}
	g.direction = DirectionRight
	g.guardActive = true

	g.RemoveGuard()

	assert.False(t, g.guardActive)
}

func Test_IsGuardActive(t *testing.T) {
	tests := []struct {
		name string
		isGuardActive bool
		want bool
	}{
		{
			name: "Active",
			isGuardActive: true,
			want: true,
		},
		{
			name: "Inactive",
			isGuardActive: false,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				guardActive: tt.isGuardActive,
			}
			assert.Equal(t, tt.want, g.IsGuardActive())
		})
	}
}


func Test_IsNextCellBlocked(t *testing.T) {
	tests := []struct {
		name string
		guardPosition [2]int
		direction Direction
		guardActive bool
		want bool
		blockedCells [][]bool
	}{
		{
			name: "BlockedTest",
			guardPosition: [2]int{0, 0},
			direction: DirectionRight,
			guardActive: true,
			want: true,
			blockedCells: [][]bool{
				{false, true, false},
				{false, false, false},
				{false, false, false},
			},
		},
		{
			name: "NotBlockedTest",
			guardPosition: [2]int{0, 0},
			direction: DirectionRight,
			guardActive: true,
			want: false,
			blockedCells: [][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGrid(3, 3)
			g.guardPosition = tt.guardPosition
			g.direction = tt.direction
			g.guardActive = tt.guardActive
			g.blockedCells = tt.blockedCells

			assert.Equal(t, tt.want, g.IsNextCellBlocked())
		})
	}
}

func Test_String(t *testing.T) {
	tests := []struct {
		name string
		rows int
		cols int
		guardPosition [2]int
		direction Direction
		guardActive bool
		blockedCells [][]bool
		data [][]int
		want string
	}{
		{
			name: "GuardTop",
			rows: 1,
			cols: 3,
			guardPosition: [2]int{0, 0},
			direction: DirectionUp,
			guardActive: true,
			blockedCells: [][]bool{
				{false, false, false},
			},
			data: [][]int{
				{0, 0, 0},
			},
			want: "^..\n",
		},
		{
			name: "GuardRight",
			rows: 1,
			cols: 3,
			guardPosition: [2]int{0, 0},
			direction: DirectionRight,
			guardActive: true,
			blockedCells: [][]bool{
				{false, false, false},
			},
			data: [][]int{
				{0, 0, 0},
			},
			want: ">..\n",
		},
		{
			name: "GuardDown",
			rows: 1,
			cols: 3,
			guardPosition: [2]int{0, 0},
			direction: DirectionDown,
			guardActive: true,
			blockedCells: [][]bool{
				{false, false, false},
			},
			data: [][]int{
				{0, 0, 0},
			},
			want: "v..\n",
		},
		{
			name: "GuardLeft",
			rows: 1,
			cols: 3,
			guardPosition: [2]int{0, 0},
			direction: DirectionLeft,
			guardActive: true,
			blockedCells: [][]bool{
				{false, false, false},
			},
			data: [][]int{
				{0, 0, 0},
			},
			want: "<..\n",
		},
		{
			name: "NoGuard",
			rows: 1,
			cols: 3,
			guardPosition: [2]int{0, 0},
			direction: DirectionRight,
			guardActive: false,
			blockedCells: [][]bool{
				{false, false, false},
			},
			data: [][]int{
				{0, 0, 0},
			},
			want: "...\n",
		},
		{
			name: "BlockedAndVisitedCells",
			rows: 1,
			cols: 3,
			guardPosition: [2]int{0, 0},
			direction: DirectionRight,
			guardActive: true,
			blockedCells: [][]bool{
				{false, true, false},
			},
			data: [][]int{
				{0, 0, 1},
			},
			want: ">#X\n",
		},
		{
			name: "MultipleLines",
			rows: 3,
			cols: 3,
			guardPosition: [2]int{0, 1},
			direction: DirectionUp,
			guardActive: true,
			blockedCells: [][]bool{
				{false, false, false},
				{true, false, false},
				{false, false, true},
			},
			data: [][]int{
				{0, 1, 1},
				{0, 1, 1},
				{0, 0, 0},
			},
			want: ".^X\n#XX\n..#\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGrid(tt.rows, tt.cols)
			g.guardPosition = tt.guardPosition
			g.direction = tt.direction
			g.guardActive = tt.guardActive
			g.blockedCells = tt.blockedCells
			g.data = tt.data

			assert.Equal(t, tt.want, g.String())
		})
	}
}
