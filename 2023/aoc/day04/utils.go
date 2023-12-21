package day04

import "fmt"

func countMatches(w map[int]empty, lst []int) (int, error) {
	if w == nil || lst == nil {
		return 0, fmt.Errorf("map or slice is nil")
	}

	var result int

	for _, n := range lst {
		if _, ok := w[n]; ok {
			result += 1
		}
	}

	return result, nil
}