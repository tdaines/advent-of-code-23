package day11_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day11"
)

func TestExpandUniverse(t *testing.T) {
	var universe = [][]byte{
		[]byte("...#......"),
		[]byte(".......#.."),
		[]byte("#........."),
		[]byte(".........."),
		[]byte("......#..."),
		[]byte(".#........"),
		[]byte(".........#"),
		[]byte(".........."),
		[]byte(".......#.."),
		[]byte("#...#....."),
	}

	var emptyRows, emptyCols = day11.ExpandUniverse(universe)
	assert.Equal(t, 2, len(emptyRows))
	assert.Equal(t, 3, emptyRows[0])
	assert.Equal(t, 7, emptyRows[1])

	assert.Equal(t, 3, len(emptyCols))
	assert.Equal(t, 2, emptyCols[0])
	assert.Equal(t, 5, emptyCols[1])
	assert.Equal(t, 8, emptyCols[2])

	// assert.Equal(t, 10, len(universe))
	// assert.Equal(t, []byte("..!#.!..!."), universe[0])
	// assert.Equal(t, []byte("..!..!.#!."), universe[1])
	// assert.Equal(t, []byte("#.!..!..!."), universe[2])
	// assert.Equal(t, []byte("!!!!!!!!!!"), universe[3])
	// assert.Equal(t, []byte("..!..!#.!."), universe[4])
	// assert.Equal(t, []byte(".#!..!..!."), universe[5])
	// assert.Equal(t, []byte("..!..!..!#"), universe[6])
	// assert.Equal(t, []byte("!!!!!!!!!!"), universe[7])
	// assert.Equal(t, []byte("..!..!.#!."), universe[8])
	// assert.Equal(t, []byte("#.!.#!..!."), universe[9])
}

func TestFindGalaxies(t *testing.T) {
	var universe = [][]byte{
		[]byte("....#........"),
		[]byte(".........#..."),
		[]byte("#............"),
		[]byte("............."),
		[]byte("............."),
		[]byte("........#...."),
		[]byte(".#..........."),
		[]byte("............#"),
		[]byte("............."),
		[]byte("............."),
		[]byte(".........#..."),
		[]byte("#....#......."),
	}

	var galaxies = day11.FindGalaxies(universe)
	assert.Equal(t, 9, len(galaxies))
	assert.Equal(t, 0, galaxies[0].Row)
	assert.Equal(t, 4, galaxies[0].Col)

	assert.Equal(t, 2, galaxies[2].Row)
	assert.Equal(t, 0, galaxies[2].Col)

	assert.Equal(t, 6, galaxies[4].Row)
	assert.Equal(t, 1, galaxies[4].Col)

	assert.Equal(t, 10, galaxies[6].Row)
	assert.Equal(t, 9, galaxies[6].Col)

	assert.Equal(t, 11, galaxies[8].Row)
	assert.Equal(t, 5, galaxies[8].Col)
}

func TestCountStepsBetween(t *testing.T) {
	var universe = [][]byte{
		[]byte("...#......"),
		[]byte(".......#.."),
		[]byte("#........."),
		[]byte(".........."),
		[]byte("......#..."),
		[]byte(".#........"),
		[]byte(".........#"),
		[]byte(".........."),
		[]byte(".......#.."),
		[]byte("#...#....."),
	}

	var emptyRows, emptyCols = day11.ExpandUniverse(universe)
	var galaxies = day11.FindGalaxies(universe)

	assert.Equal(t, 9, day11.CountStepsBetween(universe, galaxies[4], galaxies[8], emptyRows, emptyCols, 2))
	assert.Equal(t, 15, day11.CountStepsBetween(universe, galaxies[0], galaxies[6], emptyRows, emptyCols, 2))
	assert.Equal(t, 17, day11.CountStepsBetween(universe, galaxies[2], galaxies[5], emptyRows, emptyCols, 2))
	assert.Equal(t, 5, day11.CountStepsBetween(universe, galaxies[7], galaxies[8], emptyRows, emptyCols, 2))
}

func TestCountStepsBetweenAllPairs_2x(t *testing.T) {
	var universe = [][]byte{
		[]byte("...#......"),
		[]byte(".......#.."),
		[]byte("#........."),
		[]byte(".........."),
		[]byte("......#..."),
		[]byte(".#........"),
		[]byte(".........#"),
		[]byte(".........."),
		[]byte(".......#.."),
		[]byte("#...#....."),
	}

	var emptyRows, emptyCols = day11.ExpandUniverse(universe)
	var galaxies = day11.FindGalaxies(universe)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += day11.CountStepsBetween(universe, galaxies[i], galaxies[j], emptyRows, emptyCols, 2)
		}
	}

	assert.Equal(t, 374, steps)
}

func TestCountStepsBetweenAllPairs_10x(t *testing.T) {
	var universe = [][]byte{
		[]byte("...#......"),
		[]byte(".......#.."),
		[]byte("#........."),
		[]byte(".........."),
		[]byte("......#..."),
		[]byte(".#........"),
		[]byte(".........#"),
		[]byte(".........."),
		[]byte(".......#.."),
		[]byte("#...#....."),
	}

	var emptyRows, emptyCols = day11.ExpandUniverse(universe)
	var galaxies = day11.FindGalaxies(universe)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += day11.CountStepsBetween(universe, galaxies[i], galaxies[j], emptyRows, emptyCols, 10)
		}
	}

	assert.Equal(t, 1030, steps)
}

func TestCountStepsBetweenAllPairs_100x(t *testing.T) {
	var universe = [][]byte{
		[]byte("...#......"),
		[]byte(".......#.."),
		[]byte("#........."),
		[]byte(".........."),
		[]byte("......#..."),
		[]byte(".#........"),
		[]byte(".........#"),
		[]byte(".........."),
		[]byte(".......#.."),
		[]byte("#...#....."),
	}

	var emptyRows, emptyCols = day11.ExpandUniverse(universe)
	var galaxies = day11.FindGalaxies(universe)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += day11.CountStepsBetween(universe, galaxies[i], galaxies[j], emptyRows, emptyCols, 100)
		}
	}

	assert.Equal(t, 8410, steps)
}