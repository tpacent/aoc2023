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
	Hand  []rune
	JHand []rune // hand with jokers applied to maximize value
	Type  HandType
	Bid   int
}

func MaximizeHand(hand []rune) []rune {
	histogram := make(map[rune]int, len(hand))

	for _, rune := range hand {
		histogram[rune]++
	}

	jCount := histogram['J']

	if jCount == 0 {
		return hand
	}

	if jCount == len(hand) {
		return []rune{'A', 'A', 'A', 'A', 'A'}
	}

	delete(histogram, 'J')
	var maxRune rune
	var maxCount int
	for r, count := range histogram {
		if count > maxCount {
			maxCount = count
			maxRune = r
		}
	}

	maxHand := append([]rune(nil), hand...)
	for k := 0; k < len(maxHand); k++ {
		if maxHand[k] == 'J' {
			maxHand[k] = maxRune
		}
	}

	return maxHand
}

func ParseInput(input []string) (hands []*HandInfo) {
	for _, line := range input {
		handstr, bidstr, ok := strings.Cut(line, " ")
		if !ok {
			panic("unexpected input")
		}
		hand := []rune(handstr)
		hands = append(hands, &HandInfo{
			Hand:  hand,
			Bid:   lib.AsInt(bidstr),
			Type:  GetHandType(hand),
			JHand: MaximizeHand(hand),
		})
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

func Strength(card rune) int {
	return cardStrength[card]
}

func JStrength(card rune) int {
	if card == 'J' {
		return 1
	}
	return Strength(card)
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
		if cardCmp := cmp.Compare(Strength(a.Hand[k]), Strength(b.Hand[k])); cardCmp != 0 {
			return cardCmp
		}
	}

	return 0
}

func JHandSorter(a, b *HandInfo) int {
	if typeCmp := cmp.Compare(GetHandType(a.JHand), GetHandType(b.JHand)); typeCmp != 0 {
		return typeCmp
	}

	for k := 0; k < len(a.Hand); k++ {
		if cardCmp := cmp.Compare(JStrength(a.Hand[k]), JStrength(b.Hand[k])); cardCmp != 0 {
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
