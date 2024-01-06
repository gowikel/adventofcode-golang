package puzzlePartSelector

// Enum to select which part to run
type PuzzlePart int

const (
	RunAll PuzzlePart = iota
	RunPartOne
	RunPartTwo
)

func (pps PuzzlePart) String() string {
	if pps == RunAll {
		return "Run All"
	} else if pps == RunPartOne {
		return "Run Part One"
	}
	return "Run Part Two"
}
