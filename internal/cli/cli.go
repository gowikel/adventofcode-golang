package cli

import (
	"flag"
	"time"

	"github.com/rs/zerolog/log"
)

type AOCOptions struct {
	Year       int
	Day        int
	RunExample bool
}

func ParseArgs() AOCOptions {
	result := AOCOptions{}

	current := time.Now()
	defaultYear := current.Year()
	defaultDay := current.Day()

	minYear, maxYear := 2023, 2023

	if current.Month() != time.December {
		defaultDay = 1
	}

	flag.IntVar(&result.Year, "year", defaultYear, "Year to run")
	flag.IntVar(&result.Day, "day", defaultDay, "Day to run")
	flag.BoolVar(
		&result.RunExample,
		"run-example",
		false,
		"Execute with the example output",
	)
	flag.Parse()

	if result.Year < minYear || result.Year > maxYear {
		log.Fatal().
			Msgf("Year must be between [%d-%d]", minYear, maxYear)
	}

	if result.Day < 1 || result.Day > 25 {
		log.Fatal().Msgf("Day must be between [1-25]")
	}

	return result
}
