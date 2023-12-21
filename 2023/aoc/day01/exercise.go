package day01

import "fmt"

func part1(data string) int {
	numbers := ParseInput(data)
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func part2(data string) int {
	numbers := EnhancedParseInput(data)
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func Solve() {
	fmt.Printf("- Day 01\n")
	fmt.Printf("  Part 1: %d\n", part1(DATA))
	fmt.Printf("  Part 2: %d\n", part2(DATA))
}
