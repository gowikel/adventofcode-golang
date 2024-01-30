package constants

import "github.com/pterm/pterm"

var ERROR_BOX = pterm.NewStyle(
	pterm.BgLightRed,
	pterm.FgBlack,
	pterm.Bold,
)

var DONE_BOX = pterm.NewStyle(
	pterm.BgLightGreen,
	pterm.FgBlack,
	pterm.Bold,
)
