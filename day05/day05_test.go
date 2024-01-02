package day05_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day05"
)

func TestParseRange(t *testing.T) {
	var r = day05.ParseRange("50 98 2")
	assert.Equal(t, 98, r.SourceStart)
	assert.Equal(t, 99, r.SourceEnd)
	assert.Equal(t, 50, r.DestStart)
	assert.Equal(t, 51, r.DestEnd)

	r = day05.ParseRange("52 50 48")
	assert.Equal(t, 50, r.SourceStart)
	assert.Equal(t, 97, r.SourceEnd)
	assert.Equal(t, 52, r.DestStart)
	assert.Equal(t, 99, r.DestEnd)
}

func TestParseMap(t *testing.T) {
	lines := []string{
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
	}

	var m = day05.ParseMap(lines)
	assert.Equal(t, "seed-to-soil map", m.Name)
	assert.Equal(t, 2, len(m.Ranges))
	assert.Equal(t, 98, m.Ranges[0].SourceStart)
	assert.Equal(t, 99, m.Ranges[0].SourceEnd)
	assert.Equal(t, 50, m.Ranges[0].DestStart)
	assert.Equal(t, 51, m.Ranges[0].DestEnd)
	assert.Equal(t, 50, m.Ranges[1].SourceStart)
	assert.Equal(t, 97, m.Ranges[1].SourceEnd)
	assert.Equal(t, 52, m.Ranges[1].DestStart)
	assert.Equal(t, 99, m.Ranges[1].DestEnd)
}

func TestMapSourceToDestination(t *testing.T) {
	lines := []string{
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
	}

	var m = day05.ParseMap(lines)

	assert.Equal(t, 81, day05.MapSourceToDestination(m, 79))
	assert.Equal(t, 14, day05.MapSourceToDestination(m, 14))
	assert.Equal(t, 57, day05.MapSourceToDestination(m, 55))
	assert.Equal(t, 13, day05.MapSourceToDestination(m, 13))
}

func TestParseParseSeeds(t *testing.T) {
	line := "seeds: 79 14 55 13"

	var seeds = day05.ParseSeeds(line)
	assert.Equal(t, 4, len(seeds))
	assert.Equal(t, 79, seeds[0])
	assert.Equal(t, 14, seeds[1])
	assert.Equal(t, 55, seeds[2])
	assert.Equal(t, 13, seeds[3])
}

func TestParseMaps(t *testing.T) {
	lines := []string{
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
	}

	var maps = day05.ParseMaps(lines)
	assert.Equal(t, 4, len(maps))

	var m = maps[0]
	assert.Equal(t, "seed-to-soil map", m.Name)
	assert.Equal(t, 2, len(m.Ranges))
	assert.Equal(t, 98, m.Ranges[0].SourceStart)
	assert.Equal(t, 99, m.Ranges[0].SourceEnd)
	assert.Equal(t, 50, m.Ranges[0].DestStart)
	assert.Equal(t, 51, m.Ranges[0].DestEnd)

	assert.Equal(t, 50, m.Ranges[1].SourceStart)
	assert.Equal(t, 97, m.Ranges[1].SourceEnd)
	assert.Equal(t, 52, m.Ranges[1].DestStart)
	assert.Equal(t, 99, m.Ranges[1].DestEnd)

	m = maps[1]
	assert.Equal(t, "soil-to-fertilizer map", m.Name)
	assert.Equal(t, 3, len(m.Ranges))
	assert.Equal(t, 15, m.Ranges[0].SourceStart)
	assert.Equal(t, 51, m.Ranges[0].SourceEnd)
	assert.Equal(t, 0, m.Ranges[0].DestStart)
	assert.Equal(t, 36, m.Ranges[0].DestEnd)

	assert.Equal(t, 52, m.Ranges[1].SourceStart)
	assert.Equal(t, 53, m.Ranges[1].SourceEnd)
	assert.Equal(t, 37, m.Ranges[1].DestStart)
	assert.Equal(t, 38, m.Ranges[1].DestEnd)

	assert.Equal(t, 0, m.Ranges[2].SourceStart)
	assert.Equal(t, 14, m.Ranges[2].SourceEnd)
	assert.Equal(t, 39, m.Ranges[2].DestStart)
	assert.Equal(t, 53, m.Ranges[2].DestEnd)

	m = maps[2]
	assert.Equal(t, "fertilizer-to-water map", m.Name)
	assert.Equal(t, 4, len(m.Ranges))
	assert.Equal(t, 53, m.Ranges[0].SourceStart)
	assert.Equal(t, 60, m.Ranges[0].SourceEnd)
	assert.Equal(t, 49, m.Ranges[0].DestStart)
	assert.Equal(t, 56, m.Ranges[0].DestEnd)

	assert.Equal(t, 11, m.Ranges[1].SourceStart)
	assert.Equal(t, 52, m.Ranges[1].SourceEnd)
	assert.Equal(t, 0, m.Ranges[1].DestStart)
	assert.Equal(t, 41, m.Ranges[1].DestEnd)

	assert.Equal(t, 0, m.Ranges[2].SourceStart)
	assert.Equal(t, 6, m.Ranges[2].SourceEnd)
	assert.Equal(t, 42, m.Ranges[2].DestStart)
	assert.Equal(t, 48, m.Ranges[2].DestEnd)

	assert.Equal(t, 7, m.Ranges[3].SourceStart)
	assert.Equal(t, 10, m.Ranges[3].SourceEnd)
	assert.Equal(t, 57, m.Ranges[3].DestStart)
	assert.Equal(t, 60, m.Ranges[3].DestEnd)

	// "water-to-light map:",
	// 	"88 18 7",
	// 	"18 25 70",
	m = maps[3]
	assert.Equal(t, "water-to-light map", m.Name)
	assert.Equal(t, 2, len(m.Ranges))
	assert.Equal(t, 18, m.Ranges[0].SourceStart)
	assert.Equal(t, 24, m.Ranges[0].SourceEnd)
	assert.Equal(t, 88, m.Ranges[0].DestStart)
	assert.Equal(t, 94, m.Ranges[0].DestEnd)

	assert.Equal(t, 25, m.Ranges[1].SourceStart)
	assert.Equal(t, 94, m.Ranges[1].SourceEnd)
	assert.Equal(t, 18, m.Ranges[1].DestStart)
	assert.Equal(t, 87, m.Ranges[1].DestEnd)
}
