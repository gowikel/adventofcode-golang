package day08

import (
	"fmt"
	"os"

	"go.eryndalor.dev/adventofcode-golang/year2024/day08/parser"
)

func parseFile(path string) (result parser.GameInformation, err error) {
	// #nosec G304 -- Path is controlled by the user running the solution, not external input
	file, err := os.Open(path)
	if err != nil {
		return result, fmt.Errorf("parseFile: %w", err)
	}
	defer func() {
		err = file.Close()
	}()
	return parser.Parse(file)
}
