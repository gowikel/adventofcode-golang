package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var NOT_A_NUMBER = regexp.MustCompile(`\D`)

//go:embed data/2023_1.txt
var data string

func ParseInput(input string) []int {
	lines := strings.Split(input, "\n")
	result := make([]int, 0, len(lines))

	for _, line := range lines {
		parsedLine := NOT_A_NUMBER.ReplaceAllString(line, "")
		result = append(result, ParseNumber(parsedLine))
	}

	return result
}

func ParseNumber(input string) int {
	if len(input) == 1 {
		input = strings.Repeat(input, 2)
	} else if len(input) > 2 {
		input = string(input[0]) + string(input[len(input)-1])
	}

	parsedNumber, _ := strconv.Atoi(input)
	return parsedNumber
}

func part1() {
	numbers := ParseInput(data)
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func main() {
	part1()
}
