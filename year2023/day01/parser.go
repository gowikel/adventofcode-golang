package day01

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

var NUMBER = regexp.MustCompile(`\d`)

var STARTS_WITH_WORD_NUMBER = regexp.MustCompile(
	`^(one|two|three|four|five|six|seven|eight|nine|zero)`,
)
var NOT_A_NUMBER = regexp.MustCompile(`\D`)

// ParseInput takes a string and returns a slice of ints, ignoring
// any non-numeric characters.
//
// A number will be added to the slice per line. The number
// will be the first and last digit of the line.
//
// If the line has only one digit, it will be repeated twice.
func ParseInput(input string) []int {
	lines := strings.Split(input, "\n")
	result := make([]int, 0, len(lines))

	for _, line := range lines {
		parsedLine := NOT_A_NUMBER.ReplaceAllString(line, "")

		if len(parsedLine) == 0 {
			continue
		}

		result = append(result, ParseNumber(parsedLine))
	}

	return result
}

// tokenizer is a split function for a Scanner that will recognize
// numbers and words that represent numbers.
//
// It returns each token as a byte slice and converts words to numbers
// before returning them.
//
// If the token is not a number or a word, it will return the token as
// it is.
//
// This function does not return an error.
func Tokenizer(
	data []byte,
	atEOF bool,
) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	} else if NUMBER.Match(data[:1]) {
		return 1, data[:1], nil
	} else if STARTS_WITH_WORD_NUMBER.Match(data) {
		if bytes.HasPrefix(data, []byte("one")) {
			return 1, []byte("1"), nil
		} else if bytes.HasPrefix(data, []byte("two")) {
			return 1, []byte("2"), nil
		} else if bytes.HasPrefix(data, []byte("three")) {
			return 1, []byte("3"), nil
		} else if bytes.HasPrefix(data, []byte("four")) {
			return 1, []byte("4"), nil
		} else if bytes.HasPrefix(data, []byte("five")) {
			return 1, []byte("5"), nil
		} else if bytes.HasPrefix(data, []byte("six")) {
			return 1, []byte("6"), nil
		} else if bytes.HasPrefix(data, []byte("seven")) {
			return 1, []byte("7"), nil
		} else if bytes.HasPrefix(data, []byte("eight")) {
			return 1, []byte("8"), nil
		} else if bytes.HasPrefix(data, []byte("nine")) {
			return 1, []byte("9"), nil
		} else if bytes.HasPrefix(data, []byte("zero")) {
			return 1, []byte("0"), nil
		}
	}

	return 1, data[:1], nil
}

// EnhancedParseInput takes into account that the input may contain
// words that represent numbers into the parsing.
//
// It returns a slice of ints.
func EnhancedParseInput(input string) []int {
	var sb strings.Builder
	var scanner = bufio.NewScanner(strings.NewReader(input))
	scanner.Split(Tokenizer)

	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}

	return ParseInput(sb.String())
}

// ParseNumber takes a string, which should be a number, and returns
// the first and last digit of the number.
//
// If the number is only one digit, it will be repeated twice.
//
// This function panics if the input is not a number.
func ParseNumber(input string) int {
	if len(input) == 0 {
		return 0
	} else if len(input) == 1 {
		input = strings.Repeat(input, 2)
	} else {
		// handle negative numbers
		if input[0] == '-' && len(input) > 3 {
			input = input[0:2] + string(input[len(input)-1])
		} else if input[0] != '-' {
			input = string(input[0]) + string(input[len(input)-1])
		}
	}

	parsedNumber, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal().
			Str("Function", "ParseNumber").
			Err(err).
			Msgf("Error while parsing %q", input)
	}

	return parsedNumber
}
