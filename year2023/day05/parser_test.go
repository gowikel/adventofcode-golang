package day05_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	. "go.eryndalor.dev/adventofcode-golang/year2023/day05"
)

func TestParseSeedsLine_WithEmptyString(t *testing.T) {
	input := ""
	want := []int{}
	got, _ := ParseSeedsLine(input)

	assert.Equal(t, want, got)
}

func TestParseSeedsLine_WithEmptySeeds(t *testing.T) {
	input := "seeds:"
	want := []int{}
	got, _ := ParseSeedsLine(input)

	assert.Equal(t, want, got)
}

func TestParseSeedsLine_WithOneSeed(t *testing.T) {
	input := "seeds: 10"
	want := []int{10}
	got, _ := ParseSeedsLine(input)

	assert.Equal(t, want, got)
}

func TestParseSeedsLine_WithMultipleSeeds(t *testing.T) {
	input := "seeds: 10 20 30 40"
	want := []int{10, 20, 30, 40}
	got, _ := ParseSeedsLine(input)

	assert.Equal(t, want, got)
}

func TestParseSeedsLine_WithInvalidInput(t *testing.T) {
	input := "seeds: a b c d"
	_, got := ParseSeedsLine(input)

	assert.Error(t, got)
	assert.ErrorIs(t, got, strconv.ErrSyntax)
}

func TestParseSeedsLine_WithBigNumber(t *testing.T) {
	// Biggest int64 is 9223372036854775807
	input := "seeds: 9223372036854775808"
	_, got := ParseSeedsLine(input)

	assert.Error(t, got)
	assert.ErrorIs(t, got, strconv.ErrRange)
}
