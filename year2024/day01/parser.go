package day01

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Parse(data string) ([]int, []int, error) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	n := strings.Count(data, "\n")
	l1 := make([]int, 0, n)
	l2 := make([]int, 0, n)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("invalid line: %q", line)
		}

		n1, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid line: %q", line)
		}

		n2, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid line: %q", line)
		}

		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	return l1, l2, nil
}
