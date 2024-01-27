package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Enum to select which part to run
type PuzzleRunSelector int

const (
	RunAll PuzzleRunSelector = iota
	RunPartOne
	RunPartTwo
)

func (pps PuzzleRunSelector) String() string {
	if pps == RunAll {
		return "Run All"
	} else if pps == RunPartOne {
		return "Run Part One"
	}
	return "Run Part Two"
}

// Return the contents of the given puzzle
func Read(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Read: %w", err)
	}
	defer file.Close()

	buffer := bufio.NewReader(file)
	bytes, err := io.ReadAll(buffer)
	if err != nil {
		return "", fmt.Errorf("Read: %w", err)
	}

	return string(bytes), nil
}
