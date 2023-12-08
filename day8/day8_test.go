package day8_test

import (
	"aoc2023/day8"
	"aoc2023/lib"
	"testing"
)

var SolutionDay8Part1 = 19631

func TestSolveDay8Part1(t *testing.T) {
	input := lib.MustReadFile("testdata/input.txt")
	rule, network := day8.ParseInput(input)
	actual := day8.Traverse("AAA", "ZZZ", rule, network)
	if actual != SolutionDay8Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var SolutionDay8Part2 = 21003205388413

func TestSolveDay8Part2(t *testing.T) {
	input := lib.MustReadFile("testdata/input.txt")
	rule, network := day8.ParseInput(input)
	actual := day8.GhostTraverse(rule, network)
	if actual != SolutionDay8Part2 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testInput = []string{
	"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)",
}

func TestExample(t *testing.T) {
	rule, network := day8.ParseInput(testInput)
	actual := day8.Traverse("AAA", "ZZZ", rule, network)
	if actual != 6 {
		t.Error("unexpected value")
	}
}

var testInput2 = []string{
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

func TestExample2(t *testing.T) {
	rule, network := day8.ParseInput(testInput2)
	actual := day8.GhostTraverse(rule, network)
	if actual != 6 {
		t.Error("unexpected value")
	}
}
