package year2023

import (
	"log"

	D01 "github.com/gowikel/adventofcode-golang/year2023/day01"
	D02 "github.com/gowikel/adventofcode-golang/year2023/day02"
	D03 "github.com/gowikel/adventofcode-golang/year2023/day03"
	D04 "github.com/gowikel/adventofcode-golang/year2023/day04"
	D05 "github.com/gowikel/adventofcode-golang/year2023/day05"
	D06 "github.com/gowikel/adventofcode-golang/year2023/day06"
)

type solver interface {
	Solve(data string)
}

var solvers = map[int]solver{
	1: D01.Exercise{},
	2: D02.Exercise{},
	3: D03.Exercise{},
	4: D04.Exercise{},
	5: D05.Exercise{},
	6: D06.Exercise{},
}

func Run(day int, data string) {
	solver, ok := solvers[day]

	if !ok {
		log.Fatal("Not implemented")
	}

	solver.Solve(data)
}
