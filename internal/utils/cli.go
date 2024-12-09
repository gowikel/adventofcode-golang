package utils

import (
	"strconv"

	"github.com/gowikel/adventofcode-golang/internal/constants"
	"github.com/pterm/pterm"
)

func PrintStatus(p1err error, p2err error) {
	if p1err != nil || p2err != nil {
		pterm.DefaultBasicText.Printf(
			"Status: %s\n\n",
			constants.ERROR_BOX.Sprint("ERROR"),
		)
	} else {
		pterm.DefaultBasicText.Printf(
			"Status: %s\n\n",
			constants.DONE_BOX.Sprint("DONE"),
		)
	}
}

func BuildExecutionTable(p1 int, p1err error, p2 int, p2err error) *pterm.TablePrinter {
	p1Value := strconv.Itoa(p1)
	p2Value := strconv.Itoa(p2)

	if p1err != nil {
		p1Value = constants.ERROR_BOX.Sprint(p1err)
	}

	if p2err != nil {
		p2Value = constants.ERROR_BOX.Sprint(p2err)
	}

	return pterm.DefaultTable.WithHasHeader().WithData(
		pterm.TableData{
			{"P1", "P2"},
			{p1Value, p2Value},
		},
	)
}
