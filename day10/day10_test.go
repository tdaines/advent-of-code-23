package day10_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day10"
)

func TestIsNorthOf(t *testing.T) {
	var pos = day10.Position{
		Row: 5,
		Col: 5,
	}

	assert.True(t, pos.IsNorthOf(day10.Position{Row: 6, Col: 5}))
	assert.False(t, pos.IsNorthOf(day10.Position{Row: 5, Col: 5}))
	assert.False(t, pos.IsNorthOf(day10.Position{Row: 4, Col: 5}))
}

func TestIsSouthOf(t *testing.T) {
	var pos = day10.Position{
		Row: 5,
		Col: 5,
	}

	assert.False(t, pos.IsSouthOf(day10.Position{Row: 6, Col: 5}))
	assert.False(t, pos.IsSouthOf(day10.Position{Row: 5, Col: 5}))
	assert.True(t, pos.IsSouthOf(day10.Position{Row: 4, Col: 5}))
}

func TestIsEastOf(t *testing.T) {
	var pos = day10.Position{
		Row: 5,
		Col: 5,
	}

	assert.False(t, pos.IsEastOf(day10.Position{Row: 5, Col: 6}))
	assert.False(t, pos.IsEastOf(day10.Position{Row: 5, Col: 5}))
	assert.True(t, pos.IsEastOf(day10.Position{Row: 5, Col: 4}))
}

func TestIsWestOf(t *testing.T) {
	var pos = day10.Position{
		Row: 5,
		Col: 5,
	}

	assert.True(t, pos.IsWestOf(day10.Position{Row: 5, Col: 6}))
	assert.False(t, pos.IsWestOf(day10.Position{Row: 5, Col: 5}))
	assert.False(t, pos.IsWestOf(day10.Position{Row: 5, Col: 4}))
}

func TestMove(t *testing.T) {
	var pos = day10.Position{
		Row: 5,
		Col: 5,
	}

	var moved = pos.Move(day10.North)
	assert.Equal(t, 4, moved.Row)
	assert.Equal(t, 5, moved.Col)

	moved = pos.Move(day10.South)
	assert.Equal(t, 6, moved.Row)
	assert.Equal(t, 5, moved.Col)

	moved = pos.Move(day10.East)
	assert.Equal(t, 5, moved.Row)
	assert.Equal(t, 6, moved.Col)

	moved = pos.Move(day10.West)
	assert.Equal(t, 5, moved.Row)
	assert.Equal(t, 4, moved.Col)
}

func TestGetNextPosition_NorthSouth(t *testing.T) {
	var pipe = '|'

	var current = day10.Position{Row: 5, Col: 5}
	var previous = day10.Position{Row: 4, Col: 5}

	// Moving South
	var new = day10.GetNextPosition(previous, current, pipe)
	assert.Equal(t, 6, new.Row)
	assert.Equal(t, 5, new.Col)

	previous = day10.Position{Row: 6, Col: 5}

	// Moving North
	new = day10.GetNextPosition(previous, current, pipe)
	assert.Equal(t, 4, new.Row)
	assert.Equal(t, 5, new.Col)
}

func TestFindStartPosition(t *testing.T) {
	var lines = []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}

	var start = day10.FindStartPosition(lines)
	assert.Equal(t, 1, start.Row)
	assert.Equal(t, 1, start.Col)

	lines = []string{
		"7-F7-",
		".FJ|7",
		"SJLL7",
		"|F--J",
		"LJ.LJ",
	}

	start = day10.FindStartPosition(lines)
	assert.Equal(t, 2, start.Row)
	assert.Equal(t, 0, start.Col)
}

func TestFindStartPositionConnections(t *testing.T) {
	var lines = []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}

	var start = day10.FindStartPosition(lines)
	var first, second = day10.FindStartPositionConnections(start, lines)
	assert.Equal(t, 2, first.Row)
	assert.Equal(t, 1, first.Col)
	assert.Equal(t, 1, second.Row)
	assert.Equal(t, 2, second.Col)

	lines = []string{
		"7-F7-",
		".FJ|7",
		"SJLL7",
		"|F--J",
		"LJ.LJ",
	}

	start = day10.FindStartPosition(lines)
	first, second = day10.FindStartPositionConnections(start, lines)
	assert.Equal(t, 3, first.Row)
	assert.Equal(t, 0, first.Col)
	assert.Equal(t, 2, second.Row)
	assert.Equal(t, 1, second.Col)
}

func TestFindMaxStepsFromStart(t *testing.T) {
	var lines = []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}

	var steps = day10.FindMaxStepsFromStart(lines)
	assert.Equal(t, 4, steps)

	lines = []string{
		"7-F7-",
		".FJ|7",
		"SJLL7",
		"|F--J",
		"LJ.LJ",
	}

	steps = day10.FindMaxStepsFromStart(lines)
	assert.Equal(t, 8, steps)
}