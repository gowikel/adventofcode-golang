package day02

import (
	"regexp"
	"strconv"
	"strings"
)

type SetCubes struct {
	Red   int
	Green int
	Blue  int
}

var GAME_REGEX = regexp.MustCompile(`Game (\d+): (.*)`)
var RED_REGEX = regexp.MustCompile(`(\d+) red`)
var BLUE_REGEX = regexp.MustCompile(`(\d+) blue`)
var GREEN_REGEX = regexp.MustCompile(`(\d+) green`)

func ParseGame(input string) (int, []SetCubes) {
	matches := GAME_REGEX.FindAllStringSubmatch(input, -1)

	if len(matches) == 0 {
		return 0, nil
	}

	gameId, _ := strconv.Atoi(matches[0][1])
	sets := strings.Split(matches[0][2], ";")
	cubes := make([]SetCubes, 0, len(sets))

	// 6 red, 1 blue, 3 green
	for _, setExpr := range sets {
		var cube SetCubes

		redMatch := RED_REGEX.FindAllStringSubmatch(setExpr, -1)
		blueMatch := BLUE_REGEX.FindAllStringSubmatch(setExpr, -1)
		greenMatch := GREEN_REGEX.FindAllStringSubmatch(setExpr, -1)

		if len(redMatch) > 0 {
			red, _ := strconv.Atoi(redMatch[0][1])
			cube.Red = red
		}

		if len(blueMatch) > 0 {
			blue, _ := strconv.Atoi(blueMatch[0][1])
			cube.Blue = blue
		}

		if len(greenMatch) > 0 {
			green, _ := strconv.Atoi(greenMatch[0][1])
			cube.Green = green
		}

		cubes = append(cubes, cube)
	}

	return gameId, cubes
}
