package day02_test

import (
	"reflect"
	"testing"

	. "github.com/gowikel/adventofcode-golang/year2023/day02"
)

func TestParseGame(t *testing.T) {
	t.Run("Parses the game ID correctly", func(t *testing.T) {
		input := "Game 8: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
		got, _ := ParseGame(input)
		want := 8

		if got != want {
			t.Errorf(
				"ParseGame(%#v) got %#v but wants %#v",
				input,
				got,
				want,
			)
		}
	})

	t.Run("Parses the Cube Sets correctly", func(t *testing.T) {
		input := "Game 8: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
		_, got := ParseGame(input)
		want := []SetCubes{
			{Red: 6, Blue: 1, Green: 3},
			{Red: 1, Blue: 2, Green: 2},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf(
				"ParseGame(%#v) got %#v but wants %#v",
				input,
				got,
				want,
			)
		}
	})
}
