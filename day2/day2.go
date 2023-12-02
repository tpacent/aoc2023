package day2

import (
	"strconv"
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

func (gs *GameSummary) GetMax() (r int, g int, b int) {
	for _, round := range gs.Rounds {
		r = max(r, round.Red)
		g = max(g, round.Green)
		b = max(b, round.Blue)
	}

	return
}

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

func parseRound(input string) Round {
	stats := make(map[string]int)

	for _, group := range strings.Split(input, ",") {
		countstr, groupKey, ok := strings.Cut(strings.TrimSpace(group), " ")

		if !ok {
			panic("unexpected input")
		}

		if count, err := strconv.Atoi(countstr); err == nil {
			stats[groupKey] = count
		}
	}

	return Round{
		Red:   stats["red"],
		Green: stats["green"],
		Blue:  stats["blue"],
	}
}

func parseID(input string) int {
	_, input, ok := strings.Cut(input, " ")

	if !ok {
		panic("unexpected input")
	}

	id, err := strconv.Atoi(input)

	if err != nil {
		panic("unexpected input")
	}

	return id
}

func IsPlayable(game *GameSummary, haveRed, haveGreen, haveBlue int) bool {
	for _, round := range game.Rounds {
		if round.Red > haveRed || round.Green > haveGreen || round.Blue > haveBlue {
			return false
		}
	}

	return true
}

func GamePower(game *GameSummary) int {
	r, g, b := game.GetMax()
	return r * g * b
}
