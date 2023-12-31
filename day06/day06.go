package day06

import (
	_ "embed"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
}

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	var lines = strings.Split(input, "\n")

	var marginOfError = 1
	var races = ParseRaces(lines[0], lines[1])
	for _, race := range races {
		marginOfError *= CountWaysToWin(race)
	}

	answer = marginOfError
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()
	var lines = strings.Split(input, "\n")

	var race = ParseRace(lines[0], lines[1])
	answer = CountWaysToWin(race)
	return answer, time.Since(now)
}

type Race struct {
	Time     int
	Distance int
}

func ParseRaces(timeLine string, distanceLine string) []Race {
	// Time:      7  15   30
	// Distance:  9  40  200

	// split by whitespace and remove 'Time:'
	var times = strings.Fields(timeLine)[1:]
	var distances = strings.Fields(distanceLine)[1:]

	var races = []Race{}

	for i := 0; i < len(times); i++ {
		var race = Race{}
		race.Time, _ = strconv.Atoi(times[i])
		race.Distance, _ = strconv.Atoi(distances[i])
		races = append(races, race)
	}

	return races
}

func CountWaysToWin(race Race) int {
	var minDistance = race.Distance
	var startTime = 0

	var start = 0
	var end = race.Time

	for {
		var mid = start + (end-start)/2
		var speed = mid
		var distance = CalcDistance(speed, race.Time)
		if distance > minDistance {
			// Found a win, but is it the first?
			if CalcDistance(speed-1, race.Time) <= minDistance {
				startTime = speed
				break
			} else {
				// Not the first win
				end = mid
			}
		} else {
			start = mid
		}
	}

	// End time is the same distance from the race.Time as
	// start time is from zero
	var endTime = race.Time - startTime

	return endTime - startTime + 1
}

func CalcDistance(speed int, raceTime int) int {
	var remainingTime = raceTime - speed
	return speed * remainingTime
}

func ParseRace(timeLine string, distanceLine string) Race {
	// Time:      7  15   30
	// Distance:  9  40  200

	// split by whitespace and remove 'Time:'
	var times = strings.Fields(timeLine)[1:]
	var distances = strings.Fields(distanceLine)[1:]

	var timeString = strings.Join(times, "")
	var distanceString = strings.Join(distances, "")

	var race = Race{}
	race.Time, _ = strconv.Atoi(timeString)
	race.Distance, _ = strconv.Atoi(distanceString)

	return race
}
