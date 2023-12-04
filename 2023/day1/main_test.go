package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed data/2023_1_P1_example.txt
var example1 string

func TestParseInput(t *testing.T) {
	t.Run("Gets the correct output", func(t *testing.T) {
		got := ParseInput(example1)
		want := []int{12, 38, 15, 77}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ParseInput(example) = %#v; want %#v", got, want)
		}
	})
}

func TestParseNumber(t *testing.T) {
	t.Run("Returns the first and last digit", func(t *testing.T) {
		got := ParseNumber("123")
		want := 13
		if got != want {
			t.Errorf("ParseNumber(\"123\") = %#v; want %#v", got, want)
		}
	})

	t.Run("Returns the number on two digit numbers", func(t *testing.T) {
		got := ParseNumber("12")
		want := 12
		if got != want {
			t.Errorf("ParseNumber(\"12\") = %#v; want %#v", got, want)
		}
	})

	t.Run("Returns the number repeated on one digit numbers", func(t *testing.T) {
		got := ParseNumber("1")
		want := 11
		if got != want {
			t.Errorf("ParseNumber(\"1\") = %#v; want %#v", got, want)
		}
	})
}
