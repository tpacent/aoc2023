package day15_test

import (
	"aoc2023/day15"
	"aoc2023/lib"
	"bytes"
	"testing"
)

const SolutionDay15Part1 = 510792

func TestSolveDay15Part1(t *testing.T) {
	data := lib.MustReadFileBytes("testdata/input.txt")
	actual := day15.SumHashes(data[0])
	if actual != SolutionDay15Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

const SolutionDay15Part2 = 269410

func TestSolveDay15Part2(t *testing.T) {
	data := lib.MustReadFileBytes("testdata/input.txt")
	hm := new(day15.HashMap)
	for _, op := range bytes.Split([]byte(data[0]), []byte{','}) {
		day15.HashOp(hm, op)
	}
	actual := hm.Power()
	if actual != SolutionDay15Part2 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

func TestHash(t *testing.T) {
	if actual := day15.Hash([]byte("HASH")); actual != 52 {
		t.Error("unexpected value", actual)
	}
}

func TestHashMap(t *testing.T) {
	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

	var hm = new(day15.HashMap)

	for _, op := range bytes.Split([]byte(input), []byte{','}) {
		day15.HashOp(hm, op)
	}

	if actual := hm.Power(); actual != 145 {
		t.Error("unexpected value", actual)
	}
}
