package parser_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.eryndalor.dev/adventofcode-golang/year2024/day07/parser"
)

func TestParse__EmptyFile(t *testing.T) {
	data := ""
	input := strings.NewReader(data)

	result, err := parser.Parse(input)

	assert.Nil(t, err)
	assert.Empty(t, result)
}

func TestParse__EmptyLinesAreIgnored(t *testing.T) {
	data := "      \n\n\n\t   "
	input := strings.NewReader(data)

	result, err := parser.Parse(input)

	assert.Nil(t, err)
	assert.Empty(t, result)
}

func TestParse__SingleLine(t *testing.T) {
	data := "123: 4 5 6"
	input := strings.NewReader(data)
	expected := []parser.Equation{
		{
			Result:   123,
			Operands: []int{4, 5, 6},
		},
	}

	result, err := parser.Parse(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestParse__MultipleLines(t *testing.T) {
	data := "123: 4 5 6\n456: 7 8 9"
	input := strings.NewReader(data)
	expected := []parser.Equation{
		{
			Result:   123,
			Operands: []int{4, 5, 6},
		},
		{
			Result:   456,
			Operands: []int{7, 8, 9},
		},
	}

	result, err := parser.Parse(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestParse__MalformedLine(t *testing.T) {
	data := "123: 4 notanumber 5"
	input := strings.NewReader(data)

	_, err := parser.Parse(input)

	assert.ErrorContains(t, err, "failed to parse")
}

func TestParse__NoResult(t *testing.T) {
	data := "123 4 5 6"
	input := strings.NewReader(data)

	_, err := parser.Parse(input)

	assert.ErrorContains(t, err, "failed to parse")
}
