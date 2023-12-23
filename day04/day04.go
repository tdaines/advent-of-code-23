package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	set "github.com/deckarep/golang-set/v2"
)

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day04/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var sum int = 0
	for scanner.Scan() {
		line := scanner.Text()

		cardValue := ParseCardPoints(line)
		sum += cardValue
	}

	answer = sum
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day04/input.txt")
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

	var sum = CountCards(lines)
	answer = sum
	return answer, time.Since(now)
}

func CountCards(lines []string) int {
	var cards = map[int]int{}
	var cardNumber = 0
	for _, line := range lines {
		cardNumber += 1
		cards[cardNumber] += 1
		var numCopies = cards[cardNumber]

		numMatches := ParseCardNumMatches(line)
		for i := 1; i <= numMatches; i++ {
			cards[cardNumber+i] += numCopies
		}
	}

	var numCards = 0
	for _, v := range cards {
		numCards += v
	}

	return numCards
}

func ParseCardNumMatches(line string) int {
	var parts = strings.Split(line, ":")
	parts = strings.Split(parts[1], "|")

	var winningNumbers = ParseNumbers(parts[0])
	var chosenNumbers = ParseNumbers(parts[1])

	var matches = winningNumbers.Intersect(chosenNumbers)
	return matches.Cardinality()
}

func ParseCardPoints(line string) int {
	var parts = strings.Split(line, ":")
	parts = strings.Split(parts[1], "|")

	var winningNumbers = ParseNumbers(parts[0])
	var chosenNumbers = ParseNumbers(parts[1])

	var matches = winningNumbers.Intersect(chosenNumbers)
	var numMatches = matches.Cardinality()

	var cardValue = 0
	if numMatches > 0 {
		cardValue = 1 << (numMatches - 1)
	}

	return cardValue
}

func ParseNumbers(line string) set.Set[int] {
	// "41 48 83 86 17"
	// " 1 21 53 59 44"
	var numbers = set.NewSet[int]()

	parts := strings.Split(line, " ")
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err == nil {
			numbers.Add(number)
		}
	}

	return numbers
}
