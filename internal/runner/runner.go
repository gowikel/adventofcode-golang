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

func (r *Runner) Run(day int, data string) (p1 int, p2 int, err error) {
	solver, ok := r.solvers[day]

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
