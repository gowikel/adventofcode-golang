package cli

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Given a string, parse it as a given
// year for AoC. Then validate that the year
// is correct.
func ParseYear(y string) (int, error) {
	year, err := strconv.Atoi(y)
	if err != nil {
		return year, err
	}

	return year, validateYear(year)
}

// Given a string, parse it as an integer
// and validate that it is between 1 and 25.
func ParseDay(d string) (int, error) {
	day, err := strconv.Atoi(d)
	if err != nil {
		return day, err
	}

	return day, validateDay(day)
}

// Parses the given args, which are expected to have at least one
// element and at most two.
// If there are two elements, the first one will be interprted as
// the year.
// If there is one element, the year will be deducted from the current
// time. If we are on December, then the current year will be used,
// otherwise, the previous year will be used.
// Then, the element in args will be interpreted as the day
// If the len or the args is zero or greater than two, or any
// error happens while parsing, an error will be returned
func ParseYearAndDay(args []string) (year int, day int, err error) {
	if len(args) == 0 {
		err = errors.New("no args provided")
		return year, day, err
	} else if len(args) == 1 {
		// year is assumed to be the last year, except on December
		now := time.Now()
		year = now.Year() - 1

		if now.Month() == time.December {
			year += 1
		}

		day, err = ParseDay(args[0])
		if err != nil {
			return year, day, err
		}
	} else {
		year, err = ParseYear(args[0])
		if err != nil {
			return year, day, err
		}

		day, err = ParseDay(args[1])
		if err != nil {
			return year, day, err
		}
	}

	err = validateYear(year)
	if err != nil {
		return year, day, err
	}

	err = validateDay(day)
	if err != nil {
		return year, day, err
	}

	return year, day, nil
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
