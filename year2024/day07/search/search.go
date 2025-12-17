package search

import (
	"slices"

	"github.com/gowikel/adventofcode-golang/year2024/day07/parser"
)

func SearchPart1(eq parser.Equation) bool {
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

func searchAlgorithm(eq parser.Equation, generator func([]int, int) []int) bool {
	if len(eq.Operands) == 0 && eq.Result != 0 {
		return false
	}

	if len(eq.Operands) == 0 && eq.Result == 0 {
		return true
	}

	if eq.Result == 0 && len(eq.Operands) > 1 {
		return false
	}

	if eq.Result == 0 && len(eq.Operands) == 1 {
		return eq.Operands[0] == 0
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
