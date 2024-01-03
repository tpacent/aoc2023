package day15

import (
	"aoc2023/lib"
	"bytes"
	"slices"
)

type Lens struct {
	Label string
	Power uint8
}

type HashMap [256][]Lens

func (hm *HashMap) OpUpsert(lens Lens) {
	bucket := Hash([]byte(lens.Label))
	itemIndex := slices.IndexFunc(hm[bucket], func(l Lens) bool { return l.Label == lens.Label })
	if itemIndex < 0 {
		hm[bucket] = append(hm[bucket], lens)
	} else {
		hm[bucket][itemIndex] = lens
	}
}

func (hm *HashMap) OpRemove(label []byte) {
	bucket := Hash(label)
	itemIndex := labelIndex(hm[bucket], string(label))
	if itemIndex < 0 {
		return
	}
	hm[bucket] = slices.Delete(hm[bucket], itemIndex, itemIndex+1)
}

func labelIndex(lens []Lens, label string) int {
	return slices.IndexFunc(lens, func(l Lens) bool { return l.Label == string(label) })
}

func (hm *HashMap) Power() (total int) {
	for bucketIndex, bucket := range hm {
		for lensIndex, lens := range bucket {
			total += (1 + bucketIndex) * (1 + lensIndex) * int(lens.Power)
		}
	}

	return
}

func HashOp(hm *HashMap, op []byte) {
	index := bytes.IndexAny(op, "-=")
	label := op[:index]
	switch op[index] {
	case '=':
		hm.OpUpsert(Lens{Label: string(label), Power: uint8(lib.AsInt(string(op[index+1:])))})
	case '-':
		hm.OpRemove(label)
	default:
		panic("unknown op")
	}
}

func Hash(str []byte) (value uint8) {
	for _, char := range str {
		value = (value + char) * 17
	}
	return
}

func SumHashes(line []byte) (sum int) {
	for _, str := range bytes.Split(line, []byte{','}) {
		sum += int(Hash(str))
	}

	return
}
