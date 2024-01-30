package day09

import (
	"fmt"
	"os"
)

type Exercise struct{}

func (e Exercise) Part1(data string) int {
	lst, err := Parse(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var result int
	for _, l := range lst {
		sr := NewSensorRead(l)
		err := sr.Compute()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		result += sr.ExtrapolateForward()
	}

	return result
}

func (e Exercise) Part2(data string) int {
	lst, err := Parse(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var result int
	for _, l := range lst {
		sr := NewSensorRead(l)
		err := sr.Compute()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		result += sr.ExtrapolateBackward()
	}

	return result
}
