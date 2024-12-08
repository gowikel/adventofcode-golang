package year2024

import "fmt"

type solver interface {
	Part1(data string) (int, error)
	Part2(data string) (int, error)
}

var solvers = map[int]solver{}

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
