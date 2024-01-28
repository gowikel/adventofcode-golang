package day09

import (
	"fmt"
	"os"

	"github.com/gowikel/adventofcode-golang/internal/log"
)

type Exercise struct{}

func (e Exercise) Part1(data string) int {
	log := log.GetLogger(log.WithPart(1))

	lst, err := Parse(data)
	if err != nil {
		log.Fatal("Error while parsing the file", "err", err)
	}

	var result int
	for _, l := range lst {
		sr := NewSensorRead(l)
		err := sr.Compute()
		if err != nil {
			log.Fatal(
				fmt.Sprintf("Unable to compute the list %v", l),
				"err",
				err,
			)
		}
		result += sr.ExtrapolateForward()
	}

	return result
}

func (e Exercise) Part2(data string) int {
	log := log.GetLogger(log.WithPart(2))
	lst, err := Parse(data)
	if err != nil {
		log.Fatal("Error while parsing the file", "err", err)
	}

	var result int
	for _, l := range lst {
		sr := NewSensorRead(l)
		err := sr.Compute()
		if err != nil {
			log.Fatal(
				fmt.Sprintf("Unable to compute the list %v", l),
				"err",
				err,
			)
			os.Exit(1)
		}
		result += sr.ExtrapolateBackward()
	}

	return result
}
