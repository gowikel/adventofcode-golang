package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/gowikel/adventofcode-golang/internal/cli"
	"github.com/gowikel/adventofcode-golang/internal/conf"
	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	"github.com/gowikel/adventofcode-golang/year2023"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	year := viper.GetInt(conf.CLI_YEAR)
	day := viper.GetInt(conf.CLI_DAY)
	runExample := viper.GetBool(conf.CLI_RUN_EXAMPLE)

	data, err := puzzle.Read(year, day, runExample)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to read puzzle input")
	}

	fmt.Println("Year:", year)
	fmt.Println("Day:", day)
	fmt.Println("Running example?", runExample)
	fmt.Println()

	// TODO: Will be updated to run other years in the future
	year2023.Run(day, data)
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	baseConfDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to find the config dir")
	}
	confDir := path.Join(baseConfDir, "aoc")

	// Create confDir, in case it does not exists
	err = os.Mkdir(confDir, 0640)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		log.Fatal().
			Err(err).
			Msgf("Unable to create %q folder", confDir)
	}

	// Check configuration
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(confDir)
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("aoc")
	viper.AutomaticEnv()

	// Defaults
	viper.SetDefault(conf.INPUT_DIR, "inputs")
	viper.SetDefault(conf.CLI_DEBUG, false)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Warn().Msg("No config file, using defaults")
		} else {
			log.Fatal().Err(err).Msg("Error reading config file")
		}
	}
	cli.ParseArgs()

	if viper.GetBool(conf.CLI_DEBUG) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
