package day03

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
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

			if number > 0 {
				// fmt.Printf("%d - [%d][%d]\n", number, row, col)
				//fmt.Println(number)
			}

			if isPartNumber {
				sum += number
			}

			col = endCol
		}
	}

	fmt.Printf("Day03 - 1: Total: %d\n", sum)
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
