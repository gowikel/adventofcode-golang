package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PageData struct {
	BeforeRules map[int][]int
	Pages       [][]int
}

func Parse(path string) (*PageData, error) {
	before_rules := make(map[int][]int)
	pages := make([][]int, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() //nolint:errcheck

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanning_rules := true

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			scanning_rules = false
			continue
		}

		if scanning_rules {
			parts := strings.Split(line, "|")

			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid line: %q", line)
			}

			before, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("%q is invalid: %w", parts[0], err)
			}

			after, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("%q is invalid: %w", parts[1], err)
			}

			_, ok := before_rules[before]
			if !ok {
				before_rules[before] = make([]int, 0, 1)
			}

			before_rules[before] = append(before_rules[before], after)
		} else {
			data := strings.Split(line, ",")
			numbers := make([]int, 0, len(data))

			for _, n := range data {
				parsed, err := strconv.Atoi(n)
				if err != nil {
					return nil, fmt.Errorf("invalid number: %q: %w", n, err)
				}
				numbers = append(numbers, parsed)
			}

			pages = append(pages, numbers)
		}
	}

	return &PageData{
		BeforeRules: before_rules,
		Pages:       pages,
	}, nil
}
