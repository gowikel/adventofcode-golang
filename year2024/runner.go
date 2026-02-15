package year2024

import (
	"go.eryndalor.dev/adventofcode-golang/internal/runner"
	D01 "go.eryndalor.dev/adventofcode-golang/year2024/day01"
	D02 "go.eryndalor.dev/adventofcode-golang/year2024/day02"
	D03 "go.eryndalor.dev/adventofcode-golang/year2024/day03"
	D04 "go.eryndalor.dev/adventofcode-golang/year2024/day04"
	D05 "go.eryndalor.dev/adventofcode-golang/year2024/day05"
	D06 "go.eryndalor.dev/adventofcode-golang/year2024/day06"
	D07 "go.eryndalor.dev/adventofcode-golang/year2024/day07"
	D08 "go.eryndalor.dev/adventofcode-golang/year2024/day08"
	D09 "go.eryndalor.dev/adventofcode-golang/year2024/day09"
)

var Solvers = map[int]runner.Solver{
	1: D01.Exercise{},
	2: D02.Exercise{},
	3: D03.Exercise{},
	4: D04.Exercise{},
	5: D05.Exercise{},
	6: D06.Exercise{},
	7: D07.Exercise{},
	8: D08.Exercise{},
	9: D09.Exercise{},
}
