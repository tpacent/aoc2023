package day2

import (
	"aoc2023/lib"
	"strings"
)

type Round struct {
	Red   int
	Green int
	Blue  int
}

type GameSummary struct {
	ID     int
	Rounds []Round
}

// GetMax returns the number of stones required to play every round in the game
func (gs *GameSummary) GetMax() (r int, g int, b int) {
	for _, round := range gs.Rounds {
		r = max(r, round.Red)
		g = max(g, round.Green)
		b = max(b, round.Blue)
	}

	return
}

// ParseGame parses a line from input (without using regex)
func ParseGame(line string) *GameSummary {
	idStr, roundsStr, ok := strings.Cut(line, ":")

	if !ok {
		panic("unexpected input")
	}

	summary := GameSummary{
		ID:     parseID(idStr),
		Rounds: make([]Round, 0),
	}

	for _, roundstr := range strings.Split(roundsStr, ";") {
		round := parseRound(roundstr)
		summary.Rounds = append(summary.Rounds, round)
	}

	return &summary
}

// parseRound handles strings similar to "1 red, 2 green, 6 blue"
func parseRound(input string) Round {
	stats := make(map[string]int)

	for _, group := range strings.Split(input, ",") {
		countstr, groupKey, ok := strings.Cut(strings.TrimSpace(group), " ")

		if !ok {
			panic("unexpected input")
		}

		stats[groupKey] = lib.AsInt(countstr)
	}

	return Round{
		Red:   stats["red"],
		Green: stats["green"],
		Blue:  stats["blue"],
	}
}

// parseID extracts game id, expects input strings like "Game 1"
func parseID(input string) int {
	_, input, ok := strings.Cut(input, " ")

	if !ok {
		panic("unexpected input")
	}

	return lib.AsInt(input)
}

// IsPlayable tells if the game can be played with the specified number of stones
func IsPlayable(game *GameSummary, haveRed, haveGreen, haveBlue int) bool {
	for _, round := range game.Rounds {
		if round.Red > haveRed || round.Green > haveGreen || round.Blue > haveBlue {
			return false
		}
	}

	return true
}

// GamePower is defined as numbers of cubes multiplied together
func GamePower(game *GameSummary) int {
	r, g, b := game.GetMax()
	return r * g * b
}
