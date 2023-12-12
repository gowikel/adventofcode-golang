package aoc

import (
	"bytes"
	_ "embed"
	"reflect"
	"testing"
)

//go:embed data/2023_01_example1.txt
var DAY1_EXAMPLE1 string

//go:embed data/2023_01_example2.txt
var DAY1_EXAMPLE2 string

func TestParseInput_ShouldParseSingleNumberStringCorrectly(
	t *testing.T,
) {
	input := "123\n"
	expected := []int{13}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput_ShouldParseSingleNumberStringWithoutEndingOnNewLine(
	t *testing.T,
) {
	input := "123"
	expected := []int{13}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput_ShouldParseMultipleNumberStringsCorrectly(
	t *testing.T,
) {
	input := "123\n456\n789\n"
	expected := []int{13, 46, 79}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput_ShouldIgnoreEmptyLines(t *testing.T) {
	input := "123\n\n456\n\n789\n"
	expected := []int{13, 46, 79}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput_ShouldHandleInputWithLeadingTrailingSpaces(
	t *testing.T,
) {
	input := "  123  \n  456  \n  789  \n"
	expected := []int{13, 46, 79}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput_ShouldHandleInputWithLeadingTrailingNonNumericCharacters(
	t *testing.T,
) {
	input := "abc123def\nghi456jkl\nmno789pqr\n"
	expected := []int{13, 46, 79}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput_ShouldHandleInputWithOnlyOneNumericCharacter(
	t *testing.T,
) {
	input := "1\n2\n3\n"
	expected := []int{11, 22, 33}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput_ShouldIgnoreLinesWithNonNumericCharacters(
	t *testing.T,
) {
	input := "123\nabc\n456\n"
	expected := []int{13, 46}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput_ShouldParseInputWithOnlyTwoNumericCharactersCorrectly(
	t *testing.T,
) {
	input := "12\n"
	expected := []int{12}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput_ShouldHandleEmptyInput(t *testing.T) {
	input := ""
	expected := []int{}
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Returns the first and last digit of a two-digit number.
func TestParseNumber_ShouldReturnFirstAndLastDigitOfTwoDigitNumber(
	t *testing.T,
) {
	input := "42"
	expected := 42
	result := ParseNumber(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Repeats a one-digit number twice and returns it.
func TestParseNumber_ShouldRepeatOneDigitNumberTwice(t *testing.T) {
	input := "5"
	expected := 55
	result := ParseNumber(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Returns the first and last digit of a number with more than two
// digits.
func TestParseNumber_ShouldReturnFirstAndLastDigitOfNumberWithMoreThanTwoDigits(
	t *testing.T,
) {
	input := "12345"
	expected := 15
	result := ParseNumber(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Panics if the input is not a number.
func TestParseNumber_ShouldPanicIfInputIsNotANumber(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	input := "abc"
	ParseNumber(input)
}


func TestParseNumber_ShouldReturnZeroIfInputIsEmptyString(
	t *testing.T,
) {
	input := ""
	expected := 0
	result := ParseNumber(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestParseNumber_ShouldHandleNegativeNumbersCorrectly(
	t *testing.T,
) {
	input := "-42"
	expected := -42
	result := ParseNumber(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}


func TestTokenizer_ReturnsData_WhenAtEOFIsFalseAndFirstByteOfDataIsANumber(t *testing.T) {
	data := []byte("1")
	atEOF := false
	advance, token, err := tokenizer(data, atEOF)
	if advance != 1 {
			t.Errorf("Expected advance to be 1, but got %d", advance)
	}
	if !bytes.Equal(token, data[:1]) {
			t.Errorf("Expected token to be %v, but got %v", data[:1], token)
	}
	if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
	}
}

func TestTokenizer_Returns1_WhenDataStartsWithOne(t *testing.T) {
	data := []byte("one")
	atEOF := false
	advance, token, err := tokenizer(data, atEOF)
	if advance != 1 {
			t.Errorf("Expected advance to be 1, but got %d", advance)
	}
	if !bytes.Equal(token, []byte("1")) {
			t.Errorf("Expected token to be [49], but got %v", token)
	}
	if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
	}
}


func TestTokenizer_Returns2_WhenDataStartsWithTwo(t *testing.T) {
	data := []byte("two")
	atEOF := false
	advance, token, err := tokenizer(data, atEOF)
	if advance != 1 {
			t.Errorf("Expected advance to be 1, but got %d", advance)
	}
	if !bytes.Equal(token, []byte("2")) {
			t.Errorf("Expected token to be [50], but got %v", token)
	}
	if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
	}
}

func TestTokenizer_Returns0_WhenAtEOFIsTrueAndDataIsEmpty(t *testing.T) {
	data := []byte{}
	atEOF := true
	advance, token, err := tokenizer(data, atEOF)
	if advance != 0 {
			t.Errorf("Expected advance to be 0, but got %d", advance)
	}
	if token != nil {
			t.Errorf("Expected token to be nil, but got %v", token)
	}
	if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
	}
}

func TestTokenizer_ReturnsData_WhenAtEOFIsFalseAndFirstByteOfDataIsNotANumberOrWord(t *testing.T) {
	data := []byte("a")
	atEOF := false
	advance, token, err := tokenizer(data, atEOF)
	if advance != 1 {
			t.Errorf("Expected advance to be 1, but got %d", advance)
	}
	if !bytes.Equal(token, data[:1]) {
			t.Errorf("Expected token to be %v, but got %v", data[:1], token)
	}
	if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
	}
}