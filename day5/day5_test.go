package day5_test

import (
	"aoc2023/day5"
	"aoc2023/lib"
	"math"
	"testing"
)

const Day5Part1Solution = 51752125

func TestSolveDay5Part1(t *testing.T) {
	input := lib.MustReadFile("testdata/input.txt")
	almanac := day5.ParseAlmanac(input)
	actual := math.MaxInt
	for _, seed := range almanac.Seeds {
		actual = min(actual, day5.MapThrough(seed, almanac.RangeMappers))
	}
	if actual != Day5Part1Solution {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

const Day5Part2Solution = 12634632

func TestSolveDay5Part2(t *testing.T) {
	input := lib.MustReadFile("testdata/input.txt")
	almanac := day5.ParseAlmanac(input)
	ranges := day5.MakeSeedRanges(almanac.Seeds)
	for _, rm := range almanac.RangeMappers {
		ranges = day5.BreakRangeAll(ranges, rm.Mappings)
	}
	if actual := ranges[0].Start; actual != Day5Part2Solution {
		t.Error("unexpected value")
	}
}

func TestParseMapping(t *testing.T) {
	input := "39 0 15"
	actual := day5.ParseMapping(input)
	if actual.Start != 0 {
		t.Error("unexpected start")
	}
	if actual.End != 15 {
		t.Error("unexpected end")
	}
	if actual.Offset != 39 {
		t.Error("unexpected offset")
	}
}

var testmap = []string{
	"50 98 2",
	"52 50 48",
}

func TestRangeMapper(t *testing.T) {
	mappings := make([]*day5.Mapping, 0)
	for _, line := range testmap {
		mappings = append(mappings, day5.ParseMapping(line))
	}
	rm := day5.NewRangeMapper(mappings, "")
	if actual := rm.Map(40); actual != 40 {
		t.Error("unexpected value")
	}
	if actual := rm.Map(51); actual != 53 {
		t.Error("unexpected value")
	}
	if actual := rm.Map(99); actual != 51 {
		t.Error("unexpected value")
	}
}

func TestParseSeeds(t *testing.T) {
	actual := day5.ParseSeeds("seeds: 79 14 55 13")
	if len(actual) != 4 {
		t.Error("unexpected seed count")
	}
}

func TestAlmanac(t *testing.T) {
	input := lib.MustReadFile("testdata/test.txt")
	almanac := day5.ParseAlmanac(input)
	if len(almanac.RangeMappers) != 7 {
		t.Error("unexpected mapper count")
	}
}

func TestBreakAll(t *testing.T) {
	input := lib.MustReadFile("testdata/test.txt")
	almanac := day5.ParseAlmanac(input)
	ranges := day5.MakeSeedRanges(almanac.Seeds)
	for _, rm := range almanac.RangeMappers {
		ranges = day5.BreakRangeAll(ranges, rm.Mappings)
	}
	if actual := ranges[0].Start; actual != 46 {
		t.Error("unexpected value")
	}
}
