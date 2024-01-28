package day08

import (
	"log/slog"
	"math"
	"os"
	"strings"

	"github.com/gowikel/adventofcode-golang/internal/utils"
)

type Exercise struct{}

func (e Exercise) Part1(data string) int {
	lineBreak := strings.Index(data, "\n")
	if lineBreak == -1 {
		slog.Error("Invalid file")
		os.Exit(1)
	}

	directions, err := ParseDirectionsLine(data[:lineBreak])
	if err != nil {
		slog.Error(
			"Error while parsing the directions line",
			"err",
			err,
		)
		os.Exit(1)
	}

	startNodes, err := ParseNodes(
		data[lineBreak+1:],
		AAAStartNode{},
		ZZZEndNode{},
	)

	if err != nil || len(startNodes) != 1 {
		slog.Error("Error while parsing", "err", err)
	}

	current := startNodes[0]

	slog.Debug("Parse completed")

	return findStepsToEndNode(directions, current)
}

func (e Exercise) Part2(data string) int {
	lineBreak := strings.Index(data, "\n")
	if lineBreak == -1 {
		slog.Error("Invalid file")
		os.Exit(1)
	}

	directions, err := ParseDirectionsLine(data[:lineBreak])
	if err != nil {
		slog.Error(
			"Error while parsing the directions line",
			"err",
			err,
		)
		os.Exit(1)
	}

	startNodes, err := ParseNodes(
		data[lineBreak+1:],
		GhostStartNode{},
		GhostEndNode{},
	)

	if err != nil {
		slog.Error("erro while parsing nodes", "err", err)
		os.Exit(1)
	}

	slog.Debug("Parse completed")

	// This time, the amount of nodes, and the fact that I have
	// to reach the end of every single node at the same time
	// means that I cannot brute force it.
	//
	// However, there is a simple trick to exploit. Once a node
	// reaches the end node, it will either be kept forever there
	// or it will start again in the start node.
	//
	// This means that, if I detect the first time that each node
	// reaches its end node, then step that will make all reach the
	// end goal will be the Least Common Multiple of all the previous
	// numbers.
	//
	// To make things easier, I extracted the logic on the part1 into
	// a function, as we will reuse it to find the minimum number
	// of steps for each node.

	minStepsPerNode := make([]int, 0, len(startNodes))
	for _, n := range startNodes {
		s := findStepsToEndNode(directions, n)

		slog.Debug("", "mode", n.Name, "minSteps", s)

		minStepsPerNode = append(minStepsPerNode, s)
	}

	slog.Debug("", "steps", minStepsPerNode)

	return utils.LCM[int](minStepsPerNode...)
}

func findStepsToEndNode(ds []Direction, n *Node) int {
	var result int

	current := n

Loop:
	for !current.IsEndNode() {
		for _, direction := range ds {
			result += 1
			if direction == Right {
				current = current.Right
			} else {
				current = current.Left
			}

			if current.IsEndNode() {
				break Loop
			}

			if result == math.MaxInt {
				slog.Error("Unable to continue, the logic should be reviewed", "year", 2023, "day", 8, "part", 1)
				break Loop
			}

		}

	}

	return result
}

// Defines a start node as a node whose name is AAA
type AAAStartNode struct{}

func (a AAAStartNode) Determine(name string) bool {
	return name == "AAA"
}

// Defines an end node as a node whose name is ZZZ
type ZZZEndNode struct{}

func (z ZZZEndNode) Determine(name string) bool {
	return name == "ZZZ"
}

// Defines a start node as a node whose name ends in A
type GhostStartNode struct{}

func (g GhostStartNode) Determine(name string) bool {
	return strings.HasSuffix(name, "A")
}

// Defines an end node as a node whose name ends in Z
type GhostEndNode struct{}

func (g GhostEndNode) Determine(name string) bool {
	return strings.HasSuffix(name, "Z")
}
