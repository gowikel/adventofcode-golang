package day02

import (
	"bufio"
	"fmt"
	"strings"
)

const REQUIRED_RED_CUBES = 12
const REQUIRED_BLUE_CUBES = 14
const REQUIRED_GREEN_CUBES = 13

type Exercise struct{}

func (e Exercise) Solve(data string) {
	fmt.Printf("- Day 02\n")
	fmt.Printf("  Part 1: %d\n", part1(data))
	fmt.Printf("  Part 2: %d\n", part2(data))
}

func part1(data string) int {
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

func part2(data string) int {
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
