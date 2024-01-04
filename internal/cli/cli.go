package cli

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func ParseArgs() {
	var year, day int
	var runExample, debug bool

	current := time.Now()
	defaultYear := current.Year()
	defaultDay := current.Day()

	minYear, maxYear := 2023, 2023

	if current.Month() != time.December {
		defaultDay = 1
	}

	pflag.IntVar(&year, "year", defaultYear, "Year to run")
	pflag.IntVar(&day, "day", defaultDay, "Day to run")
	pflag.BoolVar(
		&runExample,
		"run-example",
		false,
		"Execute with the example output",
	)
	pflag.BoolVar(
		&debug,
		"debug",
		false,
		"Enables the debug log level",
	)

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if year < minYear || year > maxYear {
		log.Fatal().
			Msgf("Year must be between [%d-%d]", minYear, maxYear)
	}

	if day < 1 || day > 25 {
		log.Fatal().Msgf("Day must be between [1-25]")
	}
}
