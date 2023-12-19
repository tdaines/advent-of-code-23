package day01_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-23/day01"
)

func TestGetFirstDigit(t *testing.T) {
	assert.Equal(t, 1, day01.GetFirstDigit("1abc2"))
	assert.Equal(t, 3, day01.GetFirstDigit("pqr3stu8vwx"))
	assert.Equal(t, 1, day01.GetFirstDigit("a1b2c3d4e5f"))
	assert.Equal(t, 7, day01.GetFirstDigit("treb7uchet"))
}

func TestGetLastDigit(t *testing.T) {
	assert.Equal(t, 2, day01.GetLastDigit("1abc2"))
	assert.Equal(t, 8, day01.GetLastDigit("pqr3stu8vwx"))
	assert.Equal(t, 5, day01.GetLastDigit("a1b2c3d4e5f"))
	assert.Equal(t, 7, day01.GetLastDigit("treb7uchet"))
}

func TetGetCalibrationValue(t *testing.T) {
	assert.Equal(t, 12, day01.GetLastDigit("1abc2"))
	assert.Equal(t, 38, day01.GetLastDigit("pqr3stu8vwx"))
	assert.Equal(t, 15, day01.GetLastDigit("a1b2c3d4e5f"))
	assert.Equal(t, 77, day01.GetLastDigit("treb7uchet"))
}

func TestGetRealFirstDigit(t *testing.T) {
	assert.Equal(t, 2, day01.GetRealFirstDigit("two1nine"))
	assert.Equal(t, 8, day01.GetRealFirstDigit("eightwothree"))
	assert.Equal(t, 1, day01.GetRealFirstDigit("abcone2threexyz"))
	assert.Equal(t, 2, day01.GetRealFirstDigit("xtwone3four"))
	assert.Equal(t, 4, day01.GetRealFirstDigit("4nineeightseven2"))
	assert.Equal(t, 1, day01.GetRealFirstDigit("zoneight234"))
	assert.Equal(t, 7, day01.GetRealFirstDigit("7pqrstsixteen"))
}

func TestGetRealLastDigit(t *testing.T) {
	assert.Equal(t, 9, day01.GetRealLastDigit("two1nine"))
	assert.Equal(t, 3, day01.GetRealLastDigit("eightwothree"))
	assert.Equal(t, 3, day01.GetRealLastDigit("abcone2threexyz"))
	assert.Equal(t, 4, day01.GetRealLastDigit("xtwone3four"))
	assert.Equal(t, 2, day01.GetRealLastDigit("4nineeightseven2"))
	assert.Equal(t, 4, day01.GetRealLastDigit("zoneight234"))
	assert.Equal(t, 6, day01.GetRealLastDigit("7pqrstsixteen"))
}

func TestGetRealCalibrationValue(t *testing.T) {
	assert.Equal(t, 29, day01.GetRealCalibrationValue("two1nine"))
	assert.Equal(t, 83, day01.GetRealCalibrationValue("eightwothree"))
	assert.Equal(t, 13, day01.GetRealCalibrationValue("abcone2threexyz"))
	assert.Equal(t, 24, day01.GetRealCalibrationValue("xtwone3four"))
	assert.Equal(t, 42, day01.GetRealCalibrationValue("4nineeightseven2"))
	assert.Equal(t, 14, day01.GetRealCalibrationValue("zoneight234"))
	assert.Equal(t, 76, day01.GetRealCalibrationValue("7pqrstsixteen"))
}
