package day02_test

import (
	"testing"

	. "github.com/gowikel/adventofcode-golang/year2023/day02"
	"github.com/stretchr/testify/assert"
)

func TestPowerCubes(t *testing.T) {
	t.Run(
		"PowerCube(SetCubes{Red: 4, Green: 2, Blue: 6}) == 48",
		func(t *testing.T) {
			input := SetCubes{Red: 4, Green: 2, Blue: 6}
			got := PowerCube(input)
			want := 48

			assert.Equal(
				t,
				want,
				got,
				"PowerCube(%#v) got %#v but want %#v\n",
				input,
				got,
				want,
			)
		},
	)

	t.Run(
		"PowerCube(SetCubes{Red: 1, Green: 3, Blue: 4}) == 12",
		func(t *testing.T) {
			input := SetCubes{Red: 1, Green: 3, Blue: 4}
			got := PowerCube(input)
			want := 12

			assert.Equalf(
				t,
				want,
				got,
				"PowerCube(%#v) got %#v but want %#v\n",
				input,
				got,
				want,
			)
		},
	)

	t.Run(
		"PowerCube(SetCubes{Red: 20, Green: 13, Blue: 6}) == 1560",
		func(t *testing.T) {
			input := SetCubes{Red: 20, Green: 13, Blue: 6}
			got := PowerCube(input)
			want := 1560

			assert.Equalf(
				t,
				want,
				got,
				"PowerCube(%#v) got %#v but want %#v\n",
				input,
				got,
				want,
			)
		},
	)

	t.Run(
		"PowerCube(SetCubes{Red: 14, Green: 3, Blue: 15}) == 630",
		func(t *testing.T) {
			input := SetCubes{Red: 14, Green: 3, Blue: 15}
			got := PowerCube(input)
			want := 630

			assert.Equal(
				t,
				want,
				got,
				"PowerCube(%#v) got %#v but want %#v\n",
				input,
				got,
				want,
			)
		},
	)

	t.Run("PowerCube(SetCubes{Red: 6, Green: 3, Blue: 2}) == 36",
		func(t *testing.T) {
			input := SetCubes{Red: 6, Green: 3, Blue: 2}
			got := PowerCube(input)
			want := 36

			assert.Equalf(
				t,
				want,
				got,
				"PowerCube(%#v) got %#v but want %#v\n",
				input,
				got,
				want,
			)
		},
	)
}

func TestMinimumPowerSet(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		input := []SetCubes{
			{Red: 4, Blue: 3},
			{Red: 1, Green: 2, Blue: 6},
			{Green: 2},
		}
		got := MinimumPowerSet(input)
		want := SetCubes{Red: 4, Green: 2, Blue: 6}

		assert.Equalf(
			t,
			want,
			got,
			"Input: %#v\nGot: %#v\nWant: %#v\n",
			input,
			got,
			want,
		)
	})
	t.Run("Example 2", func(t *testing.T) {
		input := []SetCubes{
			{Blue: 1, Green: 2},
			{Green: 3, Blue: 4, Red: 1},
			{Green: 1, Blue: 1},
		}
		got := MinimumPowerSet(input)
		want := SetCubes{Red: 1, Green: 3, Blue: 4}

		assert.Equal(
			t,
			want,
			got,
			"Input: %#v\nGot: %#v\nWant: %#v\n",
			input,
			got,
			want,
		)
	})
	t.Run("Example 3", func(t *testing.T) {
		input := []SetCubes{
			{Green: 8, Blue: 6, Red: 20},
			{Blue: 5, Red: 4, Green: 13},
			{Green: 5, Red: 1},
		}
		got := MinimumPowerSet(input)
		want := SetCubes{Red: 20, Green: 13, Blue: 6}

		assert.Equalf(
			t,
			want,
			got,
			"Input: %#v\nGot: %#v\nWant: %#v\n",
			input,
			got,
			want,
		)
	})

	t.Run("Example 4", func(t *testing.T) {
		input := []SetCubes{
			{Green: 1, Red: 3, Blue: 6},
			{Green: 3, Red: 6},
			{Green: 3, Blue: 15, Red: 14},
		}
		got := MinimumPowerSet(input)
		want := SetCubes{Red: 14, Green: 3, Blue: 15}

		assert.Equalf(
			t,
			want,
			got,
			"Input: %#v\nGot: %#v\nWant: %#v\n",
			input,
			got,
			want,
		)
	})

	t.Run("Example 5", func(t *testing.T) {
		input := []SetCubes{
			{Red: 6, Blue: 1, Green: 3},
			{Blue: 2, Red: 1, Green: 2},
		}
		got := MinimumPowerSet(input)
		want := SetCubes{Red: 6, Green: 3, Blue: 2}

		assert.Equalf(
			t,
			want,
			got,
			"Input: %#v\nGot: %#v\nWant: %#v\n",
			input,
			got,
			want,
		)
	})
}
