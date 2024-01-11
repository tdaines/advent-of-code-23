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

	ExpandUniverse(universe)
	var galaxies = FindGalaxies(universe)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += CountStepsBetween(universe, galaxies[i], galaxies[j], 2)
		}
	}

	answer = steps
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()

	var universe = bytes.Split(input, []byte{'\n'})

	ExpandUniverse(universe)
	var galaxies = FindGalaxies(universe)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += CountStepsBetween(universe, galaxies[i], galaxies[j], 1_000_000)
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

func ExpandUniverse(universe Universe) {
	for _, row := range universe {
		if isEmptySpace(row) {
			for col := 0; col < len(row); col++ {
				row[col] = PENALTY
			}
		}
	}

	for col := 0; col < len(universe[0]); col++ {
		if isEmptySpaceCol(universe, col) {
			for row := 0; row < len(universe); row++ {
				universe[row][col] = PENALTY
			}
		}
	}
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

func CountStepsBetween(universe Universe, first, second Galaxy, penaltyMultiplier int) (steps int) {
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

	for row := rowStart; row < rowEnd; row++ {
		if universe[row][0] == PENALTY {
			steps += penaltyMultiplier
		} else {
			steps++
		}
	}

	for col := colStart; col < colEnd; col++ {
		var row = universe[0]
		if row[col] == PENALTY {
			steps += penaltyMultiplier
		} else {
			steps++
		}
	}

	return
}
