package day08_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day08"
)

func TestNewNode(t *testing.T) {
	var node = day08.NewNode("AAA")
	assert.True(t, node.IsStartingNode)
	assert.False(t, node.IsEndingNode)

	node = day08.NewNode("XYZ")
	assert.False(t, node.IsStartingNode)
	assert.True(t, node.IsEndingNode)

	node = day08.NewNode("DEF")
	assert.False(t, node.IsStartingNode)
	assert.False(t, node.IsEndingNode)
}

func TestParseNodeDefinition(t *testing.T) {
	var label, leftLabel, rightLabel = day08.ParseNodeDefinition("AAA = (BBB, CCC)")
	assert.Equal(t, "AAA", label)
	assert.Equal(t, "BBB", leftLabel)
	assert.Equal(t, "CCC", rightLabel)
}

func TestParseInstructions(t *testing.T) {
	var instructions = day08.ParseInstructions("LRRL")
	assert.Equal(t, 4, len(instructions))
	assert.Equal(t, day08.Left, instructions[0])
	assert.Equal(t, day08.Right, instructions[1])
	assert.Equal(t, day08.Right, instructions[2])
	assert.Equal(t, day08.Left, instructions[3])
}

func TestBuildNetwork_Simple(t *testing.T) {
	var lines = []string{
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	var root = day08.BuildNetwork(lines)
	assert.Equal(t, "AAA", root.Label)

	var left = root.Left
	var right = root.Right
	assert.Equal(t, "BBB", left.Label)
	assert.Equal(t, "BBB", right.Label)
	assert.True(t, left == right)

	var bbb = left
	left = bbb.Left
	right = bbb.Right
	assert.True(t, left == root)
	assert.Equal(t, "ZZZ", right.Label)

	var zzz = bbb.Right
	left = zzz.Left
	right = zzz.Right
	assert.True(t, left == right)
}

func TestBuildNetwork(t *testing.T) {
	var lines = []string{
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	var aaa = day08.BuildNetwork(lines)
	assert.Equal(t, "AAA", aaa.Label)
	assert.Equal(t, "BBB", aaa.Left.Label)
	assert.Equal(t, "CCC", aaa.Right.Label)

	var bbb = aaa.Left
	assert.Equal(t, "BBB", bbb.Label)
	assert.Equal(t, "DDD", bbb.Left.Label)
	assert.Equal(t, "EEE", bbb.Right.Label)

	var ccc = aaa.Right
	assert.Equal(t, "CCC", ccc.Label)
	assert.Equal(t, "ZZZ", ccc.Left.Label)
	assert.Equal(t, "GGG", ccc.Right.Label)

	var ddd = bbb.Left
	assert.Equal(t, "DDD", ddd.Label)
	assert.Equal(t, "DDD", ddd.Left.Label)
	assert.Equal(t, "DDD", ddd.Right.Label)

	var eee = bbb.Right
	assert.Equal(t, "EEE", eee.Label)
	assert.Equal(t, "EEE", eee.Left.Label)
	assert.Equal(t, "EEE", eee.Right.Label)

	var ggg = ccc.Right
	assert.Equal(t, "GGG", ggg.Label)
	assert.Equal(t, "GGG", ggg.Left.Label)
	assert.Equal(t, "GGG", ggg.Right.Label)

	var zzz = ccc.Left
	assert.Equal(t, "ZZZ", zzz.Label)
	assert.Equal(t, "ZZZ", zzz.Left.Label)
	assert.Equal(t, "ZZZ", zzz.Right.Label)
}

func TestTraverseNetwork(t *testing.T) {
	var instructions = day08.ParseInstructions("RL")

	var network = []string{
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	var root = day08.BuildNetwork(network)
	var stop = func(n *day08.Node) bool { return n.Label == "ZZZ" }
	var numSteps = day08.TraverseNetwork(root, instructions, stop)
	assert.Equal(t, 2, numSteps)
}

func TestTraverseNetwork_Simple(t *testing.T) {
	var instructions = day08.ParseInstructions("LLR")

	var network = []string{
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	var root = day08.BuildNetwork(network)
	var stop = func(n *day08.Node) bool { return n.Label == "ZZZ" }
	var numSteps = day08.TraverseNetwork(root, instructions, stop)
	assert.Equal(t, 6, numSteps)
}

func TestBuildGhostNetwork(t *testing.T) {
	var network = []string{
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}

	var startingNodes = day08.BuildGhostNetwork(network)
	assert.Equal(t, 2, len(startingNodes))
	assert.Equal(t, "11A", startingNodes[0].Label)
	assert.Equal(t, "22A", startingNodes[1].Label)
}

func TestGreatestCommonDivisor(t *testing.T) {
	assert.Equal(t, 4, day08.GreatestCommonDivisor(8, 12))
	assert.Equal(t, 4, day08.GreatestCommonDivisor(12, 8))
}

func TestLeastCommonMultiple(t *testing.T) {
	assert.Equal(t, 20, day08.LeastCommonMultiple(4, 5))
	assert.Equal(t, 30, day08.LeastCommonMultiple(6, 15))
}
