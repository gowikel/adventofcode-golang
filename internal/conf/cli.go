package conf

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Configuration struct {
	Year  int
	Day   int
	Input string
}

var conf *Configuration

// Given a year, y, validates that it is in a valid range
func validateYear(y int) error {
	minYear, maxYear := 2023, 2023

	if y < minYear || y > maxYear {
		return fmt.Errorf(
			"\"%d\" is not in the range [%d-%d]",
			y,
			minYear,
			maxYear,
		)
	}

	return nil
}

// Given a day, d, validate that it is between 1 and 25
func validateDay(d int) error {
	minDay, maxDay := 1, 25

	if d < minDay || d > maxDay {
		return fmt.Errorf(
			"\"%d\" is not in the range [%d-%d]",
			d,
			minDay,
			maxDay,
		)
	}

	return nil
}

func ParseCLI() {
	if conf != nil {
		return
	}

	result := Configuration{}

	now := time.Now()
	defaultYear := now.Year()
	defaultDay := now.Day()

	if now.Month() != time.December {
		defaultYear -= 1
	}

	flag.IntVar(
		&result.Year,
		"year",
		0,
		"Year to run (defaults to current on December, otherwise to previous year)",
	)

	flag.IntVar(
		&result.Day,
		"day",
		0,
		"Day to run (defaults to current on December, required otherwise)",
	)

	flag.StringVar(
		&result.Input,
		"input",
		"",
		"Input file to pass to the solver (required)",
	)

	flag.Parse()

	if result.Year == 0 {
		result.Year = defaultYear
	}

	if result.Day == 0 && now.Month() == time.December {
		result.Day = defaultDay
	} else if result.Day == 0 {
		fmt.Fprintln(os.Stderr, "Day is required")
		os.Exit(1)
	}

	err := validateYear(result.Year)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid year: %s\n", err)
		os.Exit(1)
	}

	err = validateDay(result.Day)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid day: %s\n", err)
		os.Exit(1)
	}

	conf = &result
}

func Conf() *Configuration {
	return conf
}
