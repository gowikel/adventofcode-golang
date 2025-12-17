package year2024

import (
	"github.com/gowikel/adventofcode-golang/internal/runner"
	D01 "github.com/gowikel/adventofcode-golang/year2024/day01"
	D02 "github.com/gowikel/adventofcode-golang/year2024/day02"
	D03 "github.com/gowikel/adventofcode-golang/year2024/day03"
	D04 "github.com/gowikel/adventofcode-golang/year2024/day04"
	D05 "github.com/gowikel/adventofcode-golang/year2024/day05"
	D06 "github.com/gowikel/adventofcode-golang/year2024/day06"
	D07 "github.com/gowikel/adventofcode-golang/year2024/day07"
)

var Solvers = map[int]runner.Solver{
	1: D01.Exercise{},
	2: D02.Exercise{},
	3: D03.Exercise{},
	4: D04.Exercise{},
	5: D05.Exercise{},
	6: D06.Exercise{},
	7: D07.Exercise{},
}
