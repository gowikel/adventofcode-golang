package parser_test

import (
	"strings"
	"testing"

	"github.com/gowikel/adventofcode-golang/year2024/day09/parser"
	"github.com/stretchr/testify/assert"
)

func TestEmptyInput(t *testing.T) {
	data := ""
	input := strings.NewReader(data)
	expected := &parser.MemoryMap{FileInfo: map[int]int{}, MemoryMap: []int{}}
	got, err := parser.ParsePart1(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, got)
}

func TestOneElement(t *testing.T) {
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

func TestTwoElements(t *testing.T) {
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

func TestMultipleElements(t *testing.T) {
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

func TestInvalidInput(t *testing.T) {
	data := "5a"
	input := strings.NewReader(data)
	_, err := parser.ParsePart1(input)

	assert.Error(t, err)
	assert.EqualError(
		t, err, "parse: unable to convert to number: strconv.Atoi: parsing \"a\": invalid syntax",
	)
}
