package cli

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/gowikel/adventofcode-golang/internal/puzzlePartSelector"
)

type CLIOptions struct {
	Year  int
	Day   int
	Input string
	Part  puzzlePartSelector.PuzzlePart
}

// Given a string, parse it as a given
// year for AoC. Then validate that the year
// is correct.
func ParseYear(y string) (int, error) {
	year, err := strconv.Atoi(y)
	if err != nil {
		return year, fmt.Errorf("ParseYear: %w", err)
	}

	err = validateYear(year)
	if err != nil {
		err = fmt.Errorf("ParseYear: %w", err)
	}

	return year, err
}

// Given a string, parse it as an integer
// and validate that it is between 1 and 25.
func ParseDay(d string) (int, error) {
	day, err := strconv.Atoi(d)
	if err != nil {
		return day, fmt.Errorf("ParseDay: %w", err)
	}

	err = validateDay(day)
	if err != nil {
		err = fmt.Errorf("ParseDay: %w", err)
	}

	return day, err
}

// Given an integer, converts it to puzzlePartSelector.PuzzlePart
func ParsePart(part int) (puzzlePartSelector.PuzzlePart, error) {
	switch part {
	// Default value, assume RunAll
	case 0:
		return puzzlePartSelector.RunAll, nil
	case 1:
		return puzzlePartSelector.RunPartOne, nil
	case 2:
		return puzzlePartSelector.RunPartTwo, nil
	}

	return puzzlePartSelector.RunAll, fmt.Errorf(
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

func ParseFlags() CLIOptions {
	result := CLIOptions{}
	var part int

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

	return result
}
