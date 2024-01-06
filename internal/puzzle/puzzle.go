package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Return the contents of the given puzzle
func Read(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("unable to open %q: %w", path, err)
	}
	defer file.Close()

	buffer := bufio.NewReader(file)
	bytes, err := io.ReadAll(buffer)
	if err != nil {
		return "", fmt.Errorf("unable to read %q: %w", path, err)
	}

	return string(bytes), nil
}
