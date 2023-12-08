package day8

import (
	"aoc2023/lib"
	"regexp"
	"strings"
)

type Node struct {
	ID       string
	L        string
	R        string
	GhostEnd bool
}

var nodeRe = regexp.MustCompile(`\w+`)

func ParseInput(input []string) ([]rune, map[string]*Node) {
	network := make(map[string]*Node, len(input))

	for _, line := range input[2:] {
		chunks := nodeRe.FindAllString(line, -1)
		network[chunks[0]] = &Node{ID: chunks[0], L: chunks[1], R: chunks[2], GhostEnd: strings.HasSuffix(chunks[0], "Z")}
	}

	return []rune(input[0]), network
}

func Traverse(from, to string, rule []rune, network map[string]*Node) (count int) {
	nodeID := from
	var ruleIndex int

	for {
		nodeID = selectNodeID(network[nodeID], rule[ruleIndex])
		ruleIndex++
		count++

		if nodeID == to {
			return
		}

		if ruleIndex > len(rule)-1 {
			ruleIndex = 0
		}
	}
}

func selectNodeID(node *Node, dir rune) string {
	if dir == 'L' {
		return node.L
	}

	if dir == 'R' {
		return node.R
	}

	panic("unreachable")
}

func GhostTraverse(rule []rune, network map[string]*Node) (count int) {
	startNodes := make([]*Node, 0)
	for nodeID, node := range network {
		if strings.HasSuffix(nodeID, "A") {
			startNodes = append(startNodes, node)
		}
	}

	nodeSteps := make([]int, len(startNodes))
	for index, node := range startNodes {
		nodeSteps[index] = FindGhostNode(node, rule, network)
	}

	return lib.LCM(nodeSteps[0], nodeSteps[1], nodeSteps[2:]...)
}

func FindGhostNode(node *Node, rule []rune, network map[string]*Node) int {
	var step int
	var ruleIndex int

	for {
		node = network[selectNodeID(node, rule[ruleIndex])]
		ruleIndex++
		step++

		if strings.HasSuffix(node.ID, "Z") {
			return step
		}

		if ruleIndex > len(rule)-1 {
			ruleIndex = 0
		}
	}
}

func endCondition(state []*Node) bool {
	for _, node := range state {
		if !node.GhostEnd {
			return false
		}
	}

	return true
}
