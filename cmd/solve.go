package cmd

import (
	"errors"
	"fmt"

	"github.com/gowikel/adventofcode-golang/internal/cli"
	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	"github.com/gowikel/adventofcode-golang/internal/puzzlePartSelector"
	"github.com/gowikel/adventofcode-golang/year2023"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var solveCmd = &cobra.Command{
	Use:   "solve [year] day",
	Short: "Executes the solution for a given problem",
	Args: func(cmd *cobra.Command, args []string) error {
		// Basic args validation
		if len(args) == 0 || len(args) > 2 {
			return errors.New(
				"called with invalid args: year and/or day are expected",
			)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		year, day, err := cli.ParseYearAndDay(args)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}

		puzzleFile, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to read path")
		}

		part, err := cmd.Flags().GetInt("part")
		parsedPart := puzzlePartSelector.RunAll

		if err != nil {
			log.Fatal().Err(err).Msg("Unable to read part")
		} else if part == 1 {
			parsedPart = puzzlePartSelector.RunPartOne
		} else if part == 2 {
			parsedPart = puzzlePartSelector.RunPartTwo
		}

		fmt.Println("Year:", year)
		fmt.Println("Day:", day)
		fmt.Println(parsedPart)
		fmt.Println()

		data, err := puzzle.Read(puzzleFile)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to read puzzle input")
		}

		// TODO: Will be updated to run other years in the future
		year2023.Run(day, data, parsedPart)
		fmt.Println()
	},
}

func init() {
	solveCmd.Flags().
		StringP("file", "f", "", "Path to the file to be used in the solution (required)")
	solveCmd.Flags().
		IntP("part", "p", 0, "Part to run. Omit to run both.")

	solveCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(solveCmd)
}
