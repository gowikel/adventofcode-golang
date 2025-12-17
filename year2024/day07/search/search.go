package search

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/gowikel/adventofcode-golang/year2024/day07/parser"
)

func Part1(eq parser.Equation) (bool, error) {
	generator := func(lst []int, value int) ([]int, error) {
		newValues := make([]int, 2*len(lst))

		for i, bufferedValue := range lst {
			newValues[2*i] = bufferedValue + value
			newValues[2*i+1] = bufferedValue * value
		}

		return newValues, nil
	}

	return searchAlgorithm(eq, generator)
}

func Part2(eq parser.Equation) (bool, error) {
	generator := func(lst []int, value int) ([]int, error) {
		newValues := make([]int, 3*len(lst))

		for i, bufferedValue := range lst {
			var sb strings.Builder
			sb.WriteString(strconv.Itoa(bufferedValue))
			sb.WriteString(strconv.Itoa(value))
			orValue, err := strconv.Atoi(sb.String())
			if err != nil {
				return nil, fmt.Errorf("unable to convert string to int: %w", err)
			}

			newValues[3*i] = bufferedValue + value
			newValues[3*i+1] = bufferedValue * value
			newValues[3*i+2] = orValue
		}

		return newValues, nil
	}

	return searchAlgorithm(eq, generator)
}

func searchAlgorithm(eq parser.Equation, generator func([]int, int) ([]int, error)) (bool, error) {
	if len(eq.Operands) == 0 && eq.Result != 0 {
		return false, nil
	}

	if len(eq.Operands) == 0 && eq.Result == 0 {
		return true, nil
	}

	if eq.Result == 0 && len(eq.Operands) > 1 {
		return false, nil
	}

	if eq.Result == 0 && len(eq.Operands) == 1 {
		return eq.Operands[0] == 0, nil
	}

	buffer := make([]int, 1)
	buffer[0] = eq.Operands[0]

	for _, operand := range eq.Operands[1:] {
		newValues, err := generator(buffer, operand)
		if err != nil {
			return false, fmt.Errorf("unable to generate new values: %w", err)
		}

		newValues = slices.DeleteFunc(newValues, func(value int) bool {
			return value > eq.Result
		})

		buffer = newValues
	}

	for _, value := range buffer {
		if value == eq.Result {
			return true, nil
		}
	}

	return false, nil
}
