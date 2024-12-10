package day02

import (
	"math"
)

func AreLevelsSafe(levels []int) bool {
	if len(levels) < 3 {
		return false
	}

	a := levels[0]
	b := levels[1]

	if a == b || math.Abs(float64(a-b)) > 3 {
		return false
	}

	increasing := true
	if b < a {
		increasing = false
	}

	last := b

	for _, c := range levels[2:] {
		if c == last ||
			(c > last && !increasing) ||
			(c < last && increasing) ||
			math.Abs(float64(c-last)) > 3 {
			return false
		}

		last = c
	}

	return true
}
