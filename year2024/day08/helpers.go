package day08

import (
	"fmt"
	"os"

	"go.eryndalor.dev/adventofcode-golang/year2024/day08/parser"
)

func parseFile(path string) (result parser.GameInformation, err error) {
	file, err := os.Open(path)
	if err != nil {
		return result, fmt.Errorf("parseFile: %w", err)
	}
	defer func() {
		err = file.Close()
	}()
	return parser.Parse(file)
}
