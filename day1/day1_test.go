package day1_test

import (
	"aoc2023/day1"
	"aoc2023/lib"
	"testing"
)

const Day1Part1Solution = 54634

func TestSolveDay1Part1(t *testing.T) {
	var actual int

	for _, line := range lib.MustReadFile("testdata/input.txt") {
		actual += day1.MustRecoverValue(line)
	}

	if actual != Day1Part1Solution {
		t.Error("unexpected value")
	}

	t.Log(actual)
}

const Day1Part2Solution = 53855

func TestSolveDay1Part2(t *testing.T) {
	var actual int

	for _, line := range lib.MustReadFile("testdata/input.txt") {
		n := day1.MustRecoverSpelledValue(line)
		actual += n
	}

	t.Log(actual)
}

type RecoverValueTestData struct {
	Line     string
	Expected int
}

var recoverValueTestData1 = []RecoverValueTestData{
	{Line: "1abc2", Expected: 12},
	{Line: "pqr3stu8vwx", Expected: 38},
	{Line: "a1b2c3d4e5f", Expected: 15},
	{Line: "treb7uchet", Expected: 77},
}

func TestMustRecoverValue(t *testing.T) {
	for _, test := range recoverValueTestData1 {
		actual := day1.MustRecoverValue(test.Line)
		if actual != test.Expected {
			t.Error("unexpected value")
		}
	}
}

var recoverValueTestData2 = []RecoverValueTestData{
	{Line: "two1nine", Expected: 29},
	{Line: "eightwothree", Expected: 83},
	{Line: "abcone2threexyz", Expected: 13},
	{Line: "xtwone3four", Expected: 24},
	{Line: "4nineeightseven2", Expected: 42},
	{Line: "zoneight234", Expected: 14},
	{Line: "7pqrstsixteen", Expected: 76},
}

func TestRecoverNormalizedValue(t *testing.T) {
	var sum int

	for _, test := range recoverValueTestData2 {
		actual := day1.MustRecoverSpelledValue(test.Line)
		if actual != test.Expected {
			t.Errorf("unexpected value: expected %d, got %d", test.Expected, actual)
		}
		sum += actual
	}

	if sum != 281 {
		t.Error("unexpected sum")
	}
}
