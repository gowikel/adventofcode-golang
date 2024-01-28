package conf

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gowikel/adventofcode-golang/internal/puzzle"
)

type Configuration struct {
	Year     int
	Day      int
	Input    string
	Part     puzzle.PuzzleRunSelector
	LogLevel slog.Leveler
}

var conf *Configuration

// Given an integer, converts it to puzzlePartSelector.PuzzlePart
func ParsePart(part int) (puzzle.PuzzleRunSelector, error) {
	switch part {
	// Default value, assume RunAll
	case 0:
		return puzzle.RunAll, nil
	case 1:
		return puzzle.RunPartOne, nil
	case 2:
		return puzzle.RunPartTwo, nil
	}

	return puzzle.RunAll, fmt.Errorf(
		"invalid part: %v",
		part,
	)
}

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
	var part int

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

	flag.IntVar(
		&part,
		"part",
		0,
		"Part to run. (Default both)",
	)

	flag.Parse()

	result.LogLevel = slog.LevelInfo

	if result.Year == 0 {
		result.Year = defaultYear
	}

	if result.Day == 0 && now.Month() == time.December {
		result.Day = defaultDay
	} else if result.Day == 0 {
		slog.Error("Day is required")
		os.Exit(1)
	}

	err := validateYear(result.Year)
	if err != nil {
		slog.Error("Invalid year", "err", err)
		os.Exit(1)
	}

	err = validateDay(result.Day)
	if err != nil {
		slog.Error("Invalid day", "err", err)
		os.Exit(1)
	}

	parsedPart, err := ParsePart(part)
	if err != nil {
		slog.Error("Invalid part", "err", err)
		os.Exit(1)
	}

	result.Part = parsedPart

	conf = &result
}

func Conf() *Configuration {
	return conf
}
