package day05

import (
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
)

// Given a list of pages, and a set of before Rules, it will move one
// page if it is in the wrong order. It returns the modified list of
// pages and a boolean, that indicates if any movement has happened.
func sortManualPages(pages []int, rules map[int][]int) ([]int, bool) {
	sorted := make([]int, 0, len(pages))
	written := mapset.NewSetWithSize[int](len(pages))
	pageMoved := false

	for _, p := range pages {
		r := rules[p]
		sorted = append(sorted, p)
		written.Add(p)

		if !pageMoved {
			for _, before := range r {
				if written.Contains(before) {
					pageMoved = true

					idx := slices.Index(sorted, before)
					sorted[idx] = p
					sorted[len(sorted)-1] = before
					break
				}
			}
		}
	}

	return sorted, pageMoved
}
