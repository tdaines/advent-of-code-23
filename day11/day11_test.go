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
	_, exists := emptyRows[3]
	assert.True(t, exists)

	_, exists = emptyRows[7]
	assert.True(t, exists)

	assert.Equal(t, 3, len(emptyCols))
	_, exists = emptyCols[2]
	assert.True(t, exists)

	_, exists = emptyCols[5]
	assert.True(t, exists)

	_, exists = emptyCols[8]
	assert.True(t, exists)
}

func TestFindGalaxies(t *testing.T) {
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

	var emptyRows = map[int]day11.Nothing{
		3: day11.Void,
		7: day11.Void,
	}

	var emptyCols = map[int]day11.Nothing{
		2: day11.Void,
		5: day11.Void,
		8: day11.Void,
	}

	var galaxies = day11.FindGalaxies(universe, emptyRows, emptyCols, 2)
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
	var galaxies = day11.FindGalaxies(universe, emptyRows, emptyCols, 2)

	assert.Equal(t, 9, day11.CountStepsBetween(galaxies[4], galaxies[8]))
	assert.Equal(t, 15, day11.CountStepsBetween(galaxies[0], galaxies[6]))
	assert.Equal(t, 17, day11.CountStepsBetween(galaxies[2], galaxies[5]))
	assert.Equal(t, 5, day11.CountStepsBetween(galaxies[7], galaxies[8]))
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
	var galaxies = day11.FindGalaxies(universe, emptyRows, emptyCols, 2)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += day11.CountStepsBetween(galaxies[i], galaxies[j])
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
	var galaxies = day11.FindGalaxies(universe, emptyRows, emptyCols, 10)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += day11.CountStepsBetween(galaxies[i], galaxies[j])
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
	var galaxies = day11.FindGalaxies(universe, emptyRows, emptyCols, 100)

	var steps = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			steps += day11.CountStepsBetween(galaxies[i], galaxies[j])
		}
	}

	assert.Equal(t, 8410, steps)
}