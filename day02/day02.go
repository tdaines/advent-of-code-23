package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	Id       int
	CubeSets []CubeSet
}

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

func Part1() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day02/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var sum int = 0
	for scanner.Scan() {
		line := scanner.Text()

		game := ParseGame(line)
		if IsGamePossible(game, 12, 13, 14) {
			sum += game.Id
		}
	}

	answer = sum
	return answer, time.Since(now)
}

func Part2() (answer int, elapsed time.Duration) {
	var now = time.Now()
	input, err := os.Open("./day02/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var sum int = 0
	for scanner.Scan() {
		line := scanner.Text()

		game := ParseGame(line)
		minRed, minGreen, minBlue := GetMinNumCubes(game)
		power := minRed * minGreen * minBlue
		sum += power
	}

	answer = sum
	return answer, time.Since(now)
}

func IsGamePossible(game Game, maxRed int, maxGreen int, maxBlue int) bool {
	for _, cubeSet := range game.CubeSets {
		if cubeSet.Red > maxRed || cubeSet.Green > maxGreen || cubeSet.Blue > maxBlue {
			return false
		}
	}

	return true
}

func GetMinNumCubes(game Game) (r int, g int, b int) {
	for _, cubeSet := range game.CubeSets {
		r = max(r, cubeSet.Red)
		g = max(g, cubeSet.Green)
		b = max(b, cubeSet.Blue)

	}
	return
}

func ParseGame(line string) Game {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

	var game = Game{}

	// remove "Game " from the start of line
	line = line[len("Game "):]

	parts := strings.Split(line, ":")

	game.Id, _ = strconv.Atoi(parts[0])

	parts = strings.Split(parts[1], ";")
	for _, part := range parts {
		cubeSet := ParseCubeSet(part)
		game.CubeSets = append(game.CubeSets, cubeSet)
	}

	return game
}

func ParseCubeSet(line string) CubeSet {
	// 3 blue, 4 red
	// 1 red, 2 green, 6 blue
	// 2 green

	var cubeSet = CubeSet{}

	parts := strings.Split(line, ",")
	for _, part := range parts {
		r, g, b := ParseCubeAmount(part)
		cubeSet.Red += r
		cubeSet.Green += g
		cubeSet.Blue += b
	}

	return cubeSet
}

func ParseCubeAmount(line string) (red int, green int, blue int) {
	// 3 blue
	// 4 red
	// 20 green
	line = strings.TrimLeft(line, " ")
	parts := strings.Split(line, " ")

	value, _ := strconv.Atoi(parts[0])

	if strings.HasPrefix(parts[1], "red") {
		red = value
	} else if strings.HasPrefix(parts[1], "green") {
		green = value
	} else {
		blue = value
	}
	return
}
