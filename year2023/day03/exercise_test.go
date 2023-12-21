package day03_test

import (
	_ "embed"
	"testing"

	. "github.com/gowikel/adventofcode-golang/year2023/day03"
)

//go:embed testdata/2023_03_example.txt
var EXAMPLE string

func TestDay03Part1(t *testing.T) {
	got := Part1(EXAMPLE)
	want := 4361

	if got != want {
		t.Errorf(
			"\n\nPart 1\n\nInput:\n%v\nGot: %v\nWant: %v\n\n",
			EXAMPLE,
			got,
			want,
		)
	}
}

func TestDay03Part2(t *testing.T) {
	got := Part2(EXAMPLE)
	want := 467835

	if got != want {
		t.Errorf(
			"\n\nPart 1\n\nInput:\n%v\nGot: %v\nWant: %v\n\n",
			EXAMPLE,
			got,
			want,
		)
	}
}
