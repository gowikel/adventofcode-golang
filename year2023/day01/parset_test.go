package day01_test

import (
	"testing"

	. "github.com/gowikel/adventofcode-golang/year2023/day01"
	"github.com/stretchr/testify/assert"
)

func TestParseInput_ShouldParseSingleNumberStringCorrectly(
	t *testing.T,
) {
	input := "123\n"
	expected := []int{13}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseInput_ShouldParseSingleNumberStringWithoutEndingOnNewLine(
	t *testing.T,
) {
	input := "123"
	expected := []int{13}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseInput_ShouldParseMultipleNumberStringsCorrectly(
	t *testing.T,
) {
	input := "123\n456\n789\n"
	expected := []int{13, 46, 79}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseInput_ShouldIgnoreEmptyLines(t *testing.T) {
	input := "123\n\n456\n\n789\n"
	expected := []int{13, 46, 79}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseInput_ShouldHandleInputWithLeadingTrailingSpaces(
	t *testing.T,
) {
	input := "  123  \n  456  \n  789  \n"
	expected := []int{13, 46, 79}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseInput_ShouldHandleInputWithLeadingTrailingNonNumericCharacters(
	t *testing.T,
) {
	input := "abc123def\nghi456jkl\nmno789pqr\n"
	expected := []int{13, 46, 79}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseInput_ShouldHandleInputWithOnlyOneNumericCharacter(
	t *testing.T,
) {
	input := "1\n2\n3\n"
	expected := []int{11, 22, 33}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseInput_ShouldIgnoreLinesWithNonNumericCharacters(
	t *testing.T,
) {
	input := "123\nabc\n456\n"
	expected := []int{13, 46}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseInput_ShouldParseInputWithOnlyTwoNumericCharactersCorrectly(
	t *testing.T,
) {
	input := "12\n"
	expected := []int{12}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseInput_ShouldHandleEmptyInput(t *testing.T) {
	input := ""
	expected := []int{}
	result, err := ParseInput(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

// Returns the first and last digit of a two-digit number.
func TestParseNumber_ShouldReturnFirstAndLastDigitOfTwoDigitNumber(
	t *testing.T,
) {
	input := "42"
	expected := 42
	result, err := ParseNumber(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

// Repeats a one-digit number twice and returns it.
func TestParseNumber_ShouldRepeatOneDigitNumberTwice(t *testing.T) {
	input := "5"
	expected := 55
	result, err := ParseNumber(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

// Returns the first and last digit of a number with more than two
// digits.
func TestParseNumber_ShouldReturnFirstAndLastDigitOfNumberWithMoreThanTwoDigits(
	t *testing.T,
) {
	input := "12345"
	expected := 15
	result, err := ParseNumber(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseNumber_ShouldReturnZeroIfInputIsEmptyString(
	t *testing.T,
) {
	input := ""
	expected := 0
	result, err := ParseNumber(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestParseNumber_ShouldHandleNegativeNumbersCorrectly(
	t *testing.T,
) {
	input := "-42"
	expected := -42
	result, err := ParseNumber(input)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestTokenizer_ReturnsData_WhenAtEOFIsFalseAndFirstByteOfDataIsANumber(
	t *testing.T,
) {
	data := []byte("1")
	atEOF := false
	advance, token, err := Tokenizer(data, atEOF)

	assert.Equalf(
		t,
		1,
		advance,
		"Expected advance to be 1, but got %d",
		advance,
	)
	assert.Equalf(
		t,
		token,
		data[:1],
		"Expected token to be %v, but got %v",
		data[:1],
		token,
	)
	assert.Nil(t, err)
}

func TestTokenizer_Returns1_WhenDataStartsWithOne(t *testing.T) {
	data := []byte("one")
	atEOF := false
	advance, token, err := Tokenizer(data, atEOF)

	assert.Equalf(
		t,
		1,
		advance,
		"Expected advance to be 1, but got %d",
		advance,
	)
	assert.Equalf(
		t,
		[]byte("1"),
		token,
		"Expected token to be [49], but got %v",
		token,
	)
	assert.Nil(t, err)
}

func TestTokenizer_Returns2_WhenDataStartsWithTwo(t *testing.T) {
	data := []byte("two")
	atEOF := false
	advance, token, err := Tokenizer(data, atEOF)

	assert.Equalf(
		t,
		1,
		advance,
		"Expected advance to be 1, but got %d",
		advance,
	)
	assert.Equalf(
		t,
		[]byte("2"),
		token,
		"Expected token to be [49], but got %v",
		token,
	)
	assert.Nil(t, err)
}

func TestTokenizer_Returns0_WhenAtEOFIsTrueAndDataIsEmpty(
	t *testing.T,
) {
	data := []byte{}
	atEOF := true
	advance, token, err := Tokenizer(data, atEOF)

	assert.Equalf(
		t,
		0,
		advance,
		"Expected advance to be 0, but got %d",
		advance,
	)
	assert.Nil(
		t,
		token,
		"Expected token to be nil, but got %v",
		token,
	)
	assert.Nil(t, err)
}

func TestTokenizer_ReturnsData_WhenAtEOFIsFalseAndFirstByteOfDataIsNotANumberOrWord(
	t *testing.T,
) {
	data := []byte("a")
	atEOF := false
	advance, token, err := Tokenizer(data, atEOF)

	assert.Equalf(
		t,
		1,
		advance,
		"Expected advance to be 1, but got %d",
		advance,
	)
	assert.Equalf(
		t,
		data[:1],
		token,
		"Expected token to be [49], but got %v",
		token,
	)
	assert.Nil(t, err)
}
