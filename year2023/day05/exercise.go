package day05

import (
	"math"
	"slices"
	"strings"

	"github.com/gowikel/adventofcode-golang/internal/log"
	"github.com/gowikel/adventofcode-golang/internal/utils"
)

type Exercise struct{}

func ApplyFuncs(
	almanacEntries map[string]AlmanacEntry,
	seed int,
) int {

	currentEntry, moreValues := almanacEntries["seed"]
	currentValue := seed

	for currentEntry.From != "location" && moreValues {
		currentValue = currentEntry.Mapper.Transform(currentValue)
		currentEntry, moreValues = almanacEntries[currentEntry.To]
	}

	return currentValue
}

func (e Exercise) Part1(data string) int {
	result := math.MaxInt

	// Prepare the fatal logger, in case an error occurs
	log := log.GetLogger(log.WithPart(1))

	seedsPart := strings.Index(data, "\n\n")
	seeds, err := ParseSeedsLine(data[:seedsPart])
	if err != nil {
		log.Fatal("", "err", err)
	}

	almanacEntries, err := ParseAlmanacLines(data[seedsPart+2:])
	if err != nil {
		log.Fatal("", "err", err)
	}

	for _, seed := range seeds {
		currentValue := ApplyFuncs(almanacEntries, seed)

		if currentValue < result {
			result = currentValue
		}
	}

	return result
}

// Checks if there is a variation between seedB and seedA
// big enought to consider it a changing point.
func isChangingPoint(
	ae map[string]AlmanacEntry,
	seedA, seedB int,
) bool {
	fSeedA := ApplyFuncs(ae, seedA)
	fSeedB := ApplyFuncs(ae, seedB)

	return utils.Abs[int](
		fSeedB-fSeedA,
	) != utils.Abs[int](
		seedB-seedA,
	)
}

func FindSeedPointsToTest(
	srs []SeedRange,
	ae map[string]AlmanacEntry,
) []int {
	result := []int{}

	if len(srs) == 0 {
		return result
	}

	for _, sr := range srs {
		a := sr.Start
		z := sr.End
		l := sr.Len

		switch l {
		case 1:
			result = append(result, a)
			continue
		case 2:
			result = append(result, a)
			result = append(result, z)
			continue
		case 3:
			c := sr.Start + 1

			result = append(result, a)
			result = append(result, z)

			if isChangingPoint(ae, a, c) ||
				isChangingPoint(ae, c, z) {
				result = append(result, c)
			}
			continue
		}

		// 3 extra points, just in case, middle, and 25%
		m := a + (z-a)/2
		d := a + (z-a)/4
		k := m + (z-a)/4

		// Middlepoint correction
		if l%2 == 0 {
			m += 1
			d += 1
			k += 1
		}

		if !isChangingPoint(ae, a, z) && !isChangingPoint(ae, a, m) &&
			!isChangingPoint(ae, d, k) &&
			!isChangingPoint(ae, a, k) {
			result = append(result, a)
			result = append(result, z)
			continue
		}

		ra := NewSeedRange(a, l/2)
		rb := NewSeedRange(m, l/2)

		seedsRA := FindSeedPointsToTest([]SeedRange{ra}, ae)
		seedsRB := FindSeedPointsToTest([]SeedRange{rb}, ae)

		// Get the last seed of ra, and the first of rb
		// If the difference is not good enough, we merge them
		ral := seedsRA[len(seedsRA)-1]
		rbf := seedsRB[0]

		result = append(result, seedsRA...)
		if isChangingPoint(ae, ral, rbf) {
			result = append(result, seedsRB...)
		} else {
			result = append(result, seedsRB[1:]...)
		}
	}

	slices.Sort[[]int](result)
	return result
}

func (e Exercise) Part2(data string) int {
	result := math.MaxInt

	// Prepare the fatal logger, in case an error occurs
	log := log.GetLogger(log.WithPart(2))

	seedsPart := strings.Index(data, "\n\n")
	seedRanges, err := ParseSeedLineAsRanges(data[:seedsPart])
	if err != nil {
		log.Fatal("", "err", err)
	}

	almanacEntries, err := ParseAlmanacLines(data[seedsPart+2:])
	if err != nil {
		log.Fatal("", "err", err)
	}

	srp := FindSeedPointsToTest(seedRanges, almanacEntries)

	for _, seed := range srp {
		value := ApplyFuncs(almanacEntries, seed)

		if value <= result {
			result = value
		}
	}

	return result
}
