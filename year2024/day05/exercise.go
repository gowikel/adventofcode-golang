package day05

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	pageData, err := Parse(path)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}

	result := 0

outer:
	for _, pages := range pageData.Pages {
		// First, we check the first page
		writtenPages := mapset.NewSetWithSize[int](len(pages))

		for _, page := range pages {
			beforeRules := pageData.BeforeRules[page]
			writtenPages.Add(page)

			for _, beforeRule := range beforeRules {
				if writtenPages.Contains(beforeRule) {
					continue outer
				}
			}
		}

		middlePage := pages[len(pages)/2]
		result += middlePage
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	pageData, err := Parse(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}

	result := 0
	wrongPagesLst := make([][]int, 0)

	// Extract wrong page lists
outer:
	for idx, pages := range pageData.Pages {
		writtenPages := mapset.NewSetWithSize[int](len(pages))

		for _, page := range pages {
			beforeRules := pageData.BeforeRules[page]
			writtenPages.Add(page)

			for _, beforeRule := range beforeRules {
				if writtenPages.Contains(beforeRule) {
					wrongPagesLst = append(wrongPagesLst, pageData.Pages[idx])
					continue outer
				}
			}
		}
	}

	for _, pageLst := range wrongPagesLst {
		cont := true
		sortedLst := pageLst

		// Sort the page list, one page at a time.
		for cont {
			sortedLst, cont = sortManualPages(
				sortedLst,
				pageData.BeforeRules,
			)
		}

		// Get the middle page and sum them up
		middle := sortedLst[len(sortedLst)/2]
		result += middle
	}

	return result, nil
}
