package day06

import (
	"fmt"
	"slices"
	"strings"

	"errors"
)

var ErrParse = errors.New("error while parsing file")

func Parse(data string) ([]Race, error) {
	result := make([]Race, 0)

	lines := strings.Split(strings.TrimSpace(data), "\n")
	if len(lines) != 2 {
		return result, fmt.Errorf(
			"%w: it was expected to find two lines, but %d where found instead",
			ErrParse,
			len(lines),
		)
	}

	timeLine := lines[0]
	distanceLine := lines[1]

	if !strings.HasPrefix(timeLine, "Time:") {
		return result, fmt.Errorf(
			"%w: \"Time:\" not found in %q",
			ErrParse,
			timeLine,
		)
	}

	if !strings.HasPrefix(distanceLine, "Distance:") {
		return result, fmt.Errorf(
			"%w: \"Distance:\" not found in %q",
			ErrParse,
			distanceLine,
		)
	}

	timeLine = strings.TrimPrefix(timeLine, "Time:")
	timeLine = strings.TrimSpace(timeLine)

	distanceLine = strings.TrimPrefix(distanceLine, "Distance:")
	distanceLine = strings.TrimSpace(distanceLine)

	timeFields := strings.Fields(timeLine)
	distanceFields := strings.Fields(distanceLine)

	if len(timeFields) != len(distanceFields) {
		return result, fmt.Errorf(
			"%w: found %d time fields with %d distance fields",
			ErrParse,
			len(timeFields),
			len(distanceFields),
		)
	}

	result = slices.Grow[[]Race](result, len(timeFields))

	for i := 0; i < len(timeFields); i++ {
		timeText := timeFields[i]
		distanceText := distanceFields[i]

		var time, distance int

		_, err := fmt.Sscanf(timeText, "%d", &time)

		if err != nil {
			return result, fmt.Errorf(
				"%w: error while parsing the time %q: %w",
				ErrParse,
				timeText,
				err,
			)
		}

		_, err = fmt.Sscanf(distanceText, "%d", &distance)

		if err != nil {
			return result, fmt.Errorf(
				"%w: error while parsing the distance %q: %w",
				ErrParse,
				distanceText,
				err,
			)
		}

		result = append(result, NewRace(time, distance))
	}

	return result, nil
}

func ParsePart2(data string) (Race, error) {
	result := Race{}

	lines := strings.Split(strings.TrimSpace(data), "\n")
	if len(lines) != 2 {
		return result, fmt.Errorf(
			"%w: it was expected to find two lines, but %d where found instead",
			ErrParse,
			len(lines),
		)
	}

	timeLine := lines[0]
	distanceLine := lines[1]

	if !strings.HasPrefix(timeLine, "Time:") {
		return result, fmt.Errorf(
			"%w: \"Time:\" not found in %q",
			ErrParse,
			timeLine,
		)
	}

	if !strings.HasPrefix(distanceLine, "Distance:") {
		return result, fmt.Errorf(
			"%w: \"Distance:\" not found in %q",
			ErrParse,
			distanceLine,
		)
	}

	timeLine = strings.TrimPrefix(timeLine, "Time:")
	timeLine = strings.TrimSpace(timeLine)
	timeLine = strings.ReplaceAll(timeLine, " ", "")

	distanceLine = strings.TrimPrefix(distanceLine, "Distance:")
	distanceLine = strings.TrimSpace(distanceLine)
	distanceLine = strings.ReplaceAll(distanceLine, " ", "")

	var time, distance int

	_, err := fmt.Sscanf(timeLine, "%d", &time)

	if err != nil {
		return result, fmt.Errorf(
			"%w: error while parsing the time %q: %w",
			ErrParse,
			timeLine,
			err,
		)
	}

	result.Time = time

	_, err = fmt.Sscanf(distanceLine, "%d", &distance)

	if err != nil {
		return result, fmt.Errorf(
			"%w: error while parsing the distance %q: %w",
			ErrParse,
			distanceLine,
			err,
		)
	}

	result.Distance = distance

	return result, nil
}
