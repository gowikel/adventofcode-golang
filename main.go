package main

import (
	"os"

	"github.com/gowikel/adventofcode-golang/internal/conf"
	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	"github.com/gowikel/adventofcode-golang/internal/runner"
	"github.com/gowikel/adventofcode-golang/internal/utils"
	"github.com/gowikel/adventofcode-golang/year2023"
	"github.com/gowikel/adventofcode-golang/year2024"
	"github.com/pterm/pterm"
)

func main() {
	var solvers map[int]runner.Solver
	conf.ParseCLI()
	opts := conf.Conf()

	data, err := puzzle.Read(opts.Input)
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}

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

	runner := runner.New(solvers)

	utils.MeasureExecutionTime(func() {
		p1, p1err := runner.RunPart1(opts.Day, data)
		p2, p2err := runner.RunPart2(opts.Day, data)

		spinner.Stop()

		utils.PrintStatus(p1err, p2err)

		table := utils.BuildExecutionTable(p1, p1err, p2, p2err)
		table.Render()
	})()
}
