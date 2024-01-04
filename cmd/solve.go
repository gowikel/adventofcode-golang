package cmd

import (
	"errors"
	"fmt"

	"github.com/gowikel/adventofcode-golang/internal/cli"
	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	"github.com/gowikel/adventofcode-golang/year2023"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var solveCmd = &cobra.Command{
	Use:   "solve year day",
	Short: "Executes the solution for a given problem",
	Args: func(cmd *cobra.Command, args []string) error {
		// Basic args validation
		if len(args) != 2 {
			return errors.New(
				"called with invalid args: year and day are expected",
			)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		year, err := cli.ParseYear(args[0])
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}

		day, err := cli.ParseDay(args[1])
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
		runExample := viper.GetBool("example")

		fmt.Println("Year:", year)
		fmt.Println("Day:", day)
		fmt.Println("Run example?", runExample)
		fmt.Println()

		data, err := puzzle.Read(year, day, runExample)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to read puzzle input")
		}

		// TODO: Will be updated to run other years in the future
		year2023.Run(day, data)
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	solveCmd.Flags().BoolP("example", "e", false, "Use example input")
	viper.BindPFlag("example", solveCmd.Flags().Lookup("example"))
}
