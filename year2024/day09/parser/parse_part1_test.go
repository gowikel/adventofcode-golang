package parser_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.eryndalor.dev/adventofcode-golang/year2024/day09/parser"
)

func TestParsePart1_EmptyInput(t *testing.T) {
	data := ""
	input := strings.NewReader(data)
	expected := &parser.MemoryMap{FileInfo: map[int]int{}, MemoryMap: []int{}}
	got, err := parser.ParsePart1(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, got)
}

func TestParsePart1_OneElement(t *testing.T) {
	data := "5"
	input := strings.NewReader(data)
	expected := &parser.MemoryMap{
		FileInfo:  map[int]int{0: 5},
		MemoryMap: []int{0},
		TotalSize: 5,
	}
	got, err := parser.ParsePart1(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, got)
}

func TestParsePart1_TwoElements(t *testing.T) {
	data := "56"
	input := strings.NewReader(data)
	expected := &parser.MemoryMap{
		FileInfo:  map[int]int{0: 5},
		MemoryMap: []int{0, 6},
		TotalSize: 5,
	}

	got, err := parser.ParsePart1(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, got)
}

func TestParsePart1_MultipleElements(t *testing.T) {
	data := "567810"
	input := strings.NewReader(data)
	expected := &parser.MemoryMap{
		FileInfo:  map[int]int{0: 5, 1: 7, 2: 1},
		MemoryMap: []int{0, 6, 1, 8, 2, 0},
		TotalSize: 13,
	}
	got, err := parser.ParsePart1(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, got)
}

func TestParsePart1_InvalidInput(t *testing.T) {
	data := "5a"
	input := strings.NewReader(data)
	_, err := parser.ParsePart1(input)

	assert.Error(t, err)
	assert.EqualError(
		t, err, "parse: unable to convert to number: strconv.Atoi: parsing \"a\": invalid syntax",
	)
}
