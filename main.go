package main

import (
	"fmt"

	"github.com/gowikel/adventofcode-golang/internal/cli"
	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	"github.com/gowikel/adventofcode-golang/internal/utils"
	"github.com/gowikel/adventofcode-golang/year2023"
	"github.com/rs/zerolog/log"
)

func main() {
	// cmd.Execute()
	opts := cli.ParseFlags()

	fmt.Println("Year:", opts.Year)
	fmt.Println("Day:", opts.Day)
	fmt.Println(opts.Part)
	fmt.Println()

	data, err := puzzle.Read(opts.Input)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to read puzzle input")
	}

	utils.MeasureExecutionTime(func() {
		// TODO: Will be updated to run other years in the future
		year2023.Run(opts.Day, data, opts.Part)
	})()
}
