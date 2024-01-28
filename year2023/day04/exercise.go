package day04

import (
	"strings"

	"github.com/gowikel/adventofcode-golang/internal/log"
)

type Exercise struct{}

func (e Exercise) Part1(data string) int {
	var result int
	log := log.GetLogger(log.WithPart(1))

	for _, line := range strings.Split(data, "\n") {
		var lineResult int
		w, p, err := parseLine(line)

		if err != nil {
			log.Fatal("", "err", err)
		}

		for _, playedNumber := range p {
			if w.Contains(playedNumber) && lineResult == 0 {
				lineResult = 1
			} else if w.Contains(playedNumber) {
				lineResult *= 2
			}
		}

		result += lineResult
	}

	return result
}

func (e Exercise) Part2(data string) int {
	var result int
	log := log.GetLogger(log.WithPart(2))
	lines := strings.Split(data, "\n")
	copies := make([]int, len(lines))

	for i := range copies {
		copies[i] = 1
	}

	for idx, line := range strings.Split(data, "\n") {
		w, p, err := parseLine(line)

		if err != nil {
			log.Fatal("", "err", err)
		}

		matches, _ := countMatches(w, p)

		for i := idx + 1; i < idx+matches+1; i++ {
			copies[i] += copies[idx]
		}
	}

	for _, cpy := range copies {
		result += cpy
	}

	return result
}
