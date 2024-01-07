package day10

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day10/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	var scanner = bufio.NewScanner(input)
	var lines = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var steps = FindMaxStepsFromStart(lines)

	answer = steps
	return answer, time.Since(now)
}

// func Part2() (answer int, elapsed time.Duration) {
// 	var now = time.Now()
// 	input, err := os.Open("./day10/input.txt")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer input.Close()

// 	scanner := bufio.NewScanner(input)
// 	var total = 0

// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		var history = ParseHistory(line)
// 		var previousDigit = GeneratePreviousDigit(history)

// 		total += previousDigit
// 	}

// 	answer = total
// 	return answer, time.Since(now)
// }

type Position struct {
	Row int
	Col int
}

func NewPosition(row, col int) Position {
	return Position{
		Row: row,
		Col: col,
	}
}

func (p Position) IsNorthOf(test Position) bool {
	return p.Col == test.Col && p.Row == test.Row-1
}

func (p Position) IsSouthOf(test Position) bool {
	return p.Col == test.Col && p.Row == test.Row+1
}

func (p Position) IsEastOf(test Position) bool {
	return p.Row == test.Row && p.Col == test.Col+1
}

func (p Position) IsWestOf(test Position) bool {
	return p.Row == test.Row && p.Col == test.Col-1
}

type Direction int

const (
	North Direction = 0
	South Direction = 1
	East  Direction = 2
	West  Direction = 3
)

func (p Position) Move(direction Direction) Position {
	switch direction {
	case North:
		return Position{Row: p.Row - 1, Col: p.Col}
	case South:
		return Position{Row: p.Row + 1, Col: p.Col}
	case East:
		return Position{Row: p.Row, Col: p.Col + 1}
	case West:
		return Position{Row: p.Row, Col: p.Col - 1}
	}

	return Position{Row: -1, Col: -1}
}

func GetNextPosition(previous, current Position, currentPipe rune) (next Position) {
	/*
		| is a vertical pipe connecting north and south.
		- is a horizontal pipe connecting east and west.
		L is a 90-degree bend connecting north and east.
		J is a 90-degree bend connecting north and west.
		7 is a 90-degree bend connecting south and west.
		F is a 90-degree bend connecting south and east.
	*/
	switch currentPipe {
	case '|':
		if previous.IsNorthOf(current) {
			return current.Move(South)
		} else {
			return current.Move(North)
		}
	case '-':
		if previous.IsWestOf(current) {
			return current.Move(East)
		} else {
			return current.Move(West)
		}
	case 'L':
		if previous.IsNorthOf(current) {
			return current.Move(East)
		} else {
			return current.Move(North)
		}
	case 'J':
		if previous.IsNorthOf(current) {
			return current.Move(West)
		} else {
			return current.Move(North)
		}
	case '7':
		if previous.IsWestOf(current) {
			return current.Move(South)
		} else {
			return current.Move(West)
		}
	case 'F':
		if previous.IsEastOf(current) {
			return current.Move(South)
		} else {
			return current.Move(East)
		}
	}

	return Position{Row: -1, Col: -1}
}

func FindStartPosition(lines []string) Position {
	for row := 0; row < len(lines); row++ {
		var line = lines[row]
		for col := 0; col < len(line); col++ {
			var pipe = line[col]
			if pipe == 'S' {
				return Position{Row: row, Col: col}
			}
		}
	}

	return Position{Row: -1, Col: -1}
}

func (p Position) ConnectsTo(test Position, testPipe rune) bool {
	/*
		| is a vertical pipe connecting north and south.
		- is a horizontal pipe connecting east and west.
		L is a 90-degree bend connecting north and east.
		J is a 90-degree bend connecting north and west.
		7 is a 90-degree bend connecting south and west.
		F is a 90-degree bend connecting south and east.
	*/

	switch testPipe {
	case '|':
		if p.IsNorthOf(test) || p.IsSouthOf(test) {
			return true
		}
	case '-':
		if p.IsEastOf(test) || p.IsWestOf(test) {
			return true
		}
	case 'L':
		if p.IsNorthOf(test) || p.IsEastOf(test) {
			return true
		}
	case 'J':
		if p.IsNorthOf(test) || p.IsWestOf(test) {
			return true
		}
	case '7':
		if p.IsSouthOf(test) || p.IsWestOf(test) {
			return true
		}
	case 'F':
		if p.IsSouthOf(test) || p.IsEastOf(test) {
			return true
		}
	}

	return false
}

func FindStartPositionConnections(start Position, lines []string) (first, second Position) {
	var found = []Position{}

	var diffs = []Position{
		{Row: -1, Col: 0}, // North
		{Row: 1, Col: 0},  // South
		{Row: 0, Col: 1},  // East
		{Row: 0, Col: -1}, // West
	}

	for _, diff := range diffs {
		var row = start.Row + diff.Row
		var col = start.Col + diff.Col

		if row >= 0 && row < len(lines)-1 {
			if col >= 0 && col < len(lines[row])-1 {
				var test = Position{Row: row, Col: col}
				var testPipe = rune(lines[row][col])

				if start.ConnectsTo(test, testPipe) {
					found = append(found, test)
					if len(found) == 2 {
						break
					}
				}
			}
		}
	}

	first = found[0]
	second = found[1]
	return
}

func FindMaxStepsFromStart(lines []string) int {
	var start = FindStartPosition(lines)
	var left, right = FindStartPositionConnections(start, lines)
	var previousLeft, previousRight = start, start

	var steps = 1

	for {
		// Stop when left and right paths collide
		if left.Row == right.Row && left.Col == right.Col {
			return steps
		}

		steps++

		// Advance left path
		var current = left
		var currentPipe = rune(lines[current.Row][current.Col])
		var next = GetNextPosition(previousLeft, current, currentPipe)
		previousLeft = left
		left = next

		// Advance right path
		current = right
		currentPipe = rune(lines[current.Row][current.Col])
		next = GetNextPosition(previousRight, current, currentPipe)
		previousRight = right
		right = next
	}
}
