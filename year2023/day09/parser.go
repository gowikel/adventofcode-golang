package day09

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Parse(data string) ([][]int, error) {
	var result [][]int

	lineBreaks := strings.Count(data, "\n")
	result = slices.Grow[[][]int](result, lineBreaks)

	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)

		pl := make([]int, 0, len(fields))

		for _, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				return result, fmt.Errorf("%q is not an int", field)
			}

			pl = append(pl, n)
		}

		result = append(result, pl)
	}

	return result, nil
}
