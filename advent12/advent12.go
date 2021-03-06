package advent12

import (
	"strings"

	"advent2021/util"
)

const startName = "start"
const endName = "end"

func Solution(inputFile string) (part1, part2 interface{}) {
	caveStart := parseCave(inputFile)
	part1Paths := extendPaths([]*CaveNode{caveStart}, false)
	part2Paths := extendPaths([]*CaveNode{caveStart}, true)
	return len(part1Paths), len(part2Paths)
}

type (
	CaveNode struct {
		Name      string
		Neighbors map[string]*CaveNode
	}
)

func (c *CaveNode) IsSmallCave() bool {
	return c.Name != endName && c.Name == strings.ToLower(c.Name)
}

func extendPaths(prev []*CaveNode, smallCaveVisit bool) [][]*CaveNode {
	paths := make([][]*CaveNode, 0)
	last := prev[len(prev)-1]
	for _, neighbor := range last.Neighbors {
		if neighbor.Name == startName {
			continue
		}
		futureSmallCaveVisit := smallCaveVisit
		if neighbor.IsSmallCave() && caveNodeInPath(neighbor, prev) {
			if futureSmallCaveVisit {
				futureSmallCaveVisit = false
			} else {
				continue
			}
		}

		path := append(prev, neighbor)

		if neighbor.Name == endName {
			paths = append(paths, path)
		} else {
			paths = append(paths, extendPaths(path, futureSmallCaveVisit)...)
		}
	}
	return paths
}

func caveNodeInPath(node *CaveNode, path []*CaveNode) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

func parseCave(inputFile string) *CaveNode {
	nodesByName := make(map[string]*CaveNode)

	lines := util.ReadFile(inputFile)
	for _, line := range lines {
		nodes := strings.Split(line, "-")
		if len(nodes) != 2 {
			continue
		}
		left, right := nodes[0], nodes[1]

		if _, ok := nodesByName[left]; !ok {
			nodesByName[left] = &CaveNode{
				Name:      left,
				Neighbors: make(map[string]*CaveNode),
			}
		}
		if _, ok := nodesByName[right]; !ok {
			nodesByName[right] = &CaveNode{
				Name:      right,
				Neighbors: make(map[string]*CaveNode),
			}
		}

		leftNode := nodesByName[left]
		rightNode := nodesByName[right]

		leftNode.Neighbors[right] = rightNode
		rightNode.Neighbors[left] = leftNode
	}

	return nodesByName[startName]
}
