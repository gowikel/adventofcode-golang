// Package runner provides functionality to execute Advent of Code solutions.
// It defines interfaces and types to standardize how daily puzzle solutions are run.

package runner

import (
	"errors"
	"fmt"
)

// Solver is the interface that must be implemented by each day's solution.
// It defines methods to solve both parts of an Advent of Code puzzle.
type Solver interface {
	// Part1 solves the first part of the puzzle and returns the result.
	// It should return ErrSolverNotImplemented if the part is not yet implemented
	Part1(data string) (int, error)
	// Part2 solves the second part of the puzzle and returns the result.
	// It should return ErrSolverNotImplemented if the part is not yet implemented
	Part2(data string) (int, error)
}

// Runner manages a collection of Solvers and provides methods to execute them.
type Runner struct {
	solvers map[int]Solver
}

// ErrSolverNotImplemented is returned when attempting to run a solution for a day
// that has not been implemented.
type ErrSolverNotImplemented struct {
	day int
}

func (e ErrSolverNotImplemented) Error() string {
	return fmt.Sprintf("solver for day %d not implemented", e.day)
}

// ErrPartNotImplemented is returned when a specific part (1 or 2) of a day's solution
// has not been implemented.
var ErrPartNotImplemented = errors.New("not implemented")

// New creates a new Runner instance with the provided map of solvers.
func New(solvers map[int]Solver) *Runner {
	return &Runner{
		solvers: solvers,
	}
}

// RunPart1 executes the Part1 solution for the specified day with the given input data.
// Returns ErrSolverNotImplemented if no solver exists for the specified day.
func (r *Runner) RunPart1(day int, data string) (p1 int, err error) {
	solver, ok := r.solvers[day]

	if !ok {
		return 0, ErrSolverNotImplemented{day}
	}

	return solver.Part1(data)
}

// RunPart2 executes the Part2 solution for the specified day with the given input data.
// Returns ErrSolverNotImplemented if no solver exists for the specified day.
func (r *Runner) RunPart2(day int, data string) (p2 int, err error) {
	solver, ok := r.solvers[day]

	if !ok {
		return 0, fmt.Errorf("solver not implemented")
	}

	return solver.Part2(data)
}
