package day08

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
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

// Interface to determine if a Node is a start node
type StartNodeDefiner interface {
	Determine(string) bool
}

// Interface to determine if a Node is a end node
type EndNodeDefiner interface {
	Determine(string) bool
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

func NewNode(
	name string,
	snd StartNodeDefiner,
	end EndNodeDefiner,
) *Node {
	return &Node{
		Name:      name,
		startNode: snd.Determine(name),
		endNode:   end.Determine(name),
	}
}

func ParseNodes(
	data string,
	snd StartNodeDefiner,
	end EndNodeDefiner,
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

		log.Debug().
			Str("Node", no.Name).
			Str("Left", nl.Name).
			Str("Right", nr.Name).
			Msg("")

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
		if d == 'R' {
			result = append(result, Right)
		} else if d == 'L' {
			result = append(result, Left)
		} else {
			return result, fmt.Errorf("invalid direction found: %q", d)
		}
	}

	log.Debug().Any("Direction", result).Msg("Directions parsed")

	return result, nil
}

func getOrCreateNode(
	nodes map[string]*Node,
	key string,
	snd StartNodeDefiner,
	end EndNodeDefiner,
) *Node {
	key = strings.TrimSpace(key)
	if node, ok := nodes[key]; ok {
		return node
	}

	node := NewNode(key, snd, end)
	nodes[key] = node

	log.Debug().
		Str("Name", key).
		Bool("IsStartNode?", node.IsStartNode()).
		Bool("IsEndNode?", node.IsEndNode()).
		Msg("Node created")

	return node
}
