package day06

import (
	"fmt"
	"log"
)

type Exercise struct{}

func (e Exercise) Solve(data string) {
	fmt.Printf("- Day 06\n")
	fmt.Printf("  Part 1: %d\n", part1(data))
	fmt.Printf("  Part 2: %d\n", part2(data))
}

func part1(data string) int {
	result := 1
	races, err := Parse(data)

	if err != nil {
		log.Fatal(err)
	}

	for _, race := range races {
		waysToWin := CountWaysToWin(race)
		result *= waysToWin
	}

	return result
}

func part2(data string) int {
	result := 1
	race, err := ParsePart2(data)

	if err != nil {
		log.Fatal(err)
	}

	// Surprised it was so easy, compared
	// to the previous day
	waysToWin := CountWaysToWin(race)
	result *= waysToWin

	return result
}
