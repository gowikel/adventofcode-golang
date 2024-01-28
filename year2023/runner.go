package year2023

import (
	"log/slog"
	"os"

	"github.com/gowikel/adventofcode-golang/internal/puzzle"
	D01 "github.com/gowikel/adventofcode-golang/year2023/day01"
	D02 "github.com/gowikel/adventofcode-golang/year2023/day02"
	D03 "github.com/gowikel/adventofcode-golang/year2023/day03"
	D04 "github.com/gowikel/adventofcode-golang/year2023/day04"
	D05 "github.com/gowikel/adventofcode-golang/year2023/day05"
	D06 "github.com/gowikel/adventofcode-golang/year2023/day06"
	D07 "github.com/gowikel/adventofcode-golang/year2023/day07"
	D08 "github.com/gowikel/adventofcode-golang/year2023/day08"
	D09 "github.com/gowikel/adventofcode-golang/year2023/day09"
)

type solver interface {
	Part1(data string) int
	Part2(data string) int
}

var solvers = map[int]solver{
	1: D01.Exercise{},
	2: D02.Exercise{},
	3: D03.Exercise{},
	4: D04.Exercise{},
	5: D05.Exercise{},
	6: D06.Exercise{},
	7: D07.Exercise{},
	8: D08.Exercise{},
	9: D09.Exercise{},
}

func Run(day int, data string, pps puzzle.PuzzleRunSelector) {
	solver, ok := solvers[day]

	if !ok {
		slog.Error("Solver not implemented", "year", 2023, "day", day)
		os.Exit(1)
	}

	switch pps {
	case puzzle.RunAll:
		p1 := solver.Part1(data)
		p2 := solver.Part2(data)

		slog.Info(
			"Run completed",
			"year",
			2023,
			"day",
			day,
			"part1",
			p1,
			"part2",
			p2,
		)
	case puzzle.RunPartOne:
		p1 := solver.Part1(data)

		slog.Info(
			"Run completed",
			"year",
			2023,
			"day",
			day,
			"part1",
			p1,
		)
	case puzzle.RunPartTwo:
		p2 := solver.Part2(data)

		slog.Info(
			"Run completed",
			"year",
			2023,
			"day",
			day,
			"part2",
			p2,
		)
	}
}
