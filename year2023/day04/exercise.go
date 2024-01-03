package day04

import (
	"fmt"
	"log"
	"strings"
)

type Exercise struct{}

func (e Exercise) Solve(data string) {
	fmt.Printf("- Day 04\n")
	fmt.Printf("  Part 1: %d\n", part1(data))
	fmt.Printf("  Part 2: %d\n", part2(data))
}

func part1(data string) int {
	var result int

	for _, line := range strings.Split(data, "\n") {
		var lineResult int
		w, p, err := parseLine(line)

		if err != nil {
			log.Fatal(err)
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

func part2(data string) int {
	var result int
	lines := strings.Split(data, "\n")
	copies := make([]int, len(lines))

	for i := range copies {
		copies[i] = 1
	}

	for idx, line := range strings.Split(data, "\n") {
		w, p, err := parseLine(line)

		if err != nil {
			log.Fatal(err)
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
