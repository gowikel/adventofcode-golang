package main

import (
	"fmt"
	"os"

	"github.com/gowikel/adventofcode-golang/internal/conf"
	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	"github.com/gowikel/adventofcode-golang/internal/utils"
	"github.com/gowikel/adventofcode-golang/year2023"
)

func main() {
	conf.ParseCLI()
	opts := conf.Conf()

	data, err := puzzle.Read(opts.Input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Running exercise")
	fmt.Printf("  Year: %d\n", opts.Year)
	fmt.Printf("  Day: %d\n", opts.Day)
	fmt.Printf("  Part: %s\n", opts.Part)

	utils.MeasureExecutionTime(func() {
		// TODO: Will be updated to run other years in the future
		year2023.Run(opts.Day, data, opts.Part)
	})()
}
