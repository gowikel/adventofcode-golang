package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Equation struct {
	Result   int
	Operands []int
}

func Parse(input io.Reader) (result []Equation, err error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return result, fmt.Errorf("failed to scan line: %w", err)
		}

		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		eq, err := parseLine(line)
		if err != nil {
			return result, fmt.Errorf("failed to parse line: %w", err)
		}
		result = append(result, eq)
	}

	if err := scanner.Err(); err != nil {
		return result, fmt.Errorf("failed to finish scanning: %w", err)
	}

	return result, nil
}

func parseLine(line string) (Equation, error) {
	result := Equation{}

	fields := strings.Split(line, ":")
	if len(fields) != 2 {
		return result, fmt.Errorf("no separation between result and operands found: %q", line)
	}

	r := fields[0]
	operands := strings.Fields(fields[1])

	parsedResult, err := strconv.Atoi(r)
	if err != nil {
		return result, fmt.Errorf("failed to parse %q: %w", line, err)
	}
	result.Result = parsedResult

	result.Operands = make([]int, 0, len(operands))

	for _, operand := range operands {
		parsedOperand, err := strconv.Atoi(operand)
		if err != nil {
			return result, fmt.Errorf("failed to parse %q: %w", line, err)
		}
		result.Operands = append(result.Operands, parsedOperand)
	}

	return result, nil
}
