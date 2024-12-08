package main

import (
	"fmt"
	"os"

	"github.com/gowikel/adventofcode-golang/internal/conf"
	"github.com/gowikel/adventofcode-golang/internal/constants"
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
		p1, p2, err := runner.Run(opts.Day, data)

		spinner.Stop()

		if err != nil {
			pterm.DefaultBasicText.Printf(
				"Status: %s\n\n",
				constants.ERROR_BOX.Sprint("ERROR"),
			)
			pterm.DefaultBasicText.Printf("%s\n\n", err.Error())
			return
		}

		pterm.DefaultBasicText.Printf(
			"Status: %s\n\n",
			constants.DONE_BOX.Sprint("DONE"),
		)

		pterm.DefaultTable.WithHasHeader().WithData(
			pterm.TableData{
				{"P1", "P2"},
				{fmt.Sprint(p1), fmt.Sprint(p2)},
			},
		).Render()
	})()
}
