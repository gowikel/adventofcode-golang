package parser

import (
	"bufio"
	"fmt"
	"io"
)

type Cell [2]int
type RoofMap map[byte][]Cell

type GameInformation struct {
	RoofMap RoofMap
	Columns int
	Rows    int
}

func Parse(reader io.Reader) (GameInformation, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanBytes)
	result := GameInformation{}
	result.RoofMap = make(RoofMap)

	r := 0
	c := -1  // Start on -1 to account for initialization
	rwd := 0 // Rows With Data
	cwd := 0 // Columns With Data

	for scanner.Scan() {
		cells := scanner.Bytes()
		if len(cells) == 0 || len(cells) > 1 {
			return result, fmt.Errorf("invalid line: %q", scanner.Text())
		}
		cell := cells[0]

		if cell == '\n' {
			r++
			c = -1
			continue
		} else {
			c++
		}

		if (c + 1) > cwd {
			cwd = c + 1
		}
		if (r + 1) > rwd {
			rwd = r + 1
		}

		if cell == '.' {
			continue
		}

		if len(result.RoofMap[cell]) == 0 {
			result.RoofMap[cell] = make([]Cell, 0)
		}

		result.RoofMap[cell] = append(result.RoofMap[cell], Cell{r, c})
	}

	result.Rows = rwd
	result.Columns = cwd

	return result, scanner.Err()
}
