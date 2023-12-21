package year2023

import (
	"fmt"
	"os"

	D01 "github.com/gowikel/adventofcode-golang/year2023/day01"
	D02 "github.com/gowikel/adventofcode-golang/year2023/day02"
	D03 "github.com/gowikel/adventofcode-golang/year2023/day03"
	D04 "github.com/gowikel/adventofcode-golang/year2023/day04"
	D05 "github.com/gowikel/adventofcode-golang/year2023/day05"
)

type Solver interface {
	Solve(data string)
}

var solvers = map[int]Solver{
	1: D01.Exercise{},
	2: D02.Exercise{},
	3: D03.Exercise{},
	4: D04.Exercise{},
	5: D05.Exercise{},
}

func Run(day int, data string) {
	solver, ok := solvers[day]

	if !ok {
		fmt.Fprintf(os.Stderr, "Not implemented\n")
		os.Exit(125)
	}

	solver.Solve(data)
}