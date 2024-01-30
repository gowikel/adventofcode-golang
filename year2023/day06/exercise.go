package day06

import (
	"fmt"
	"os"
)

type Exercise struct{}

func (e Exercise) Part1(data string) int {
	result := 1
	races, err := Parse(data)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, race := range races {
		waysToWin := CountWaysToWin(race)
		result *= waysToWin
	}

	return result
}

func (e Exercise) Part2(data string) int {

	result := 1
	race, err := ParsePart2(data)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Surprised it was so easy, compared
	// to the previous day
	waysToWin := CountWaysToWin(race)
	result *= waysToWin

	return result
}
