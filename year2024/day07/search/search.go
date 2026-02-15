package search

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"go.eryndalor.dev/adventofcode-golang/year2024/day07/parser"
)

func Part1(eq parser.Equation) bool {
	generator := func(lst []int, value int) []int {
		newValues := make([]int, 2*len(lst))

		for i, bufferedValue := range lst {
			newValues[2*i] = bufferedValue + value
			newValues[2*i+1] = bufferedValue * value
		}

		return newValues
	}

	return searchAlgorithm(eq, generator)
}

func Part2(eq parser.Equation) bool {
	generator := func(lst []int, value int) []int {
		newValues := make([]int, 3*len(lst))

		for i, bufferedValue := range lst {
			var sb strings.Builder
			sb.WriteString(strconv.Itoa(bufferedValue))
			sb.WriteString(strconv.Itoa(value))
			orValue, err := strconv.Atoi(sb.String())
			if err != nil {
				panic(fmt.Sprintf("An unexpected error has occurred while concatenating two strings: %v ", err.Error()))
			}

			newValues[3*i] = bufferedValue + value
			newValues[3*i+1] = bufferedValue * value
			newValues[3*i+2] = orValue
		}

		return newValues
	}

	return searchAlgorithm(eq, generator)
}

func searchAlgorithm(eq parser.Equation, generator func([]int, int) []int) bool {
	if len(eq.Operands) == 0 {
		return eq.Result == 0
	}

	buffer := make([]int, 1)
	buffer[0] = eq.Operands[0]

	for _, operand := range eq.Operands[1:] {
		newValues := generator(buffer, operand)
		newValues = slices.DeleteFunc(newValues, func(value int) bool {
			return value > eq.Result
		})
		buffer = newValues
	}

	for _, value := range buffer {
		if value == eq.Result {
			return true
		}
	}

	return false
}
