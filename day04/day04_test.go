package day04_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day04"
)

func TestParseNumbers(t *testing.T) {
	numbers := day04.ParseNumbers("41 48 83 86 17")
	assert.Equal(t, 5, numbers.Cardinality())
	assert.True(t, numbers.Contains(41))
	assert.True(t, numbers.Contains(48))
	assert.True(t, numbers.Contains(83))
	assert.True(t, numbers.Contains(86))
	assert.True(t, numbers.Contains(17))

	numbers = day04.ParseNumbers(" 1 21 53 59 44")
	assert.Equal(t, 5, numbers.Cardinality())
	assert.True(t, numbers.Contains(1))
	assert.True(t, numbers.Contains(21))
	assert.True(t, numbers.Contains(53))
	assert.True(t, numbers.Contains(59))
	assert.True(t, numbers.Contains(44))

	numbers = day04.ParseNumbers("83 86  6 31 17  9 48 53")
	assert.Equal(t, 8, numbers.Cardinality())
	assert.True(t, numbers.Contains(83))
	assert.True(t, numbers.Contains(86))
	assert.True(t, numbers.Contains(6))
	assert.True(t, numbers.Contains(31))
	assert.True(t, numbers.Contains(17))
	assert.True(t, numbers.Contains(9))
	assert.True(t, numbers.Contains(48))
	assert.True(t, numbers.Contains(53))
}

func TestParseCard(t *testing.T) {
	assert.Equal(t, 8, day04.ParseCardPoints("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"))
	assert.Equal(t, 2, day04.ParseCardPoints("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"))
	assert.Equal(t, 2, day04.ParseCardPoints("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"))
	assert.Equal(t, 1, day04.ParseCardPoints("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"))
	assert.Equal(t, 0, day04.ParseCardPoints("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"))
	assert.Equal(t, 0, day04.ParseCardPoints("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"))
}

func TestCountCards(t *testing.T) {
	lines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	assert.Equal(t, 30, day04.CountCards(lines))
}
