package aoc

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
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

var REQUIRED_RED_CUBES = 12
var REQUIRED_BLUE_CUBES = 14
var REQUIRED_GREEN_CUBES = 13

//go:embed data/2023_02.txt
var DAY2_DATA string

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

func Day2Part1(data string) int {
	var result int

	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		gameLine := scanner.Text()
		gameId, cubes := ParseGame(gameLine)

		isValid := true
		for _, cubeSet := range cubes {
			if (cubeSet.Red > REQUIRED_RED_CUBES) ||
				(cubeSet.Blue > REQUIRED_BLUE_CUBES) ||
				(cubeSet.Green > REQUIRED_GREEN_CUBES) {
				isValid = false
				break
			}
		}

		if isValid {
			result += gameId
		}
	}

	return result
}

func PowerCube(cube SetCubes) int {
	return cube.Red * cube.Green * cube.Blue
}

func MinimumPowerSet(cubes []SetCubes) SetCubes {
	result := SetCubes{
		Red:   math.MinInt,
		Green: math.MinInt,
		Blue:  math.MinInt,
	}

	for _, cube := range cubes {
		if cube.Red > result.Red {
			result.Red = cube.Red
		}

		if cube.Green > result.Green {
			result.Green = cube.Green
		}

		if cube.Blue > result.Blue {
			result.Blue = cube.Blue
		}
	}

	return result
}

func Day2Part2(data string) int {
	var result int

	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		gameLine := scanner.Text()
		_, cubes := ParseGame(gameLine)

		result += PowerCube(MinimumPowerSet(cubes))
	}

	return result
}

func Day2() {
	fmt.Printf("- Day 02\n")
	fmt.Printf("  Part 1: %d\n", Day2Part1(DAY2_DATA))
	fmt.Printf("  Part 2: %d\n", Day2Part2(DAY2_DATA))
}
