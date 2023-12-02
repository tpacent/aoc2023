package day2_test

import (
	"aoc2023/day2"
	"aoc2023/lib"
	"testing"
)

const Day2Part1Solution = 2771

func TestSolveDay2Part1(t *testing.T) {
	var actual int

	for _, line := range lib.MustReadFile("testdata/input.txt") {
		game := day2.ParseGame(line)
		if day2.IsPlayable(game, 12, 13, 14) {
			actual += game.ID
		}
	}

	if actual != Day2Part1Solution {
		t.Error("unexpected value")
	}

	t.Log(actual)
}

const Day2Part2Solution = 70924

func TestSolveDay2Part2(t *testing.T) {
	var actual int // sum of game powers

	for _, line := range lib.MustReadFile("testdata/input.txt") {
		game := day2.ParseGame(line)
		actual += day2.GamePower(game)
	}

	if actual != Day2Part2Solution {
		t.Error("unexpected value")
	}

	t.Log(actual)
}

type GameParseTest struct {
	Input    string
	Rounds   int
	Playable bool
}

var gameParseTests = []GameParseTest{
	{Input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", Rounds: 3, Playable: true},
	{Input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", Rounds: 3, Playable: true},
	{Input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", Rounds: 3, Playable: false},
	{Input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", Rounds: 3, Playable: false},
	{Input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", Rounds: 2, Playable: true},
}

func TestParseGame(t *testing.T) {
	for _, test := range gameParseTests {
		game := day2.ParseGame(test.Input)
		if len(game.Rounds) != test.Rounds {
			t.Error("unexpected round count")
		}
		if day2.IsPlayable(game, 12, 13, 14) != test.Playable {
			t.Error("playable wrong result")
		}
	}
}

type PowerTest struct {
	Input    string
	Expected int
}

var gamePowerTests = []PowerTest{
	{Input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", Expected: 48},
	{Input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", Expected: 12},
	{Input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", Expected: 1560},
	{Input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", Expected: 630},
	{Input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", Expected: 36},
}

func TestGamePower(t *testing.T) {
	for _, test := range gamePowerTests {
		game := day2.ParseGame(test.Input)

		if day2.GamePower(game) != test.Expected {
			t.Error("unexpected value")
		}
	}
}
