package day03

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type PartNumber struct {
	Row    int
	Col    int
	Number int
}

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day03/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	schematic := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}

	var sum int = 0

	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[row]); col++ {
			number, isPartNumber, endCol := ParseNumber(schematic, row, col)

			if isPartNumber {
				sum += number
			}

			col = endCol
		}
	}

	answer = sum
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day03/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	schematic := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}

	var sum int = 0

	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[row]); col++ {
			// "*" == 42
			if schematic[row][col] == 42 {
				sum += ParseGearRatio(schematic, row, col)
			}
		}
	}

	answer = sum
	return answer, time.Since(now)
}

func ParseNumber(schematic []string, row int, col int) (number int, isPartNumber bool, endCol int) {
	isPartNumber = false
	endCol = col
	foundDigit := false

	// while we are looking at a digit
	for col < len(schematic[row]) && schematic[row][col] >= 48 && schematic[row][col] <= 57 {
		foundDigit = true
		number = (number * 10) + int(schematic[row][col]-48)

		// check neighbors
		isPartNumber = isPartNumber || IsSymbol(schematic, row, col-1)   // left
		isPartNumber = isPartNumber || IsSymbol(schematic, row, col+1)   // right
		isPartNumber = isPartNumber || IsSymbol(schematic, row-1, col)   // up
		isPartNumber = isPartNumber || IsSymbol(schematic, row+1, col)   // down
		isPartNumber = isPartNumber || IsSymbol(schematic, row-1, col-1) // up-left
		isPartNumber = isPartNumber || IsSymbol(schematic, row-1, col+1) // up-right
		isPartNumber = isPartNumber || IsSymbol(schematic, row+1, col-1) // down-left
		isPartNumber = isPartNumber || IsSymbol(schematic, row+1, col+1) // down-right

		col++
	}

	if foundDigit {
		endCol = col - 1
	}

	return
}

func IsSymbol(schematic []string, row int, col int) bool {
	if row < 0 {
		return false
	}

	if row >= len(schematic) {
		return false
	}

	if col < 0 {
		return false
	}

	if col >= len(schematic[row]) {
		return false
	}

	// "." == 46
	if schematic[row][col] == 46 {
		return false
	}

	if schematic[row][col] >= 48 && schematic[row][col] <= 57 {
		return false
	}

	return true
}

type Pair struct {
	Row int
	Col int
}

func ParseGearRatio(schematic []string, row int, col int) int {
	// schematic[row][col] points to a '*'

	neighbors := []Pair{
		{Row: 0, Col: -1},  // left
		{Row: 0, Col: 1},   // right
		{Row: -1, Col: 0},  // up
		{Row: 1, Col: 0},   // down
		{Row: -1, Col: -1}, // up-left
		{Row: -1, Col: 1},  // up-right
		{Row: 1, Col: -1},  // down-left
		{Row: 1, Col: 1},   // down-right
	}

	partNumbers := []PartNumber{}

	for _, neighbor := range neighbors {
		partNumber := ParsePartNumber(schematic, row+neighbor.Row, col+neighbor.Col)
		if partNumber.Number > 0 {
			partNumbers = AddPartNumber(partNumbers, partNumber)
		}
	}

	if len(partNumbers) == 2 {
		// found valid gear ratio
		return partNumbers[0].Number * partNumbers[1].Number
	}

	return 0
}

func AddPartNumber(partNumbers []PartNumber, partNumber PartNumber) []PartNumber {
	for _, pn := range partNumbers {
		if pn.Row == partNumber.Row && pn.Col == partNumber.Col {
			// don't change the slice
			return partNumbers
		}
	}

	return append(partNumbers, partNumber)
}

func ParsePartNumber(schematic []string, row int, col int) (partNumber PartNumber) {
	if row < 0 {
		return
	}

	if row >= len(schematic) {
		return
	}

	if col < 0 {
		return
	}

	if col >= len(schematic[row]) {
		return
	}

	if schematic[row][col] < 48 || schematic[row][col] > 57 {
		return
	}

	// walk left till we get to the front of the number
	for col > 0 && schematic[row][col-1] >= 48 && schematic[row][col-1] <= 57 {
		col--
	}

	// record the location of the part number
	partNumber.Row = row
	partNumber.Col = col

	// while we are looking at a digit
	for col < len(schematic[row]) && schematic[row][col] >= 48 && schematic[row][col] <= 57 {
		partNumber.Number = (partNumber.Number * 10) + int(schematic[row][col]-48)
		col++
	}

	return
}
