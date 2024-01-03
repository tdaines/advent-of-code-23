package day06

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day06/input.txt")
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
	input, err := os.Open("./day06/input.txt")
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
	var count = 0

	for holdTime := 0; holdTime <= race.Time; holdTime++ {
		var speed = holdTime
		var remainingTime = race.Time - holdTime

		var distance = speed * remainingTime
		if distance > minDistance {
			count++
		}
	}

	return count
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
