package day06

import (
	"github.com/gowikel/adventofcode-golang/internal/log"
)

type Exercise struct{}

func (e Exercise) Part1(data string) int {
	log := log.GetLogger(log.WithPart(1))
	result := 1
	races, err := Parse(data)

	if err != nil {
		log.Fatal("", "err", err)
	}

	for _, race := range races {
		waysToWin := CountWaysToWin(race)
		result *= waysToWin
	}

	return result
}

func (e Exercise) Part2(data string) int {
	log := log.GetLogger(log.WithPart(2))

	result := 1
	race, err := ParsePart2(data)

	if err != nil {
		log.Fatal("", "err", err)
	}

	// Surprised it was so easy, compared
	// to the previous day
	waysToWin := CountWaysToWin(race)
	result *= waysToWin

	return result
}
