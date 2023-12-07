package day7

import (
	"aoc2023/lib"
	"cmp"
	"strings"
)

type HandType byte

const (
	HighCard HandType = 1 + iota
	Pair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

type HandInfo struct {
	Hand []rune
	Type HandType
	Bid  int
}

func ParseInput(input []string) (hands []*HandInfo) {
	for _, line := range input {
		handstr, bidstr, ok := strings.Cut(line, " ")
		if !ok {
			panic("unexpected input")
		}
		hand := []rune(handstr)
		hands = append(hands, &HandInfo{Hand: hand, Bid: lib.AsInt(bidstr), Type: GetHandType(hand)})
	}
	return
}

var cardStrength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func GetHandType(hand []rune) HandType {
	histogram := make(map[rune]int, len(hand))
	var maxStack int
	for _, r := range hand {
		histogram[r]++
	}
	for _, count := range histogram {
		if count > maxStack {
			maxStack = count
		}
	}
	if maxStack == 5 {
		return FiveOfKind
	}
	if maxStack == 4 {
		return FourOfKind
	}
	if maxStack == 3 {
		if len(histogram) == 2 {
			return FullHouse
		} else {
			return ThreeOfKind
		}
	}
	if maxStack == 2 {
		if len(histogram) == 3 {
			return TwoPair
		} else {
			return Pair
		}
	}
	if maxStack == 1 && len(histogram) == 5 {
		return HighCard
	}
	panic("unreachable")
}

func HandSorter(a, b *HandInfo) int {
	if typeCmp := cmp.Compare(GetHandType(a.Hand), GetHandType(b.Hand)); typeCmp != 0 {
		return typeCmp
	}

	for k := 0; k < len(a.Hand); k++ {
		if cardCmp := cmp.Compare(cardStrength[a.Hand[k]], cardStrength[b.Hand[k]]); cardCmp != 0 {
			return cardCmp
		}
	}

	return 0
}

func Winnings(hands []*HandInfo) (winnings int) {
	for index, hand := range hands {
		winnings += (index + 1) * hand.Bid
	}

	return
}
