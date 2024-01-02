package main

import (
	"github.com/cheynewallace/tabby"
	"github.com/tdaines/advent-of-code-23/day01"
	"github.com/tdaines/advent-of-code-23/day02"
	"github.com/tdaines/advent-of-code-23/day03"
	"github.com/tdaines/advent-of-code-23/day04"
	"github.com/tdaines/advent-of-code-23/day05"
)

func main() {
	t := tabby.New()
	t.AddHeader("DAY", "PART", "ANSWER", "ELAPSED")

	var answer, elapsed = day01.Part1()
	t.AddLine("1", "1", answer, elapsed)

	answer, elapsed = day01.Part2()
	t.AddLine("1", "2", answer, elapsed)

	answer, elapsed = day02.Part1()
	t.AddLine("2", "1", answer, elapsed)

	answer, elapsed = day02.Part2()
	t.AddLine("2", "2", answer, elapsed)

	answer, elapsed = day03.Part1()
	t.AddLine("3", "1", answer, elapsed)

	answer, elapsed = day03.Part2()
	t.AddLine("3", "2", answer, elapsed)

	answer, elapsed = day04.Part1()
	t.AddLine("4", "1", answer, elapsed)

	answer, elapsed = day04.Part2()
	t.AddLine("4", "2", answer, elapsed)

	answer, elapsed = day05.Part1()
	t.AddLine("5", "1", answer, elapsed)

	// answer, elapsed = day05.Part2()
	// t.AddLine("5", "2", answer, elapsed)

	t.Print()
}
