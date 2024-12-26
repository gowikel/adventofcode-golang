package parser_test

import (
	"testing"

	. "github.com/gowikel/adventofcode-golang/year2024/day06/parser"
	"github.com/stretchr/testify/assert"
)

func TestParse_EmptyString(t *testing.T) {
	data := ""

	g, err := Parse(data)

	assert.NoError(t, err)
	assert.Equal(t, 0, g.Rows())
	assert.Equal(t, 0, g.Columns())
	assert.Equal(t, "", g.String())
}

func TestParse_SingleLine(t *testing.T) {
	data := "...#x>.."

	g, err := Parse(data)

	assert.NoError(t, err)
	assert.Equal(t, 1, g.Rows())
	assert.Equal(t, 8, g.Columns())
	assert.Equal(t, "...#x>..", g.String())
}

func TestParse_MultipleLines(t *testing.T) {
	data := "" +
		"...#\n" +
		"xxx>"

	g, err := Parse(data)

	assert.NoError(t, err)
	assert.Equal(t, 2, g.Rows())
	assert.Equal(t, 4, g.Columns())
	assert.Equal(t, "...#\nxxx>", g.String())
}

func TestParse_TrailingLines(t *testing.T) {
	data := "...\n"

	g, err := Parse(data)

	assert.NoError(t, err)
	assert.Equal(t, 1, g.Rows())
	assert.Equal(t, 3, g.Columns())
	assert.Equal(t, "...", g.String())
}

func TestParse_DetectsInvalidCells(t *testing.T) {
	data := "..................I.........."

	_, err := Parse(data)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid cell: I")
}
