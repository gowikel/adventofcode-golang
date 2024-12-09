package runner

import "fmt"

type Solver interface {
	Part1(data string) (int, error)
	Part2(data string) (int, error)
}

type Runner struct {
	solvers map[int]Solver
}

func New(solvers map[int]Solver) *Runner {
	return &Runner{
		solvers: solvers,
	}
}

func (r *Runner) RunPart1(day int, data string) (p1 int, err error) {
	solver, ok := r.solvers[day]

	if !ok {
		return 0, fmt.Errorf("solver not implemented")
	}

	return solver.Part1(data)
}

func (r *Runner) RunPart2(day int, data string) (p2 int, err error) {
	solver, ok := r.solvers[day]

	if !ok {
		return 0, fmt.Errorf("solver not implemented")
	}

	return solver.Part2(data)
}
