package day08

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

// Direction represents the left and right options
// that are available
type Direction int

const (
	Right Direction = iota
	Left
)

type Node struct {
	Name      string
	Left      *Node
	Right     *Node
	startNode bool
	endNode   bool
}

func NewNode(
	name string,
	snd NodeDefiner,
	end NodeDefiner,
) *Node {
	return &Node{
		Name:      name,
		startNode: snd.Determine(name),
		endNode:   end.Determine(name),
	}
}

// Interface to determine if a Node is a start node
type NodeDefiner interface {
	Determine(name string) bool
}

func (n Node) String() string {
	return n.Name
}

func (n Node) IsStartNode() bool {
	return n.startNode
}

func (n Node) IsEndNode() bool {
	return n.endNode
}

func ParseNodes(
	data string,
	snd NodeDefiner,
	end NodeDefiner,
) (startNodes []*Node, err error) {
	var nodes = make(map[string]*Node)

	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		nodeDefs := strings.Split(line, "=")
		if len(nodeDefs) != 2 {
			return nil, fmt.Errorf(
				"invalid line: %q",
				line,
			)
		}

		o := strings.TrimSpace(nodeDefs[0])
		d := strings.TrimSpace(nodeDefs[1])

		ds := strings.Split(d, ",")
		if len(ds) != 2 {
			return nil, fmt.Errorf(
				"invalid children: %q",
				d,
			)
		}

		l := ds[0][1:]
		r := ds[1][:len(ds[1])-1]

		// Create the nodes if they don't exist
		no := getOrCreateNode(nodes, o, snd, end)
		nl := getOrCreateNode(nodes, l, snd, end)
		nr := getOrCreateNode(nodes, r, snd, end)

		no.Left = nl
		no.Right = nr
	}

	// All nodes processed, fetching start nodes
	for _, n := range nodes {
		if n.IsStartNode() {
			startNodes = append(startNodes, n)
		}
	}

	if len(startNodes) == 0 {
		return nil, errors.New("start node not found")
	}

	return startNodes, nil
}

func ParseDirectionsLine(line string) ([]Direction, error) {
	line = strings.TrimSpace(line)
	var result = make([]Direction, 0, len(line))

	for _, d := range line {
		switch d {
		case 'R':
			result = append(result, Right)
		case 'L':
			result = append(result, Left)
		default:
			return result, fmt.Errorf("invalid direction found: %q", d)
		}
	}

	return result, nil
}

func getOrCreateNode(
	nodes map[string]*Node,
	key string,
	snd NodeDefiner,
	end NodeDefiner,
) *Node {
	key = strings.TrimSpace(key)
	if node, ok := nodes[key]; ok {
		return node
	}

	node := NewNode(key, snd, end)
	nodes[key] = node

	return node
}
