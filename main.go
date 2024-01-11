package main

import (
	"time"

	"github.com/cheynewallace/tabby"
	"github.com/tdaines/advent-of-code-23/day01"
	"github.com/tdaines/advent-of-code-23/day02"
	"github.com/tdaines/advent-of-code-23/day03"
	"github.com/tdaines/advent-of-code-23/day04"
	"github.com/tdaines/advent-of-code-23/day05"
	"github.com/tdaines/advent-of-code-23/day06"
	"github.com/tdaines/advent-of-code-23/day07"
	"github.com/tdaines/advent-of-code-23/day08"
	"github.com/tdaines/advent-of-code-23/day09"
	"github.com/tdaines/advent-of-code-23/day10"
	"github.com/tdaines/advent-of-code-23/day11"
)

func main() {
	var dayBreak = "---"
	var partBreak = "----"
	var answerBreak = "-------------"
	var elapsedBreak = "-----------"

	t := tabby.New()
	t.AddHeader("DAY", "PART", "ANSWER", "ELAPSED")
	var total time.Duration = 0

	var answer, elapsed = day01.Part1()
	total += elapsed
	t.AddLine("1", "1", answer, elapsed)

	answer, elapsed = day01.Part2()
	total += elapsed
	t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day02.Part1()
	total += elapsed
	t.AddLine("2", "1", answer, elapsed)

	answer, elapsed = day02.Part2()
	total += elapsed
	t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day03.Part1()
	total += elapsed
	t.AddLine("3", "1", answer, elapsed)

	answer, elapsed = day03.Part2()
	total += elapsed
	t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day04.Part1()
	total += elapsed
	t.AddLine("4", "1", answer, elapsed)

	answer, elapsed = day04.Part2()
	total += elapsed
	t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day05.Part1()
	total += elapsed
	t.AddLine("5", "1", answer, elapsed)

	// answer, elapsed = day05.Part2()
	t.AddLine("", "2", "DNF", "DNF")

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day06.Part1()
	total += elapsed
	t.AddLine("6", "1", answer, elapsed)

	answer, elapsed = day06.Part2()
	total += elapsed
	t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day07.Part1()
	total += elapsed
	t.AddLine("7", "1", answer, elapsed)

	answer, elapsed = day07.Part2()
	total += elapsed
	t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day08.Part1()
	total += elapsed
	t.AddLine("8", "1", answer, elapsed)

	answer, elapsed = day08.Part2()
	total += elapsed
	t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day09.Part1()
	total += elapsed
	t.AddLine("9", "1", answer, elapsed)

	answer, elapsed = day09.Part2()
	total += elapsed
	t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day10.Part1()
	total += elapsed
	t.AddLine("10", "1", answer, elapsed)
	t.AddLine("", "2", "DNF", "DNF")

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	answer, elapsed = day11.Part1()
	total += elapsed
	t.AddLine("11", "1", answer, elapsed)

	answer, elapsed = day11.Part2()
	total += elapsed
	t.AddLine("", "2", answer, elapsed)

	t.AddLine(dayBreak, partBreak, answerBreak, elapsedBreak)

	t.AddLine("", "", "TOTAL", total)
	t.Print()
}
