package year2023

import (
	"fmt"
	"os"

	D01 "github.com/gowikel/adventofcode-golang/year2023/day01"
	D02 "github.com/gowikel/adventofcode-golang/year2023/day02"
	D03 "github.com/gowikel/adventofcode-golang/year2023/day03"
	D04 "github.com/gowikel/adventofcode-golang/year2023/day04"
)

func Run(day int, data string) {
	switch day {
	case 1:
		D01.Solve(data)
	case 2:
		D02.Solve(data)
	case 3:
		D03.Solve(data)
	case 4:
		D04.Solve(data)
	default:
		fmt.Fprintf(os.Stderr, "Not implemented\n")
		os.Exit(125)
	}
}