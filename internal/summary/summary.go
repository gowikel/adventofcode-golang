package summary

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pterm/pterm"
	"go.eryndalor.dev/adventofcode-golang/internal/constants"
	"go.eryndalor.dev/adventofcode-golang/internal/runner"
)

type Summary struct {
	year  int
	day   int
	p1    int
	p1Err error
	p2    int
	p2Err error
}

func BuildSummary(
	year int,
	day int,
	p1 int,
	p1Err error,
	p2 int,
	p2Err error,
) Summary {
	return Summary{
		year:  year,
		day:   day,
		p1:    p1,
		p1Err: p1Err,
		p2:    p2,
		p2Err: p2Err,
	}
}

func (s Summary) RenderSummary() {
	s.renderStatus()
	s.renderExecutionTable()

	if errors.Is(s.p1Err, runner.ErrSolverNotImplemented) {
		s.renderSolverImplementationHelp()
	}
}

func (s Summary) renderStatus() {
	if errors.Is(s.p1Err, runner.ErrSolverNotImplemented) {
		pterm.DefaultBasicText.Printf(
			"Status: %s\n\n",
			constants.ERROR_BOX.Sprint("Solver not implemented"),
		)
	} else if (s.p1Err != nil && !errors.Is(s.p1Err, runner.ErrPartNotImplemented)) || (s.p2Err != nil && !errors.Is(s.p2Err, runner.ErrPartNotImplemented)) {
		pterm.DefaultBasicText.Printf(
			"Status: %s\n\n",
			constants.ERROR_BOX.Sprint("ERROR"),
		)
	} else if s.p1Err != nil || s.p2Err != nil {
		pterm.DefaultBasicText.Printf(
			"Status: %s\n\n",
			constants.WARNING_BOX.Sprint("IN PROGRESS"),
		)
	} else {
		pterm.DefaultBasicText.Printf(
			"Status: %s\n\n",
			constants.DONE_BOX.Sprint("DONE"),
		)
	}
}

func (s Summary) renderExecutionTable() {
	p1Value := strconv.Itoa(s.p1)
	p2Value := strconv.Itoa(s.p2)

	if s.p1Err != nil && errors.Is(s.p1Err, runner.ErrPartNotImplemented) {
		p1Value = constants.WARNING_BOX.Sprint("NOT IMPLEMENTED")
	} else if s.p1Err != nil {
		p1Value = constants.ERROR_BOX.Sprint(s.p1Err)
	}

	if s.p2Err != nil && errors.Is(s.p2Err, runner.ErrPartNotImplemented) {
		p2Value = constants.WARNING_BOX.Sprint("NOT IMPLEMENTED")
	} else if s.p2Err != nil {
		p2Value = constants.ERROR_BOX.Sprint(s.p2Err)
	}

	if !errors.Is(s.p1Err, runner.ErrSolverNotImplemented) {
		pterm.DefaultTable.WithHasHeader().WithData(
			pterm.TableData{
				{"P1", "P2"},
				{p1Value, p2Value},
			},
		).Render()
	}
}

func (s Summary) renderSolverImplementationHelp() {
	pterm.DefaultBasicText.Println(
		"Try to implement " + pterm.LightMagenta(
			fmt.Sprintf("year%d/day%02d", s.year, s.day),
		) + " first.",
	)
	pterm.DefaultBasicText.Println(
		"And don't forget to update the list of solvers at " + pterm.LightMagenta(
			fmt.Sprintf("year%d/runner.go", s.year),
		),
	)
}
