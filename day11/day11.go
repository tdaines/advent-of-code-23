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
	var galaxies = FindGalaxies(universe, emptyRows, emptyCols, 2)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += CountStepsBetween(galaxies[i], galaxies[j])
		}
	}

	answer = steps
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()

	var universe = bytes.Split(input, []byte{'\n'})

	var emptyRows, emptyCols = ExpandUniverse(universe)
	var galaxies = FindGalaxies(universe, emptyRows, emptyCols, 1_000_000)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += CountStepsBetween(galaxies[i], galaxies[j])
		}
	}

	answer = steps
	return answer, time.Since(now)
}

type Universe [][]byte

const (
	GALAXY byte = '#'
	SPACE  byte = '.'
)

type Galaxy struct {
	Row int
	Col int
}

type Nothing struct{}

var Void = Nothing{}

func ExpandUniverse(universe Universe) (emptyRows, emptyCols map[int]Nothing) {
	emptyRows = map[int]Nothing{}
	emptyCols = map[int]Nothing{}

	for i, row := range universe {
		if isEmptySpace(row) {
			emptyRows[i] = Void
		}
	}

	for col := 0; col < len(universe[0]); col++ {
		if isEmptySpaceCol(universe, col) {
			emptyCols[col] = Void
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

func FindGalaxies(universe Universe, emptyRows, emptyCols map[int]Nothing, penaltyMultiplier int) (galaxies []Galaxy) {

	var rowActual = 0
	for row := 0; row < len(universe); row++ {
		if _, exists := emptyRows[row]; exists {
			rowActual += (penaltyMultiplier - 1)
		}

		var colActual = 0
		for col := 0; col < len(universe[row]); col++ {
			if _, exists := emptyCols[col]; exists {
				colActual += (penaltyMultiplier - 1)
			}

			if universe[row][col] == GALAXY {
				galaxies = append(galaxies, Galaxy{Row: rowActual, Col: colActual})
			}

			colActual++
		}
		rowActual++
	}

	return
}

func CountStepsBetween(first, second Galaxy) (steps int) {

	var rowDiff = abs(first.Row - second.Row)
	var colDiff = abs(first.Col - second.Col)

	return rowDiff + colDiff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
