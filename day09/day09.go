package day09

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
	var total = 0
	var lines = strings.Split(input, "\n")

	for _, line := range lines {
		var history = ParseHistory(line)
		var nextDigit = GenerateNextDigit(history)

		total += nextDigit
	}

	answer = total
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()
	var total = 0
	var lines = strings.Split(input, "\n")

	for _, line := range lines {
		var history = ParseHistory(line)
		var previousDigit = GeneratePreviousDigit(history)

		total += previousDigit
	}

	answer = total
	return answer, time.Since(now)
}

func ParseHistory(line string) (history []int) {
	// 0 3 6 9 12 15
	for _, field := range strings.Fields(line) {
		var num, _ = strconv.Atoi(field)
		history = append(history, num)
	}

	return
}

func GenerateNextDigit(history []int) (nextDigit int) {
	/* 0   3   6   9  12  15
		 3   3   3   3   3
	  	   0   0   0   0
	*/
	nextDigit = history[len(history)-1]

	for {
		var sequence = make([]int, 0, len(history)-1)
		var allZeroes = true

		for i := 1; i < len(history); i++ {
			var val = history[i] - history[i-1]
			sequence = append(sequence, val)

			if val != 0 {
				allZeroes = false
			}
		}

		nextDigit += sequence[len(sequence)-1]
		history = sequence

		if allZeroes {
			return
		}
	}
}

func GeneratePreviousDigit(history []int) (previousDigit int) {
	var startDigits = []int{}
	startDigits = append(startDigits, history[0])

	for {
		var sequence = make([]int, 0, len(history)-1)
		var allZeroes = true

		for i := 1; i < len(history); i++ {
			var val = history[i] - history[i-1]
			sequence = append(sequence, val)

			if val != 0 {
				allZeroes = false
			}
		}

		startDigits = append(startDigits, sequence[0])
		history = sequence

		if allZeroes {
			break
		}
	}

	for i := len(startDigits) - 1; i >= 0; i-- {
		previousDigit = startDigits[i] - previousDigit
	}

	return
}
