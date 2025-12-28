package day08

import (
	"fmt"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	data, err := parseFile(path)
	if err != nil {
		return 0, fmt.Errorf("part1: %w", err)
	}

	gameMap := make(map[int]map[int]struct{})
	var result int

	for _, antennas := range data.RoofMap {
		for _, a := range antennas {

			for _, b := range antennas {
				if a == b {
					continue
				}

				Δx := b[0] - a[0]
				Δy := b[1] - a[1]

				x := a[0] - Δx
				y := a[1] - Δy

				if x < 0 || y < 0 || x >= data.Rows || y >= data.Columns {
					continue
				}

				if _, ok := gameMap[x]; !ok {
					gameMap[x] = make(map[int]struct{})
				}

				if _, ok := gameMap[x][y]; ok {
					continue
				}
				gameMap[x][y] = struct{}{}
				result++
			}
		}
	}

	return result, nil
}

func (e Exercise) Part2(path string) (int, error) {
	data, err := parseFile(path)
	if err != nil {
		return 0, fmt.Errorf("part2: %w", err)
	}

	var result int
	gameMap := make(map[int]map[int]struct{})

	for _, antennas := range data.RoofMap {
		for _, a := range antennas {
			if _, ok := gameMap[a[0]]; !ok {
				gameMap[a[0]] = make(map[int]struct{})
			}

			for _, b := range antennas {
				if a == b {
					continue
				}

				Δx := b[0] - a[0]
				Δy := b[1] - a[1]

				x := a[0]
				y := a[1]

				if _, ok := gameMap[x][y]; !ok {
					result++
					gameMap[x][y] = struct{}{}
				}

				for x+Δx >= 0 && y+Δy >= 0 && x+Δx < data.Rows && y+Δy < data.Columns {
					x += Δx
					y += Δy

					if _, ok := gameMap[x]; !ok {
						gameMap[x] = make(map[int]struct{})
					}

					if _, ok := gameMap[x][y]; !ok {
						result++
						gameMap[x][y] = struct{}{}
					}
				}
			}
		}
	}

	return result, nil
}
