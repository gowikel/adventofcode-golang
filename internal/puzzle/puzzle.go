package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/gowikel/adventofcode-golang/internal/conf"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Return the contents of the given puzzle
func Read(year, day int, example bool) (string, error) {
	var filename string

	if example {
		filename = fmt.Sprintf("%04d_%02d_example.txt", year, day)
	} else {
		filename = fmt.Sprintf("%04d_%02d.txt", year, day)
	}
	inputDir := viper.GetString(conf.INPUT_DIR)
	yearDir := fmt.Sprintf("%04d", year)
	path := path.Join(inputDir, yearDir, filename)

	log.Debug().
		Str("module", "puzzle").
		Str("func", "Read").
		Str("path", path).
		Int("year", year).
		Int("day", day).
		Bool("example", example).
		Msg("Attempting to read puzzle")

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	text := string(bytes)

	log.Debug().
		Str("path", path).
		Int("year", year).
		Int("day", day).
		Bool("example", example).
		Str("module", "puzzle").
		Str("func", "Read").
		Msg("File succesfully read")

	return text, nil
}
