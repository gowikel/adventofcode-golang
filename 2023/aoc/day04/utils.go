package day04

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

func countMatches(w mapset.Set[int], lst []int) (int, error) {
	if w == nil || lst == nil {
		return 0, fmt.Errorf("map or slice is nil")
	}

	var result int

	for _, n := range lst {
		if w.Contains(n) {
			result += 1
		}
	}

	return result, nil
}