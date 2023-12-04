package day4

import (
	"aoc2023/lib"
	"math"
	"strings"
)

type CardInfo struct {
	ID      int
	Numbers []int
	Pile    map[int]struct{}
}

func ParseLine(line string) *CardInfo {
	cardnum, line, ok := strings.Cut(line, ":")
	if !ok {
		panic("unexpected input")
	}

	idstr := cardnum[strings.LastIndex(cardnum, " ")+1:]
	numstr, pilestr, ok := strings.Cut(line, "|")
	if !ok {
		panic("unexpected input")
	}

	info := &CardInfo{
		ID:      lib.AsInt(idstr),
		Numbers: ParseFields(numstr),
		Pile:    make(map[int]struct{}),
	}

	for _, n := range ParseFields(pilestr) {
		info.Pile[n] = struct{}{}
	}

	return info
}

func ParseFields(chunk string) (nums []int) {
	for _, nstr := range strings.Fields(chunk) {
		nums = append(nums, lib.AsInt(nstr))
	}

	return
}

func MatchCount(card *CardInfo) (count int) {
	for _, n := range card.Numbers {
		if _, ok := card.Pile[n]; ok {
			count++
		}
	}

	return
}

func CardPoints(card *CardInfo) int {
	matchCount := MatchCount(card)

	if matchCount == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matchCount-1)))
}

func CountCards(initialHand []*CardInfo) (total int) {
	cardRegistry := make(map[int]*CardInfo, len(initialHand))
	for _, card := range initialHand {
		cardRegistry[card.ID] = card
	}

	total += len(initialHand)

	prevHand := make(map[int]int, len(initialHand))
	for _, card := range initialHand {
		prevHand[card.ID]++
	}

	for {
		currHand := make(map[int]int, 0)

		for id, copies := range prevHand {
			card := cardRegistry[id]
			count := MatchCount(card)
			from := card.ID // off-by-one
			upto := from + count
			total += count * copies

			for _, c := range initialHand[from:upto] {
				currHand[c.ID] += copies
			}
		}

		if len(currHand) == 0 {
			break
		}

		prevHand = currHand
	}

	return
}
