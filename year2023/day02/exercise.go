package day02

import (
	"bufio"
	"strings"
)

const REQUIRED_RED_CUBES = 12
const REQUIRED_BLUE_CUBES = 14
const REQUIRED_GREEN_CUBES = 13

type Exercise struct{}

func (e Exercise) Part1(data string) (int, error) {
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

	return result, nil
}

func (e Exercise) Part2(data string) (int, error) {
	var result int

	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		gameLine := scanner.Text()
		_, cubes := ParseGame(gameLine)

		result += PowerCube(MinimumPowerSet(cubes))
	}

	return result, nil
}
