package day02

import (
	"bufio"
	"fmt"
	"os"
)

const REQUIRED_RED_CUBES = 12
const REQUIRED_BLUE_CUBES = 14
const REQUIRED_GREEN_CUBES = 13

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	var result int

	file, err := os.Open(path)
	if err != nil {
		return result, fmt.Errorf("Part1: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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

func (e Exercise) Part2(path string) (int, error) {
	var result int

	file, err := os.Open(path)
	if err != nil {
		return result, fmt.Errorf("Part2: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		gameLine := scanner.Text()
		_, cubes := ParseGame(gameLine)

		result += PowerCube(MinimumPowerSet(cubes))
	}

	return result, nil
}
