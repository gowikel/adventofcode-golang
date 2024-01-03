package main

import (
	"fmt"
	"os"

	"github.com/gowikel/adventofcode-golang/cli"
	"github.com/gowikel/adventofcode-golang/year2023"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	options := cli.ParseArgs()
	data := fetchData(options)

	fmt.Println("Year:", options.Year)
	fmt.Println("Day:", options.Day)
	fmt.Println("Running example?", options.RunExample)
	fmt.Println()

	// TODO: Will be updated to run other years in the future
	year2023.Run(options.Day, data)
}
