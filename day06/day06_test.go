package day06_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day06"
)

func TestParseRaces(t *testing.T) {
	var time = "Time:      7  15   30"
	var distance = "Distance:  9  40  200"

	var races = day06.ParseRaces(time, distance)
	assert.Equal(t, 3, len(races))

	assert.Equal(t, 7, races[0].Time)
	assert.Equal(t, 9, races[0].Distance)
	assert.Equal(t, 15, races[1].Time)
	assert.Equal(t, 40, races[1].Distance)
	assert.Equal(t, 30, races[2].Time)
	assert.Equal(t, 200, races[2].Distance)
}

func TestCountWaysToWin(t *testing.T) {
	var race = day06.Race{
		Time:     7,
		Distance: 9,
	}

	assert.Equal(t, 4, day06.CountWaysToWin(race))

	race = day06.Race{
		Time:     15,
		Distance: 40,
	}

	assert.Equal(t, 8, day06.CountWaysToWin(race))

	race = day06.Race{
		Time:     30,
		Distance: 200,
	}

	assert.Equal(t, 9, day06.CountWaysToWin(race))

	race = day06.Race{
		Time:     71530,
		Distance: 940200,
	}

	assert.Equal(t, 71503, day06.CountWaysToWin(race))
}

func TestParseRace(t *testing.T) {
	var time = "Time:      7  15   30"
	var distance = "Distance:  9  40  200"

	var race = day06.ParseRace(time, distance)
	assert.Equal(t, 71530, race.Time)
	assert.Equal(t, 940200, race.Distance)
}
