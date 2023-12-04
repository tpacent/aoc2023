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

	numstr, pilestr, ok := strings.Cut(line, "|")

	if !ok {
		panic("unexpected input")
	}

	info := &CardInfo{
		ID:      lib.AsInt(cardnum[strings.LastIndex(cardnum, " ")+1:]),
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

func FindMatching(numbers []int, pile map[int]struct{}) (matches []int) {
	for _, n := range numbers {
		if _, ok := pile[n]; ok {
			matches = append(matches, n)
		}
	}

	return
}

func MatchCount(numbers []int, pile map[int]struct{}) int {
	return len(FindMatching(numbers, pile))
}

func CardPoints(matchCount int) int {
	if matchCount == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matchCount-1)))
}

func CountCards(initialHand []*CardInfo) (total int) {
	allCards := make(map[int]*CardInfo, len(initialHand))
	for _, card := range initialHand {
		allCards[card.ID] = card
	}

	total += len(initialHand)

	prevHand := make(map[int]int, len(initialHand))
	for _, card := range initialHand {
		prevHand[card.ID]++
	}

	for {
		currHand := make(map[int]int, 0)

		for id, copies := range prevHand {
			card := allCards[id]
			count := MatchCount(card.Numbers, card.Pile)
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
