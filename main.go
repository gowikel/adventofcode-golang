package main

import (
	"log/slog"
	"os"

	"github.com/gowikel/adventofcode-golang/internal/cli"
	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	"github.com/gowikel/adventofcode-golang/internal/utils"
	"github.com/gowikel/adventofcode-golang/year2023"
)

func main() {
	opts := cli.ParseFlags()

	logger := slog.New(
		slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{
			Level: opts.LogLevel,
		}),
	)

	slog.SetDefault(logger)

	data, err := puzzle.Read(opts.Input)
	if err != nil {
		slog.Error("Unable to read puzzle input", "error", err)
		os.Exit(1)
	}

	slog.Info(
		"Running exercise",
		"year",
		opts.Year,
		"day",
		opts.Day,
		"part",
		opts.Part,
	)

	utils.MeasureExecutionTime(func() {
		// TODO: Will be updated to run other years in the future
		year2023.Run(opts.Day, data, opts.Part)
	})()
}
