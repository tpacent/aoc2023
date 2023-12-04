package day4_test

import (
	"aoc2023/day4"
	"aoc2023/lib"
	"testing"
)

const Day4Part1Solution = 17803

func TestSolveDay4Part1(t *testing.T) {
	actual := 0
	for _, line := range lib.MustReadFile("testdata/input.txt") {
		info := day4.ParseLine(line)
		actual += day4.CardPoints(day4.MatchCount(info.Numbers, info.Pile))
	}
	if actual != Day4Part1Solution {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

const Day4Part2Solution = 5554894

func TestSolveDay4Part2(t *testing.T) {
	cards := make([]*day4.CardInfo, 0)
	for _, line := range lib.MustReadFile("testdata/input.txt") {
		cards = append(cards, day4.ParseLine(line))
	}
	actual := day4.CountCards(cards)
	if actual != Day4Part2Solution {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

func TestParseLine(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	info := day4.ParseLine(input)
	if info.ID != 1 {
		t.Error("unexpected card ID")
	}
	if len(info.Numbers) != 5 {
		t.Error("unexpected numbers length")
	}
	if len(info.Pile) != 8 {
		t.Error("unexpected pile length")
	}
}

var testcards = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

func TestCountScratchcards(t *testing.T) {
	cards := make([]*day4.CardInfo, 0)
	for _, line := range testcards {
		cards = append(cards, day4.ParseLine(line))
	}
	actual := day4.CountCards(cards)
	if actual != 30 {
		t.Error("unexpected value")
	}
}
