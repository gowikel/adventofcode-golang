package day07

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Parses the given text as a list of hands
func Parse(
	data string,
	handDeterminer HandTypeDeterminer,
	cardStrengthDeterminer CardStrengthDeterminer,
) ([]Rank, error) {
	linesCount := strings.Count(data, "\n")
	result := make([]Rank, 0, linesCount)

	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return result, fmt.Errorf(
				"expected a hand and a bid in %q",
				line,
			)
		}

		hand, err := NewHand(
			fields[0],
			handDeterminer,
			cardStrengthDeterminer,
		)
		if err != nil {
			return result, errors.Join(
				fmt.Errorf("error while parsing %q", line),
				err,
			)
		}

		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			return result, errors.Join(
				fmt.Errorf("error while parsing %q", line),
				err,
			)
		}

		result = append(result, NewRank(hand, bid))
	}

	return result, nil
}
