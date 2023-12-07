package day7_test

import (
	"aoc2023/day7"
	"aoc2023/lib"
	"slices"
	"testing"
)

const SolutionDay7Part1 = 250347426

func TestSolveDay7Part1(t *testing.T) {
	input := lib.MustReadFile("testdata/input.txt")
	hands := day7.ParseInput(input)
	slices.SortFunc(hands, day7.HandSorter)
	actual := day7.Winnings(hands)
	if actual != SolutionDay7Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testinput = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestExample(t *testing.T) {
	hands := day7.ParseInput(testinput)
	slices.SortFunc(hands, day7.HandSorter)
	if actual := day7.Winnings(hands); actual != 6440 {
		t.Error("unexpected value", actual)
	}
}

type HandTypeTestCase struct {
	Hand     []rune
	Expected day7.HandType
}

var handTypeTestCases = []*HandTypeTestCase{
	{Hand: []rune{'3', '2', 'T', '3', 'K'}, Expected: day7.Pair},
	{Hand: []rune{'T', '5', '5', 'J', '5'}, Expected: day7.ThreeOfKind},
	{Hand: []rune{'K', 'K', '6', '7', '7'}, Expected: day7.TwoPair},
	{Hand: []rune{'K', 'T', 'J', 'J', 'T'}, Expected: day7.TwoPair},
	{Hand: []rune{'Q', 'Q', 'Q', 'J', 'A'}, Expected: day7.ThreeOfKind},
}

func TestHandType(t *testing.T) {
	for _, test := range handTypeTestCases {
		actual := day7.GetHandType(test.Hand)
		if actual != test.Expected {
			t.Error("unexpected hand type")
		}
	}
}
