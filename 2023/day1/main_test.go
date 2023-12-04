package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed data/2023_1_P1_example.txt
var example1 string

//go:embed data/2023_1_P2_example.txt
var example2 string

func TestParseInput(t *testing.T) {
	t.Run("Gets the correct output", func(t *testing.T) {
		got := ParseInput(example1)
		want := []int{12, 38, 15, 77}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ParseInput(example) = %#v; want %#v", got, want)
		}
	})

	t.Run("Handles empty input", func(t *testing.T) {
		got := ParseInput("")
		want := []int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ParseInput(\"\") = %#v; want %#v", got, want)
		}
	})

	t.Run("Handles empty lines", func(t *testing.T) {
		got := ParseInput("\n\n")
		want := []int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ParseInput(\"\\n\\n\") = %#v; want %#v", got, want)
		}
	})
}

func TestEnhancedParseInput(t *testing.T) {
	t.Run("Gets the correct output (example1)", func(t *testing.T) {
		got := EnhancedParseInput(example1)
		want := []int{12, 38, 15, 77}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(example1) = %#v; want %#v", got, want)
		}
	})

	t.Run("Gets the correct output (example2)", func(t *testing.T) {
		got := EnhancedParseInput(example2)
		want := []int{29, 83, 13, 24, 42, 14, 76}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(example2) = %#v; want %#v", got, want)
		}
	})

	t.Run("Handles one", func(t *testing.T) {
		input := "one"
		got := EnhancedParseInput(input)
		want := []int{11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles two", func(t *testing.T) {
		input := "two"
		got := EnhancedParseInput(input)
		want := []int{22}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles three", func(t *testing.T) {
		input := "three"
		got := EnhancedParseInput(input)
		want := []int{33}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles four", func(t *testing.T) {
		input := "four"
		got := EnhancedParseInput(input)
		want := []int{44}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles five", func(t *testing.T) {
		input := "five"
		got := EnhancedParseInput(input)
		want := []int{55}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles six", func(t *testing.T) {
		input := "six"
		got := EnhancedParseInput(input)
		want := []int{66}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles seven", func(t *testing.T) {
		input := "seven"
		got := EnhancedParseInput(input)
		want := []int{77}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles eight", func(t *testing.T) {
		input := "eight"
		got := EnhancedParseInput(input)
		want := []int{88}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles nine", func(t *testing.T) {
		input := "nine"
		got := EnhancedParseInput(input)
		want := []int{99}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles zero", func(t *testing.T) {
		input := "zero"
		got := EnhancedParseInput(input)
		want := []int{00}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles mixed inputs", func(t *testing.T) {
		input := "one3"
		got := EnhancedParseInput(input)
		want := []int{13}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Skips other characters", func(t *testing.T) {
		input := "one$3"
		got := EnhancedParseInput(input)
		want := []int{13}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles tricky inputs", func(t *testing.T) {
		input := "eightwo"
		got := EnhancedParseInput(input)
		want := []int{88}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
		}
	})

	t.Run("Handles tricky inputs (v2)", func(t *testing.T) {
		input := "threeight"
		got := EnhancedParseInput(input)
		want := []int{33}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("EnhancedParseInput(%#v) = %#v; want %#v", input, got, want)
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
