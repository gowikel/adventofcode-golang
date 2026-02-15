package day01_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "go.eryndalor.dev/adventofcode-golang/year2024/day01"
)

func TestParseReturnsTwoLists(t *testing.T) {
	t.Skip("It needs review")
	data := "1 2\n3 4\n5 6\n7 8"
	r1, r2, err := Parse(data)

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 3, 5, 7}, r1)
	assert.Equal(t, []int{2, 4, 6, 8}, r2)
}

func TestParseHandlesEmptyLines(t *testing.T) {
	t.Skip("It needs review")
	data := "\n1 2\n\n3 4\n\n5 6\n\n7 8\n"
	r1, r2, err := Parse(data)

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 3, 5, 7}, r1)
	assert.Equal(t, []int{2, 4, 6, 8}, r2)
}

func TestParseHandlesMultipleSpaces(t *testing.T) {
	t.Skip("It needs review")
	data := "   1 2\n3      4\n5 6    \n  7  8  "
	r1, r2, err := Parse(data)

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 3, 5, 7}, r1)
	assert.Equal(t, []int{2, 4, 6, 8}, r2)
}

func TestParseErrorWithIncompleteLine(t *testing.T) {
	t.Skip("It needs review")
	data := "1 2\n3 4\n5 6\n7"
	_, _, err := Parse(data)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid line: \"7\"")
}

func TestParseErrorWithAlphaValues(t *testing.T) {
	t.Skip("It needs review")
	data := "a 2\n3 4\n5 6\n7 8"
	_, _, err := Parse(data)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid line: \"a 2\"")
}

func TestParseErrorWithTooManyValues(t *testing.T) {
	t.Skip("It needs review")
	data := "1 2\n3 4\n5 6\n7 8 9"
	_, _, err := Parse(data)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid line: \"7 8 9\"")
}
