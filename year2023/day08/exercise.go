package day08

import (
	"fmt"
	"math"
	"os"
	"strings"

	"go.eryndalor.dev/adventofcode-golang/internal/utils"
)

type Exercise struct{}

func (e Exercise) Part1(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}
	data := string(contents)

	lineBreak := strings.Index(data, "\n")
	if lineBreak == -1 {
		return 0, fmt.Errorf("Part1: invalid file")
	}

	directions, err := ParseDirectionsLine(data[:lineBreak])
	if err != nil {
		return 0, fmt.Errorf("Part1: %w", err)
	}

	startNodes, err := ParseNodes(
		data[lineBreak+1:],
		AAAStartNode{},
		ZZZEndNode{},
	)

	if err != nil || len(startNodes) != 1 {
		return 0, fmt.Errorf(
			"Part1: %w || len(startNodes) = %d",
			err,
			len(startNodes),
		)
	}

	current := startNodes[0]

	return findStepsToEndNode(directions, current), nil
}

func (e Exercise) Part2(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}
	data := string(contents)

	lineBreak := strings.Index(data, "\n")
	if lineBreak == -1 {
		return 0, fmt.Errorf("Part2: invalid file")
	}

	directions, err := ParseDirectionsLine(data[:lineBreak])
	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}

	startNodes, err := ParseNodes(
		data[lineBreak+1:],
		GhostStartNode{},
		GhostEndNode{},
	)

	if err != nil {
		return 0, fmt.Errorf("Part2: %w", err)
	}

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

		minStepsPerNode = append(minStepsPerNode, s)
	}

	return utils.LCM[int](minStepsPerNode...), nil
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
				fmt.Fprintln(os.Stderr, "Unable to continue, the logic should be reviewed")
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
