package fs

import (
	"errors"
	"io/fs"
	"os"
	"path"

	"github.com/rs/zerolog/log"
)

func EnsureConfigDirExists() {
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
}
