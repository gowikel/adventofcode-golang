package day09

import (
	"github.com/rs/zerolog/log"
)

type Exercise struct{}

func (e Exercise) Part1(data string) int {

	lst, err := Parse(data)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while parsing the file")
	}

	var result int
	for _, l := range lst {
		sr := NewSensorRead(l)
		err := sr.Compute()
		if err != nil {
			log.Fatal().
				Err(err).
				Msgf("Unable to compute the list %v", l)
		}
		result += sr.ExtrapolateForward()
	}

	return result
}

func (e Exercise) Part2(data string) int {
	lst, err := Parse(data)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while parsing the file")
	}

	var result int
	for _, l := range lst {
		sr := NewSensorRead(l)
		err := sr.Compute()
		if err != nil {
			log.Fatal().
				Err(err).
				Msgf("Unable to compute the list %v", l)
		}
		result += sr.ExtrapolateBackward()
	}

	return result
}
