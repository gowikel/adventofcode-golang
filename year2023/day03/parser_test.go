package day03_test

import (
	"testing"

	"github.com/gowikel/adventofcode-golang/utils/algo"
	. "github.com/gowikel/adventofcode-golang/year2023/day03"
)

type GetPointsTestCase struct {
	description string
	input       string
	expected    [][2]int
}

type GetRangesTestCase struct {
	description string
	input       string
	points      [][2]int
	expected    [][3]int
}

type LocateNumbersTestCase struct {
	description string
	input       string
	ranges      [][3]int
	expected    [][3]int
}

type GetGearsTestCase struct {
	description string
	input       string
	expected    [][2]int
}

//	0123456789
//
// 0 123......9
// 1 10*.......
// 2 11.456.78.
// 3 ..........
// 4 ...#......
// 5 ........09
// 6 666....33$
const input = "123......9\n" +
	"10*.......\n" +
	"11.456.78.\n" +
	"..........\n" +
	"...#......\n" +
	"........09\n" +
	"666....33$\n"

func TestGetPoints(t *testing.T) {
	testCases := []GetPointsTestCase{
		{
			description: "Symbols are located in a string",
			input:       "...123...+\n...*45....\n67...&....",
			expected: [][2]int{
				{0, 9},
				{1, 3},
				{2, 5},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := GetPoints(testCase.input)
			want := testCase.expected

			if !algo.UnorderedEqualSlices[[2]int](got, want) {
				t.Errorf(
					"\n\n%s\n\nInput: %v\nGot: %v\nWant: %v\n\n",
					testCase.description,
					testCase.input,
					got,
					testCase.expected,
				)
			}
		})
	}
}

//   0123456789
// 0 123......9
// 1 10*.......
// 2 11.456.78.
// 3 ..........
// 4 ...#......
// 5 ........09
// 6 666....33$

func TestGetRanges(t *testing.T) {
	testCases := []GetRangesTestCase{
		{
			description: "No points, no ranges",
			input:       input,
			points:      [][2]int{},
			expected:    [][3]int{},
		},
		{
			description: "A point on the first line will trigger " +
				"a range in that line and the next",
			input: input,
			points: [][2]int{
				{0, 1},
				{0, 5},
			},
			expected: [][3]int{
				{0, 0, 3},
				{1, 0, 3},
				{0, 4, 7},
				{1, 4, 7},
			},
		},
		{
			description: "A point on the last line will trigger " +
				"a range in that line and the previous one",
			input: input,
			points: [][2]int{
				{6, 1},
			},
			expected: [][3]int{
				{5, 0, 3},
				{6, 0, 3},
			},
		},
		{
			description: "A point not close to any boundary will trigger " +
				"a range in the previous line, that line and the next one",
			input: input,
			points: [][2]int{
				{2, 1},
			},
			expected: [][3]int{
				{1, 0, 3},
				{2, 0, 3},
				{3, 0, 3},
			},
		},
		{
			description: "Points at the beginning will have shorter ranges",
			input:       input,
			points: [][2]int{
				{0, 0},
				{3, 0},
			},
			expected: [][3]int{
				{0, 0, 2},
				{1, 0, 2},
				{2, 0, 2},
				{3, 0, 2},
				{4, 0, 2},
			},
		},
		{
			description: "Points at the end of the line will have shorter ranges",
			input:       input,
			points: [][2]int{
				{0, 9},
				{3, 9},
				{6, 9},
			},
			expected: [][3]int{
				{0, 8, 9},
				{1, 8, 9},
				{2, 8, 9},
				{3, 8, 9},
				{4, 8, 9},
				{5, 8, 9},
				{6, 8, 9},
			},
		},
		{
			description: "Close points will cause overlaping ranges",
			input:       input,
			points: [][2]int{
				{1, 1},
				{1, 2},
				{2, 1},
			},
			expected: [][3]int{
				{0, 0, 3},
				{1, 0, 3},
				{2, 0, 3},
				{0, 1, 4},
				{1, 1, 4},
				{2, 1, 4},
				{1, 0, 3},
				{2, 0, 3},
				{3, 0, 3},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := GetRanges(testCase.input, testCase.points)
			want := testCase.expected

			if !algo.UnorderedEqualSlices[[3]int](got, want) {
				t.Errorf(
					"\n\n%s\n\nInput:\n%v\nPoints: %v\nGot: %v\nWant: %v\n\n",
					testCase.description,
					testCase.input,
					testCase.points,
					got,
					want,
				)
			}
		})
	}
}

//   0123456789
// 0 123......9
// 1 10*.......
// 2 11.456.78.
// 3 ..........
// 4 ...#......
// 5 ........09
// 6 666....33$

func TestLocateNumbers(t *testing.T) {
	testCases := []LocateNumbersTestCase{
		{
			description: "A range over a number will just return it",
			input:       input,
			ranges: [][3]int{
				{0, 0, 2},
				{2, 7, 8},
			},
			expected: [][3]int{
				{0, 0, 3},
				{2, 7, 9},
			},
		},
		{
			description: "A range partially over a number, will expand to the cover the number",
			input:       input,
			ranges: [][3]int{
				{0, 0, 1},
				{1, 1, 2},
				{2, 2, 4},
				{6, 2, 4},
			},
			expected: [][3]int{
				{0, 0, 3},
				{1, 0, 2},
				{2, 3, 6},
				{6, 0, 3},
			},
		},
		{
			description: "A range that spans multiple numbers, will return all of them",
			input:       input,
			ranges: [][3]int{
				{2, 0, 6},
			},
			expected: [][3]int{
				{2, 0, 2},
				{2, 3, 6},
			},
		},
		{
			description: "An empty range will not return anything",
			input:       input,
			ranges: [][3]int{
				{0, 0, 0},
			},
			expected: [][3]int{},
		},
		{
			description: "Multiple ranges over the same number do not repeat the number",
			input:       input,
			ranges: [][3]int{
				{0, 0, 1},
				{0, 1, 2},
			},
			expected: [][3]int{
				{0, 0, 3},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := LocateNumbers(testCase.input, testCase.ranges)
			want := testCase.expected

			if !algo.UnorderedEqualSlices[[3]int](got, want) {
				t.Errorf(
					"\n\n%s\n\nInput:\n%v\nRanges: %v\nGot: %v\nWant: %v\n\n",
					testCase.description,
					testCase.input,
					testCase.ranges,
					got,
					want,
				)
			}
		})
	}
}

func TestGetGears(t *testing.T) {
	testCases := []GetGearsTestCase{
		{
			description: "Empty input, no points",
			input:       "",
			expected:    [][2]int{},
		},
		{
			description: "* positions are returned",
			input: ".*%$.*.2.*\n" +
				"..........\n" +
				"...*&*....",
			expected: [][2]int{
				{0, 1},
				{0, 5},
				{0, 9},
				{2, 3},
				{2, 5},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := GetGears(testCase.input)
			want := testCase.expected

			if !algo.UnorderedEqualSlices[[2]int](got, want) {
				t.Errorf(
					"\n\n%s\n\nInput:\n%s\nGot: %v\nWant: %v\n\n",
					testCase.description,
					testCase.input,
					got,
					want,
				)
			}
		})
	}
}
