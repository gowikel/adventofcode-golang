package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/gowikel/adventofcode-golang/cmd/internal/fs"
	"github.com/gowikel/adventofcode-golang/internal/conf"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var now time.Time

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "aoc is a simple CLI tool to solve Advent of Code problems",
	Long: `A tool to generate boilerplate and solve the exercises of
	the Advent of Code (https://adventofcode.com/) in Golang
	`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		now = time.Now()

		// If something goes wrong, a SIGINT could be raised
		// and it will be intercepted here. Thus, this will
		// end the program, and print the execution time.
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT)
		// The magic happens here, in a separated goroutine
		go func() {
			<-sigs

			d := time.Since(now)
			fmt.Printf("Executed in %s\n", d)
			os.Exit(0)
		}()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		d := time.Since(now)
		fmt.Printf("Executed in %s\n", d)
	},
}

// Cobra entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

// Configuration
func init() {
	cobra.OnInitialize(initConfig)
	fs.EnsureConfigDirExists()

	rootCmd.PersistentFlags().
		Bool("debug", false, "Enables the debug logger")

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
