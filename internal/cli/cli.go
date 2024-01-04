package cli

import (
	"fmt"
	"strconv"
)

func ParseYear(y string) (int, error) {
	minYear, maxYear := 2023, 2023
	year, err := strconv.Atoi(y)
	if err != nil {
		return year, err
	}

	if year < minYear || year > maxYear {
		return year, fmt.Errorf(
			"%q is not in the range [%d-%d]",
			y,
			minYear,
			maxYear,
		)
	}

	return year, nil
}

func ParseDay(d string) (int, error) {
	minDay, maxDay := 1, 25
	day, err := strconv.Atoi(d)
	if err != nil {
		return day, err
	}

	if day < minDay || day > maxDay {
		return day, fmt.Errorf(
			"%q is not in the range [%d-%d]",
			d,
			minDay,
			maxDay,
		)
	}

	return day, nil
}
