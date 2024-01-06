package day03

import (
	_ "embed"
	"testing"
)

//go:embed testdata/2023_03_example.txt
var EXAMPLE string

var e Exercise

func TestDay03Part1(t *testing.T) {
	got := e.Part1(EXAMPLE)
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
	got := e.Part2(EXAMPLE)
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
