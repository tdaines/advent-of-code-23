package day01

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1() {
	input, err := os.Open("./day01/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var sum int = 0
	for scanner.Scan() {
		line := scanner.Text()
		value := GetCalibrationValue(line)
		sum += value
	}
	fmt.Printf("Day01 - 1: Total: %d\n", sum)
}

func Part2() {
	input, err := os.Open("./day01/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var sum int = 0
	for scanner.Scan() {
		line := scanner.Text()
		value := GetRealCalibrationValue(line)
		sum += value
	}

	fmt.Printf("Day01 - 2: Total: %d\n", sum)
}

func GetCalibrationValue(line string) int {
	var firstDigit int = GetFirstDigit(line)
	var secondDigit int = GetLastDigit(line)

	return (firstDigit * 10) + secondDigit
}

func GetFirstDigit(line string) int {
	for _, char := range line {
		asciiValue := int(char)
		if asciiValue >= 48 && asciiValue <= 57 {
			return asciiValue - 48
		}
	}
	return 0
}

func GetLastDigit(line string) int {
	for x := len(line) - 1; x >= 0; x-- {
		asciiValue := int(line[x])
		if asciiValue >= 48 && asciiValue <= 57 {
			return asciiValue - 48
		}
	}
	return 0
}

func GetRealCalibrationValue(line string) int {
	var firstDigit int = GetRealFirstDigit(line)
	var secondDigit int = GetRealLastDigit(line)

	return (firstDigit * 10) + secondDigit
}

var prefixes = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func GetRealFirstDigit(line string) int {
	for len(line) > 0 {
		for key, val := range prefixes {
			if strings.HasPrefix(line, key) {
				return val
			}
		}

		line = line[1:]
	}

	return 0
}

func GetRealLastDigit(line string) int {
	var digit int
	for len(line) > 0 {
		for key, val := range prefixes {
			if strings.HasPrefix(line, key) {
				digit = val
			}
		}

		line = line[1:]
	}

	return digit
}
