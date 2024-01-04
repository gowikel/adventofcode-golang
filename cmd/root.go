package cmd

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gowikel/adventofcode-golang/cmd/internal/fs"
	"github.com/gowikel/adventofcode-golang/internal/conf"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var debug bool
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "aoc is a simple CLI tool to solve Advent of Code problems",
	Long: `A tool to generate boilerplate and solve the exercises of
	the Advent of Code (https://adventofcode.com/) in Golang
	`,
}

// Cobra entrypoint
func Execute() {
	now := time.Now()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
	d := time.Since(now)
	fmt.Printf("Executed in %s\n", d)
}

// Configuration
func init() {
	cobra.OnInitialize(initConfig)
	fs.EnsureConfigDirExists()

	rootCmd.PersistentFlags().
		BoolVar(&debug, "debug", false, "Enables the debug logger")

	viper.BindPFlag(
		"debug",
		rootCmd.PersistentFlags().Lookup("debug"),
	)

	// Defaults
	viper.SetDefault(conf.INPUT_DIR, "inputs")
}

func initConfig() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	baseConfDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to find the config dir")
	}
	confDir := path.Join(baseConfDir, "aoc")

	// Check configuration
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(confDir)
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("aoc")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Warn().Msg("No config file, using defaults")
		} else {
			log.Fatal().Err(err).Msg("Error reading config file")
		}
	}

	if viper.GetBool("debug") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
