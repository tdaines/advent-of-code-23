package day02_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day02"
)

func TestParseCubeAmount(t *testing.T) {
	r, g, b := day02.ParseCubeAmount(" 2 red")
	assert.Equal(t, 2, r)
	assert.Equal(t, 0, g)
	assert.Equal(t, 0, b)

	r, g, b = day02.ParseCubeAmount(" 8 green")
	assert.Equal(t, 0, r)
	assert.Equal(t, 8, g)
	assert.Equal(t, 0, b)

	r, g, b = day02.ParseCubeAmount(" 20 blue")
	assert.Equal(t, 0, r)
	assert.Equal(t, 0, g)
	assert.Equal(t, 20, b)
}

func TestParseCubeSet(t *testing.T) {
	cubeSet := day02.ParseCubeSet("3 blue, 4 red")
	assert.Equal(t, 4, cubeSet.Red)
	assert.Equal(t, 0, cubeSet.Green)
	assert.Equal(t, 3, cubeSet.Blue)

	cubeSet = day02.ParseCubeSet("1 red, 2 green, 6 blue")
	assert.Equal(t, 1, cubeSet.Red)
	assert.Equal(t, 2, cubeSet.Green)
	assert.Equal(t, 6, cubeSet.Blue)

	cubeSet = day02.ParseCubeSet("2 green")
	assert.Equal(t, 0, cubeSet.Red)
	assert.Equal(t, 2, cubeSet.Green)
	assert.Equal(t, 0, cubeSet.Blue)

	cubeSet = day02.ParseCubeSet("8 green, 6 blue, 20 red")
	assert.Equal(t, 20, cubeSet.Red)
	assert.Equal(t, 8, cubeSet.Green)
	assert.Equal(t, 6, cubeSet.Blue)
}

func TestParseGame(t *testing.T) {
	game := day02.ParseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	assert.Equal(t, 1, game.Id)
	assert.Equal(t, 3, len(game.CubeSets))
	assert.Equal(t, 4, game.CubeSets[0].Red)
	assert.Equal(t, 0, game.CubeSets[0].Green)
	assert.Equal(t, 3, game.CubeSets[0].Blue)
	assert.Equal(t, 1, game.CubeSets[1].Red)
	assert.Equal(t, 2, game.CubeSets[1].Green)
	assert.Equal(t, 6, game.CubeSets[1].Blue)
	assert.Equal(t, 0, game.CubeSets[2].Red)
	assert.Equal(t, 2, game.CubeSets[2].Green)
	assert.Equal(t, 0, game.CubeSets[2].Blue)

	game = day02.ParseGame("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")
	assert.Equal(t, 2, game.Id)
	assert.Equal(t, 3, len(game.CubeSets))
	assert.Equal(t, 0, game.CubeSets[0].Red)
	assert.Equal(t, 2, game.CubeSets[0].Green)
	assert.Equal(t, 1, game.CubeSets[0].Blue)
	assert.Equal(t, 1, game.CubeSets[1].Red)
	assert.Equal(t, 3, game.CubeSets[1].Green)
	assert.Equal(t, 4, game.CubeSets[1].Blue)
	assert.Equal(t, 0, game.CubeSets[2].Red)
	assert.Equal(t, 1, game.CubeSets[2].Green)
	assert.Equal(t, 1, game.CubeSets[2].Blue)

	game = day02.ParseGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	assert.Equal(t, 3, game.Id)
	assert.Equal(t, 3, len(game.CubeSets))
	assert.Equal(t, 20, game.CubeSets[0].Red)
	assert.Equal(t, 8, game.CubeSets[0].Green)
	assert.Equal(t, 6, game.CubeSets[0].Blue)
	assert.Equal(t, 4, game.CubeSets[1].Red)
	assert.Equal(t, 13, game.CubeSets[1].Green)
	assert.Equal(t, 5, game.CubeSets[1].Blue)
	assert.Equal(t, 1, game.CubeSets[2].Red)
	assert.Equal(t, 5, game.CubeSets[2].Green)
	assert.Equal(t, 0, game.CubeSets[2].Blue)

	game = day02.ParseGame("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")
	assert.Equal(t, 4, game.Id)
	assert.Equal(t, 3, len(game.CubeSets))
	assert.Equal(t, 3, game.CubeSets[0].Red)
	assert.Equal(t, 1, game.CubeSets[0].Green)
	assert.Equal(t, 6, game.CubeSets[0].Blue)
	assert.Equal(t, 6, game.CubeSets[1].Red)
	assert.Equal(t, 3, game.CubeSets[1].Green)
	assert.Equal(t, 0, game.CubeSets[1].Blue)
	assert.Equal(t, 14, game.CubeSets[2].Red)
	assert.Equal(t, 3, game.CubeSets[2].Green)
	assert.Equal(t, 15, game.CubeSets[2].Blue)

	game = day02.ParseGame("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")
	assert.Equal(t, 5, game.Id)
	assert.Equal(t, 2, len(game.CubeSets))
	assert.Equal(t, 6, game.CubeSets[0].Red)
	assert.Equal(t, 3, game.CubeSets[0].Green)
	assert.Equal(t, 1, game.CubeSets[0].Blue)
	assert.Equal(t, 1, game.CubeSets[1].Red)
	assert.Equal(t, 2, game.CubeSets[1].Green)
	assert.Equal(t, 2, game.CubeSets[1].Blue)
}
