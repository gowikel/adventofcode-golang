package day05

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"go.eryndalor.dev/adventofcode-golang/year2023/day05/intmapper"
)

// ErrParse is returned when an error has happened during parsing
var ErrParse = errors.New("error while parsing")

// ErrNotEnoughData is returned when the input given does not
// have enough data to complete.
var ErrNotEnoughData = errors.New("no enought data to complete")

// ErrInvalidAlmanacHeader is returned when the almanac header is not
// correctly parsed
var ErrInvalidAlmanacHeader = errors.New("invalid almanac header")

// ErrInvalidAlmanacRange is returned when the almanac range is
// not valid
var ErrInvalidAlmanacRange = errors.New("invalid almanac range")

// ParseSeedsLine parses a line of seeds and returns a slice of
// integers.
// It expects the line to start with the prefix "seeds:".
// If the line does not start with the prefix or is empty, it returns
// an empty slice. Each field in the line is parsed as an integer and
// added to the result slice. If any field fails to parse as an
// integer, it returns an error with the field and the parsing error.
func ParseSeedsLine(line string) ([]int, error) {
	result := []int{}

	if !strings.HasPrefix(line, "seeds:") {
		return result, nil
	}

	line = strings.TrimSpace(line[6:])

	if len(line) == 0 {
		return result, nil
	}

	for _, field := range strings.Fields(line) {
		parsedInt, err := strconv.Atoi(field)

		if err != nil {
			return result, errors.Join(ErrParse, err)
		}

		result = append(result, parsedInt)
	}

	return result, nil
}

// SeedRange contains information about a range
// of seeds. It can be used to create iterators
// over those ranges, without allocatiing the list
// of integers first (python range inspired).
type SeedRange struct {
	Start int
	End   int
	Len   int
}

// Returns a new initialized SeedRange
func NewSeedRange(start, length int) SeedRange {
	return SeedRange{
		Start: start,
		Len:   length,
		End:   start + length - 1,
	}
}

// ParseSeedLineAsRanges will parse the seeds line, but
// interpreting the line as a sequence of tuples, where
// the first element of the tuple specifies the start
// of the seed range, and the second the length of the
// range.
//
// E.g. 10 5, will be seeds 10, 11, 12, 13, 14 and 15
func ParseSeedLineAsRanges(line string) ([]SeedRange, error) {
	result := []SeedRange{}

	if !strings.HasPrefix(line, "seeds:") {
		return result, nil
	}

	line = strings.TrimSpace(line[6:])
	if len(line) == 0 {
		return result, nil
	}

	fields := strings.Fields(line)
	for i := 0; i < len(fields)/2; i++ {
		startIndex := i * 2
		lenIndex := startIndex + 1

		lineToParse := fmt.Sprintf(
			"%s %s",
			fields[startIndex],
			fields[lenIndex],
		)

		var start, length int
		_, err := fmt.Sscanf(lineToParse, "%d %d", &start, &length)
		if err != nil {
			return result, errors.Join(ErrParse, err)
		}

		result = append(result, NewSeedRange(start, length))
	}

	return result, nil
}

// AlamanacEntry represents an entry on the almanac
type AlmanacEntry struct {
	From   string
	To     string
	Mapper intmapper.IntMapper
}

// NewAlamacEntry will return an almanac with an empty list
// of ranges
func NewAlamacEntry(
	from string,
	to string,
) AlmanacEntry {
	return AlmanacEntry{
		From:   from,
		To:     to,
		Mapper: intmapper.New(),
	}
}

// ParseAlmanacEntry gets an almanac entry, which is composed
// of a title and a list of ranges, and returns the right
// AlmanacEntry structure, with the Mapper initialized
func ParseAlmanacEntry(data string) (AlmanacEntry, error) {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	if len(lines) < 2 {
		return AlmanacEntry{}, errors.Join(ErrParse, ErrNotEnoughData)
	}

	header := lines[0]
	header = strings.TrimSuffix(header, " map:")
	headerData := strings.Split(header, "-to-")

	if len(headerData) != 2 {
		return AlmanacEntry{}, errors.Join(
			ErrParse,
			ErrInvalidAlmanacHeader,
		)
	}

	from := headerData[0]
	to := headerData[1]

	result := NewAlamacEntry(from, to)

	for _, line := range lines[1:] {
		var source, destination, length int

		n, err := fmt.Sscanf(
			strings.TrimSpace(line),
			"%d %d %d",
			&destination,
			&source,
			&length,
		)

		if n != 3 || err != nil {
			return result, errors.Join(
				ErrParse,
				ErrInvalidAlmanacRange,
				err,
			)
		}

		result.Mapper.AddMapping(source, destination, length)
	}

	return result, nil
}

// ParseAlmanacLines gets all the almanac lines and converts
// them into a map with the From acting as key, and a AlmanacEntry
// object as value. The AlmanacEntry contains information about
// the ranges in that entry, and the From and To values.
//
// If an error occurs while parsing, it stops and returns the error.
func ParseAlmanacLines(data string) (map[string]AlmanacEntry, error) {
	almanacMaps := make(map[string]AlmanacEntry)
	entities := strings.Split(data, "\n\n")

	for _, entity := range entities {
		almanac, err := ParseAlmanacEntry(entity)

		if err != nil {
			return almanacMaps, errors.Join(ErrParse, err)
		}

		almanacMaps[almanac.From] = almanac
	}

	return almanacMaps, nil
}

// Parse will get the whole input, and return both,
// the list of seeds and a map between an entity
// and its destination
// func Parse(data string) ([]int, map[string]AlmanacEntry, error) {
// 	entities := strings.Split(data, "\n\n")
// 	almanacMaps := make(map[string]AlmanacEntry)

// 	if len(entities) < 2 {
// 		return []int{}, almanacMaps, errors.Join(ErrParse)
// 	}

// 	seedsLine, err := ParseSeedsLine(entities[0])
// 	if err != nil {
// 		return seedsLine, almanacMaps, errors.Join(ErrParse, err)
// 	}

// 	for _, entity := range entities[1:] {
// 		almanac, err := ParseAlmanacEntry(entity)

// 		if err != nil {
// 			return seedsLine, almanacMaps, errors.Join(ErrParse, err)
// 		}

// 		almanacMaps[almanac.From] = almanac
// 	}

// 	return seedsLine, almanacMaps, nil
// }
