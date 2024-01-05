package day08

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day08/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var lines = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var instructions = ParseInstructions(lines[0])
	var root = BuildNetwork(lines[2:])
	var stop = func(n *Node) bool { return n.Label == "ZZZ" }

	var numSteps = TraverseNetwork(root, instructions, stop)

	answer = numSteps
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day08/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var lines = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var instructions = ParseInstructions(lines[0])
	var roots = BuildGhostNetwork(lines[2:])
	var stop = func(n *Node) bool { return n.Label[len(n.Label)-1] == 'Z' }

	var result = 1

	for _, root := range roots {
		var numSteps = TraverseNetwork(root, instructions, stop)
		result = LeastCommonMultiple(result, numSteps)
	}

	answer = result
	return answer, time.Since(now)
}

type Instruction int

const (
	Left  Instruction = 0
	Right Instruction = 1
)

type Node struct {
	Label          string
	IsStartingNode bool
	IsEndingNode   bool
	Left           *Node
	Right          *Node
}

func NewNode(label string) *Node {
	return &Node{
		Label: label,
	}
}

func ParseNodeDefinition(line string) (label, leftLabel, rightLabel string) {
	// AAA = (BBB, CCC)
	label = line[0:3]
	leftLabel = line[7:10]
	rightLabel = line[12:15]
	return
}

func ParseInstructions(line string) []Instruction {
	// LRRLRRLRLLLRLLRLRRLRRLRRL
	var instructions = []Instruction{}

	for _, val := range line {
		if val == 'L' {
			instructions = append(instructions, Left)
		} else {
			instructions = append(instructions, Right)
		}
	}

	return instructions
}

func BuildNetwork(lines []string) *Node {
	// AAA = (BBB, BBB)
	// BBB = (AAA, ZZZ)
	// ZZZ = (ZZZ, ZZZ)

	var nodes = map[string]*Node{}
	for _, line := range lines {
		var label, leftLabel, rightLabel = ParseNodeDefinition(line)
		var node = nodes[label]
		if node == nil {
			node = NewNode(label)
			nodes[label] = node
		}

		var leftNode = nodes[leftLabel]
		if leftNode == nil {
			leftNode = NewNode(leftLabel)
			nodes[leftLabel] = leftNode
		}

		node.Left = leftNode

		var rightNode = nodes[rightLabel]
		if rightNode == nil {
			rightNode = NewNode(rightLabel)
			nodes[rightLabel] = rightNode
		}

		node.Right = rightNode
	}

	return nodes["AAA"]
}

func TraverseNetwork(root *Node, instructions []Instruction, stop func(n *Node) bool) (numSteps int) {
	numSteps = 0
	var currentNode = root
	var i = 0

	for !stop(currentNode) {
		var instruction = instructions[i]
		if instruction == Left {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}

		numSteps++

		i++
		if i >= len(instructions) {
			i = 0
		}
	}

	return
}

func BuildGhostNetwork(lines []string) (startingNodes []*Node) {
	// 11A = (11B, XXX)
	// 11B = (XXX, 11Z)
	// 11Z = (11B, XXX)
	// 22A = (22B, XXX)
	// 22B = (22C, 22C)
	// 22C = (22Z, 22Z)
	// 22Z = (22B, 22B)
	// XXX = (XXX, XXX)

	var nodes = map[string]*Node{}
	for _, line := range lines {
		var label, leftLabel, rightLabel = ParseNodeDefinition(line)
		var node = nodes[label]
		if node == nil {
			node = NewNode(label)
			nodes[label] = node
		}

		if node.Label[len(node.Label)-1] == 'A' {
			startingNodes = append(startingNodes, node)
		}

		var leftNode = nodes[leftLabel]
		if leftNode == nil {
			leftNode = NewNode(leftLabel)
			nodes[leftLabel] = leftNode
		}

		node.Left = leftNode

		var rightNode = nodes[rightLabel]
		if rightNode == nil {
			rightNode = NewNode(rightLabel)
			nodes[rightLabel] = rightNode
		}

		node.Right = rightNode
	}

	return
}

func LeastCommonMultiple(a, b int) int {
	return (a * b) / GreatestCommonDivisor(a, b)
}

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
