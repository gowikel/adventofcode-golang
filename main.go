package main

import (
	"github.com/gowikel/adventofcode-golang/internal/conf"
	"github.com/gowikel/adventofcode-golang/internal/runner"
	"github.com/gowikel/adventofcode-golang/internal/summary"
	"github.com/gowikel/adventofcode-golang/internal/utils"
	"github.com/gowikel/adventofcode-golang/year2023"
	"github.com/gowikel/adventofcode-golang/year2024"
	"github.com/pterm/pterm"
)

func main() {
	var solvers map[int]runner.Solver
	conf.ParseCLI()
	opts := conf.Conf()

	pterm.DefaultSection.Println("Advent of Code")

	pterm.Printf("Year: %d\n", opts.Year)
	pterm.Printf("Day: %d\n", opts.Day)

	spinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone().
		Start("Running solver...")

	switch opts.Year {
	case 2023:
		solvers = year2023.Solvers

	case 2024:
		solvers = year2024.Solvers
	}

	r := runner.New(solvers)

	utils.MeasureExecutionTime(func() {
		p1, p1err := r.RunPart1(opts.Day, opts.Input)
		p2, p2err := r.RunPart2(opts.Day, opts.Input)

		spinner.Stop()

		summary := summary.BuildSummary(
			opts.Year,
			opts.Day,
			p1,
			p1err,
			p2,
			p2err,
		)

		summary.RenderSummary()
	})()
}
