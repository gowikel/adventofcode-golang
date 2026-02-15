package parser_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.eryndalor.dev/adventofcode-golang/year2024/day08/parser"
)

func TestParse_EmptyInput(t *testing.T) {
	input := ""
	reader := strings.NewReader(input)
	result, err := parser.Parse(reader)

	assert.Nil(t, err)
	assert.Empty(t, result.RoofMap)
	assert.Equal(t, 0, result.Columns)
	assert.Equal(t, 0, result.Rows)
}

func TestParse_SingleLine_OnlyDots(t *testing.T) {
	input := "......."
	reader := strings.NewReader(input)
	result, err := parser.Parse(reader)

	assert.Nil(t, err)
	assert.Empty(t, result.RoofMap)
	assert.Equal(t, 7, result.Columns)
	assert.Equal(t, 1, result.Rows)
}

func TestParse_SingleLine_NewLineEnding(t *testing.T) {
	input := ".......\n"
	reader := strings.NewReader(input)
	result, err := parser.Parse(reader)

	assert.Nil(t, err)
	assert.Empty(t, result.RoofMap)
	assert.Equal(t, 7, result.Columns)
	assert.Equal(t, 1, result.Rows)
}

func TestParse_SingleLine_WithCharacters(t *testing.T) {
	input := "A..AA..\n"
	reader := strings.NewReader(input)
	result, err := parser.Parse(reader)

	assert.Nil(t, err)
	assert.Len(t, result.RoofMap, 1)
	assert.ElementsMatch(t, result.RoofMap['A'], []parser.Cell{
		{0, 0},
		{0, 3},
		{0, 4},
	})
	assert.Equal(t, 7, result.Columns)
	assert.Equal(t, 1, result.Rows)
}

func TestParse_SingleLine_WithMultipleCharacters(t *testing.T) {
	input := "AB..ABA\n"
	reader := strings.NewReader(input)
	result, err := parser.Parse(reader)

	assert.Nil(t, err)
	assert.Len(t, result.RoofMap, 2)
	assert.ElementsMatch(t, result.RoofMap['A'], []parser.Cell{
		{0, 0},
		{0, 4},
		{0, 6},
	})
	assert.ElementsMatch(t, result.RoofMap['B'], []parser.Cell{
		{0, 1},
		{0, 5},
	})
	assert.Equal(t, 7, result.Columns)
	assert.Equal(t, 1, result.Rows)
}

func TestParse_MultipleLines_EmptyLines(t *testing.T) {
	input := "\n\n\n"
	reader := strings.NewReader(input)
	result, err := parser.Parse(reader)

	assert.Nil(t, err)
	assert.Empty(t, result.RoofMap)
	assert.Equal(t, 0, result.Columns)
	assert.Equal(t, 0, result.Rows)
}

func TestParse_MultipleLines_WithCharacters(t *testing.T) {
	input := "A..AA..\nAB..ABA\n.BB.A.."
	reader := strings.NewReader(input)
	result, err := parser.Parse(reader)

	assert.Nil(t, err)
	assert.Len(t, result.RoofMap, 2)

	assert.ElementsMatch(t, result.RoofMap['A'], []parser.Cell{
		{0, 0},
		{0, 3},
		{0, 4},
		{1, 0},
		{1, 4},
		{1, 6},
		{2, 4},
	})
	assert.ElementsMatch(t, result.RoofMap['B'], []parser.Cell{
		{1, 1},
		{1, 5},
		{2, 1},
		{2, 2},
	})
	assert.Equal(t, 7, result.Columns)
	assert.Equal(t, 3, result.Rows)
}
