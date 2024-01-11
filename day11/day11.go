package day11

import (
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var input []byte

func init() {
	input = bytes.TrimRight(input, "\n")
}

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	var universe = bytes.Split(input, []byte{'\n'})

	var emptyRows, emptyCols = ExpandUniverse(universe)
	var galaxies = FindGalaxies(universe)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += CountStepsBetween(universe, galaxies[i], galaxies[j], emptyRows, emptyCols, 2)
		}
	}

	answer = steps
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()

	var universe = bytes.Split(input, []byte{'\n'})

	var emptyRows, emptyCols = ExpandUniverse(universe)
	var galaxies = FindGalaxies(universe)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += CountStepsBetween(universe, galaxies[i], galaxies[j], emptyRows, emptyCols, 1_000_000)
		}
	}

	answer = steps
	return answer, time.Since(now)
}

type Universe [][]byte

const (
	GALAXY  byte = '#'
	SPACE   byte = '.'
	PENALTY byte = '!'
)

type Galaxy struct {
	Row int
	Col int
}

func ExpandUniverse(universe Universe) (emptyRows []int, emptyCols []int) {
	for i, row := range universe {
		if isEmptySpace(row) {
			emptyRows = append(emptyRows, i)
		}
	}

	for col := 0; col < len(universe[0]); col++ {
		if isEmptySpaceCol(universe, col) {
			emptyCols = append(emptyCols, col)
		}
	}

	return
}

func isEmptySpace(imageArea []byte) bool {
	return bytes.IndexByte(imageArea, GALAXY) == -1
}

func isEmptySpaceCol(universe Universe, col int) bool {
	for row := 0; row < len(universe); row++ {
		if universe[row][col] == GALAXY {
			return false
		}
	}

	return true
}

func FindGalaxies(universe Universe) (galaxies []Galaxy) {
	for row := 0; row < len(universe); row++ {
		for col := 0; col < len(universe[row]); col++ {
			if universe[row][col] == GALAXY {
				galaxies = append(galaxies, Galaxy{Row: row, Col: col})
			}
		}
	}

	return
}

func CountStepsBetween(universe Universe, first, second Galaxy, emptyRows, emptyCols []int, penaltyMultiplier int) (steps int) {
	var rowStart = first.Row
	var rowEnd = second.Row
	if rowStart > rowEnd {
		rowStart, rowEnd = rowEnd, rowStart
	}

	var colStart = first.Col
	var colEnd = second.Col
	if colStart > colEnd {
		colStart, colEnd = colEnd, colStart
	}

	// Count number of emptry rows between startRow and endRow
	var numEmptyRows = CountNumsWithinRange(emptyRows, rowStart, rowEnd)
	var rowDistance = (rowEnd - rowStart) - numEmptyRows + (numEmptyRows * penaltyMultiplier)

	// Count number of emptry cols between startCol and endCol
	var numEmptyCols = CountNumsWithinRange(emptyCols, colStart, colEnd)
	var colDistance = (colEnd - colStart) - numEmptyCols + (numEmptyCols * penaltyMultiplier)

	steps = rowDistance + colDistance

	return steps
}

func CountNumsWithinRange(nums []int, start, end int) (count int) {
	for _, num := range nums {
		if num >= end {
			return
		}

		if num > start && num < end {
			count++
		}
	}

	return
}
