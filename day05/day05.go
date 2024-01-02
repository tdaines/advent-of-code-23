package day05

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day05/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var seeds = ParseSeeds(lines[0])
	lines = lines[1:]
	var minLocation = math.MaxInt64

	var maps = ParseMaps(lines)
	for _, seed := range seeds {
		var location = seed
		for _, m := range maps {
			location = MapSourceToDestination(m, location)
		}

		minLocation = min(minLocation, location)
	}

	answer = minLocation
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day05/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var seeds = ParseSeedRanges(lines[0])
	lines = lines[1:]
	var minLocation = math.MaxInt64

	var maps = ParseMaps(lines)
	for _, seed := range seeds {
		var location = seed
		for _, m := range maps {
			location = MapSourceToDestination(m, location)
		}

		minLocation = min(minLocation, location)
	}

	answer = minLocation
	return answer, time.Since(now)
}

func ParseSeeds(line string) []int {
	// seeds: 79 14 55 13
	var parts = strings.Split(line, ":")
	var seeds = []int{}

	parts = strings.Split(parts[1], " ")

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			seeds = append(seeds, num)
		}
	}

	return seeds
}

func ParseSeedRanges(line string) []int {
	// seeds: 79 14 55 13
	var parts = strings.Split(line, ": ")
	var seeds = []int{}

	parts = strings.Split(parts[1], " ")

	for len(parts) > 0 {
		var start, _ = strconv.Atoi(parts[0])
		var length, _ = strconv.Atoi(parts[1])

		for i := 0; i < length; i++ {
			seeds = append(seeds, start+i)
		}

		parts = parts[2:]
	}

	return seeds
}

func ParseMaps(lines []string) []Map {
	var maps = []Map{}
	for len(lines) > 0 {
		var line = lines[0]
		if len(line) == 0 {
			lines = lines[1:]
			continue
		}

		var mapLines = []string{}
		for len(lines) > 0 && len(lines[0]) > 0 {
			line = lines[0]
			mapLines = append(mapLines, line)
			lines = lines[1:]
		}
		var m = ParseMap(mapLines)
		maps = append(maps, m)
	}

	return maps
}

type Range struct {
	SourceStart int
	SourceEnd   int
	DestStart   int
	DestEnd     int
}

type Map struct {
	Name   string
	Ranges []Range
}

func ParseMap(lines []string) Map {
	// seed-to-soil map:
	// 50 98 2
	// 52 50 48
	var m = Map{}
	m.Name = lines[0]
	m.Name = m.Name[:len(m.Name)-1]

	lines = lines[1:]
	for _, line := range lines {
		var r = ParseRange(line)
		m.Ranges = append(m.Ranges, r)
	}

	return m
}

func ParseRange(line string) Range {
	// 50 98 2
	var parts = strings.Split(line, " ")
	var r = Range{}

	r.DestStart, _ = strconv.Atoi(parts[0])
	r.SourceStart, _ = strconv.Atoi(parts[1])
	var length, _ = strconv.Atoi(parts[2])

	r.DestEnd = r.DestStart + (length - 1)
	r.SourceEnd = r.SourceStart + (length - 1)

	return r
}

func MapSourceToDestination(mapping Map, source int) int {
	for _, r := range mapping.Ranges {
		if source >= r.SourceStart && source <= r.SourceEnd {
			var offset = r.DestStart - r.SourceStart
			return source + offset
		}
	}

	return source
}
