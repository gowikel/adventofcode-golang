package main

import (
	"embed"
	"fmt"

	"github.com/gowikel/adventofcode-golang/internal/cli"
	"github.com/rs/zerolog/log"
)

//go:embed inputs/*
var inputFiles embed.FS

func fetchData(o cli.AOCOptions) string {
	example_chunk := ""

	if o.RunExample && o.Day == 1 && o.Year == 2023 {
		example_chunk = "_example2"
	} else if o.RunExample {
		example_chunk = "_example"
	}

	path := fmt.Sprintf(
		"inputs/%04d/%04d_%02d%s.txt",
		o.Year,
		o.Year,
		o.Day,
		example_chunk,
	)
	data, err := inputFiles.ReadFile(path)

	if err != nil {
		log.Fatal().
			Str("path", path).
			Msg("Unable to read the puzzle data")
	}

	return string(data)
}
