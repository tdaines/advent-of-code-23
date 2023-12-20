package day03_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day03"
)

func TestIsSymbol(t *testing.T) {
	schematic := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	assert.False(t, day03.IsSymbol(schematic, -1, 0))
	assert.False(t, day03.IsSymbol(schematic, 0, -1))
	assert.False(t, day03.IsSymbol(schematic, 10, 0))
	assert.False(t, day03.IsSymbol(schematic, 0, 10))
	assert.False(t, day03.IsSymbol(schematic, 0, 0))
	assert.False(t, day03.IsSymbol(schematic, 1, 0))
	assert.True(t, day03.IsSymbol(schematic, 1, 3))
	assert.False(t, day03.IsSymbol(schematic, 1, 9))
}

func TestParseNumber(t *testing.T) {
	schematic := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	num, isPartNumber, endCol := day03.ParseNumber(schematic, 0, 0)
	assert.Equal(t, 467, num)
	assert.True(t, isPartNumber)
	assert.Equal(t, 2, endCol)

	num, isPartNumber, endCol = day03.ParseNumber(schematic, 0, 5)
	assert.Equal(t, 114, num)
	assert.False(t, isPartNumber)
	assert.Equal(t, 7, endCol)

	num, isPartNumber, endCol = day03.ParseNumber(schematic, 1, 0)
	assert.Equal(t, 0, num)
	assert.False(t, isPartNumber)
	assert.Equal(t, 0, endCol)

	num, isPartNumber, endCol = day03.ParseNumber(schematic, 2, 2)
	assert.Equal(t, 35, num)
	assert.True(t, isPartNumber)
	assert.Equal(t, 3, endCol)

	num, isPartNumber, endCol = day03.ParseNumber(schematic, 4, 0)
	assert.Equal(t, 617, num)
	assert.True(t, isPartNumber)
	assert.Equal(t, 2, endCol)

	num, isPartNumber, endCol = day03.ParseNumber(schematic, 7, 3)
	assert.Equal(t, 0, num)
	assert.False(t, isPartNumber)
	assert.Equal(t, 3, endCol)
}
