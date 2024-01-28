package conf_test

import (
	"testing"

	. "github.com/gowikel/adventofcode-golang/internal/conf"
	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	"github.com/stretchr/testify/assert"
)

func TestParsePart_RunAll(t *testing.T) {
	input := 0
	want := puzzle.RunAll
	got, err := ParsePart(input)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestParsePart_RunPartOne(t *testing.T) {
	input := 1
	want := puzzle.RunPartOne
	got, err := ParsePart(input)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestParsePart_RunPartTwo(t *testing.T) {
	input := 2
	want := puzzle.RunPartTwo
	got, err := ParsePart(input)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestParsePart_InvalidNumber_Below(t *testing.T) {
	input := -1
	_, err := ParsePart(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid part: -1")
}

func TestParsePart_InvalidNumber_Upper(t *testing.T) {
	input := 3
	_, err := ParsePart(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid part: 3")
}
