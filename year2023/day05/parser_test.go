package day05_test

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/gowikel/adventofcode-golang/utils/algo"
	. "github.com/gowikel/adventofcode-golang/year2023/day05"
)

func parseSeedsLineError(
	t *testing.T,
	input string,
	got []int,
	want []int,
) {
	t.Errorf(
		"ParseSeedsLine(%q) got %v but wants %v\n",
		input,
		got,
		want,
	)
}

func TestParseSeedsLine_WithEmptyString(t *testing.T) {
	input := ""
	want := []int{}
	got, _ := ParseSeedsLine(input)

	if !reflect.DeepEqual(got, want) {
		parseSeedsLineError(t, input, got, want)
	}
}

func TestParseSeedsLine_WithEmptySeeds(t *testing.T) {
	input := "seeds:"
	want := []int{}
	got, _ := ParseSeedsLine(input)

	if !reflect.DeepEqual(got, want) {
		parseSeedsLineError(t, input, got, want)
	}
}

func TestParseSeedsLine_WithOneSeed(t *testing.T) {
	input := "seeds: 10"
	want := []int{10}
	got, _ := ParseSeedsLine(input)

	if !reflect.DeepEqual(got, want) {
		parseSeedsLineError(t, input, got, want)
	}
}

func TestParseSeedsLine_WithMultipleSeeds(t *testing.T) {
	input := "seeds: 10 20 30 40"
	want := []int{10, 20, 30, 40}
	got, _ := ParseSeedsLine(input)

	if !algo.UnorderedEqualSlices[int](got, want) {
		parseSeedsLineError(t, input, got, want)
	}
}

func TestParseSeedsLine_WithInvalidInput(t *testing.T) {
	input := "seeds: a b c d"
	_, got := ParseSeedsLine(input)

	if !errors.Is(got, strconv.ErrSyntax) {
		t.Errorf(
			"ParseSeedsLine didn't fail with an error syntax. Got %#v instead\n",
			got,
		)
	}
}

func TestParseSeedsLine_WithBigNumber(t *testing.T) {
	// Biggest int64 is 9223372036854775807
	input := "seeds: 9223372036854775808"
	_, got := ParseSeedsLine(input)

	if !errors.Is(got, strconv.ErrRange) {
		t.Errorf(
			"ParseSeedsLine didn't fail with an error range. Got %#v instead\n",
			got,
		)
	}
}
