package day09_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day09"
)

func TestParseHistory(t *testing.T) {
	var history = day09.ParseHistory("0 3 6 9 12 15")
	assert.Equal(t, 6, len(history))
	assert.Equal(t, 0, history[0])
	assert.Equal(t, 3, history[1])
	assert.Equal(t, 6, history[2])
	assert.Equal(t, 9, history[3])
	assert.Equal(t, 12, history[4])
	assert.Equal(t, 15, history[5])

	history = day09.ParseHistory("5 3 1 -1 -3 -5")
	assert.Equal(t, 6, len(history))
	assert.Equal(t, 5, history[0])
	assert.Equal(t, 3, history[1])
	assert.Equal(t, 1, history[2])
	assert.Equal(t, -1, history[3])
	assert.Equal(t, -3, history[4])
	assert.Equal(t, -5, history[5])
}

func TestGenerateNextDigit(t *testing.T) {
	var history = day09.ParseHistory("0 3 6 9 12 15")
	assert.Equal(t, 18, day09.GenerateNextDigit(history))

	history = day09.ParseHistory("10 13 16 21 30 45")
	assert.Equal(t, 68, day09.GenerateNextDigit(history))
}

func TestGeneratePreviousDigit(t *testing.T) {
	var history = day09.ParseHistory("0 3 6 9 12 15")
	assert.Equal(t, -3, day09.GeneratePreviousDigit(history))

	history = day09.ParseHistory("1 3 6 10 15 21")
	assert.Equal(t, 0, day09.GeneratePreviousDigit(history))

	history = day09.ParseHistory("10 13 16 21 30 45")
	assert.Equal(t, 5, day09.GeneratePreviousDigit(history))
}
