package day09

import (
	"os"

	"github.com/gowikel/adventofcode-golang/year2024/day09/parser"
)

func parseFilePart1(path string) (result *parser.MemoryMap, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parser.Parse(file)
}
