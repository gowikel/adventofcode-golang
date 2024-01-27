package cli_test

import (
	"testing"

	"github.com/gowikel/adventofcode-golang/internal/cli"
	"github.com/gowikel/adventofcode-golang/internal/puzzlePartSelector"
	"github.com/stretchr/testify/assert"
)

func TestParseYear_Number(t *testing.T) {
	input := "2023"
	want := 2023
	got, err := cli.ParseYear(input)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestParseYear_NaN(t *testing.T) {
	input := "abc"
	_, err := cli.ParseYear(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "ParseYear:")
}

func TestParseYear_OutOfRange(t *testing.T) {
	input := "2024"
	_, err := cli.ParseYear(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "ParseYear:")
	assert.ErrorContains(
		t,
		err,
		"\"2024\" is not in the range [2023-2023]",
	)
}

func TestParseDay_Number(t *testing.T) {
	input := "23"
	want := 23
	got, err := cli.ParseDay(input)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestParseDay_NaN(t *testing.T) {
	input := "abc"
	_, err := cli.ParseDay(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "ParseDay:")
}

func TestParseDay_OutOfRange_Below(t *testing.T) {
	input := "-1"
	_, err := cli.ParseDay(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "ParseDay:")
	assert.ErrorContains(t, err, "\"-1\" is not in the range [1-25]")
}

func TestParseDay_OutOfRange_Above(t *testing.T) {
	input := "26"
	_, err := cli.ParseDay(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "ParseDay:")
	assert.ErrorContains(t, err, "\"26\" is not in the range [1-25]")
}

func TestParsePart_RunAll(t *testing.T) {
	input := 0
	want := puzzlePartSelector.RunAll
	got, err := cli.ParsePart(input)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestParsePart_RunPartOne(t *testing.T) {
	input := 1
	want := puzzlePartSelector.RunPartOne
	got, err := cli.ParsePart(input)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestParsePart_RunPartTwo(t *testing.T) {
	input := 2
	want := puzzlePartSelector.RunPartTwo
	got, err := cli.ParsePart(input)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestParsePart_InvalidNumber_Below(t *testing.T) {
	input := -1
	_, err := cli.ParsePart(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid part: -1")
}

func TestParsePart_InvalidNumber_Upper(t *testing.T) {
	input := 3
	_, err := cli.ParsePart(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid part: 3")
}
