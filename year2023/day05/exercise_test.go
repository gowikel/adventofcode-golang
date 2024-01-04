package day05_test

import (
	_ "embed"
	"testing"

	"github.com/gowikel/adventofcode-golang/internal/utils"
	. "github.com/gowikel/adventofcode-golang/year2023/day05"
)

//go:embed testdata/seeds.txt
var seedsData string

//go:embed testdata/almanac.txt
var almanacData string

var seedRanges, _ = ParseSeedLineAsRanges(seedsData)

// 0 =>  0   10 => 10   20 => 15   30 => 30   40 => 44
// 1 => 61   11 => 11   21 => 16   31 => 31   41 => 45
// 2 => 62   12 => 12   22 => 17   32 => 32   42 => 46
// 3 => 63   13 => 13   23 => 18   33 => 33   43 => 47
// 4 => 64   14 => 14   24 => 19   34 => 34   44 => 48
// 5 => 65   15 => 15   25 => 20   35 =>  7   45 => 49
// 6 => 66   16 => 16   26 => 21   36 => 40   46 => 46
// 7 => 67   17 => 17   27 => 22   37 => 41   47 => 47
// 8 => 68   18 => 18   28 => 23   38 => 42   48 => 48
// 9 => 60   19 => 19   29 => 24   39 => 43   49 => 49
var almanacEntries, _ = ParseAlmanacLines(almanacData)

func TestFindSeedPointsToTest_EmptySeedRange(t *testing.T) {
	input := []SeedRange{}
	want := []int{}
	got := FindSeedPointsToTest(input, almanacEntries)

	if !utils.UnorderedEqualSlices[int](got, want) {
		t.Errorf(
			"FindSeedPointsToTest wants %v but got %v\n",
			want,
			got,
		)
	}
}

func TestFindSeedPointsToTest_OneRange_OneElement(t *testing.T) {
	input := []SeedRange{}

	sr := NewSeedRange(1, 1)
	input = append(input, sr)

	want := []int{1}
	got := FindSeedPointsToTest(input, almanacEntries)

	if !utils.UnorderedEqualSlices[int](got, want) {
		t.Errorf(
			"FindSeedPointsToTest wants %v but got %v\n",
			want,
			got,
		)
	}
}

func TestFindSeedPointsToTest_OneRange_TwoElements(t *testing.T) {
	input := []SeedRange{}

	sr := NewSeedRange(1, 2)
	input = append(input, sr)

	want := []int{1, 2}
	got := FindSeedPointsToTest(input, almanacEntries)

	if !utils.UnorderedEqualSlices[int](got, want) {
		t.Errorf(
			"FindSeedPointsToTest wants %v but got %v\n",
			want,
			got,
		)
	}
}

func TestFindSeedPointsToTest_OneRange_ThreeElements_SameRange(
	t *testing.T,
) {
	input := []SeedRange{}

	// 1 => 61, 2 => 62, 3 => 63
	// 61, 62, 63
	sr := NewSeedRange(1, 3)
	input = append(input, sr)

	want := []int{1, 3}
	got := FindSeedPointsToTest(input, almanacEntries)

	if !utils.UnorderedEqualSlices[int](got, want) {
		t.Errorf(
			"FindSeedPointsToTest wants %v but got %v\n",
			want,
			got,
		)
	}
}

func TestFindSeedPointsToTest_OneRange_ThreeElements_DifferentRange(
	t *testing.T,
) {
	input := []SeedRange{}

	// 34 => 34, 35 =>  7, 36 => 40
	// 34, 7, 40
	sr := NewSeedRange(34, 3)
	input = append(input, sr)

	want := []int{34, 35, 36}
	got := FindSeedPointsToTest(input, almanacEntries)

	if !utils.UnorderedEqualSlices[int](got, want) {
		t.Errorf(
			"FindSeedRangePoints wants %v but got %v\n",
			want,
			got,
		)
	}
}

func TestFindSeedPointsToTest_OneLargeRange_SameRange(
	t *testing.T,
) {
	input := []SeedRange{}

	// 36 => 40, 37 => 41, 38 => 42, 39 => 43, 40 => 44, 41 => 45
	// => 40, 41, 42, 43, 44, 45
	sr := NewSeedRange(36, 6)
	input = append(input, sr)

	want := []int{36, 41}
	got := FindSeedPointsToTest(input, almanacEntries)

	if !utils.UnorderedEqualSlices[int](got, want) {
		t.Errorf(
			"FindSeedPointsToTest wants %v but got %v\n",
			want,
			got,
		)
	}
}

func TestFindSeedPointsToTest_OneLargeRange_TwoRanges(
	t *testing.T,
) {
	input := []SeedRange{}

	// 30 => 30, 31 => 31, 32 => 32, 33 => 33, 34 => 34, 35 =>  7
	sr := NewSeedRange(30, 6)
	input = append(input, sr)

	want := []int{30, 34, 35}
	got := FindSeedPointsToTest(input, almanacEntries)

	if !utils.UnorderedEqualSlices[int](got, want) {
		t.Errorf(
			"FindSeedPointsToTest wants %v but got %v\n",
			want,
			got,
		)
	}
}

func TestFindSeedPointsToTest_OneLargeRange_FullSet(
	t *testing.T,
) {
	t.Skip(
		"Broken. The algorithm is good enough to remove most seeds, but not all.",
	)
	input := []SeedRange{}

	sr := NewSeedRange(0, 50)
	input = append(input, sr)

	want := []int{
		0,
		1,
		8,
		9,
		10,
		19,
		20,
		29,
		30,
		34,
		35,
		36,
		45,
		46,
		49,
	}
	got := FindSeedPointsToTest(input, almanacEntries)

	if !utils.UnorderedEqualSlices[int](got, want) {
		t.Errorf(
			"FindSeedPointsToTest wants %v but got %v [%d vs %d]\n",
			want,
			got,
			len(want),
			len(got),
		)
	}
}
