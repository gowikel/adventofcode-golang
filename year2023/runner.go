package year2023

import (
	"fmt"

	D01 "github.com/gowikel/adventofcode-golang/year2023/day01"
	D02 "github.com/gowikel/adventofcode-golang/year2023/day02"
	D03 "github.com/gowikel/adventofcode-golang/year2023/day03"
	D04 "github.com/gowikel/adventofcode-golang/year2023/day04"
	D05 "github.com/gowikel/adventofcode-golang/year2023/day05"
	D06 "github.com/gowikel/adventofcode-golang/year2023/day06"
	D07 "github.com/gowikel/adventofcode-golang/year2023/day07"
	D08 "github.com/gowikel/adventofcode-golang/year2023/day08"
	D09 "github.com/gowikel/adventofcode-golang/year2023/day09"
)

type solver interface {
	Part1(data string) (int, error)
	Part2(data string) (int, error)
}

var solvers = map[int]solver{
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

func Run(day int, data string) (p1 int, p2 int, err error) {
	solver, ok := solvers[day]

	if !ok {
		err = fmt.Errorf("solver not implemented")
		return
	}

	p1, err = solver.Part1(data)
	if err != nil {
		return
	}

	p2, err = solver.Part2(data)
	return
}
