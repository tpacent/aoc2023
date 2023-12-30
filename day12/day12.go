package day12

import (
	"aoc2023/lib"
	"bytes"
	"errors"
	"strconv"
	"strings"
)

const (
	Empty   = '.'
	Damaged = '#'
	Unknown = '?'
)

// map key type to elide expensive string operations
type CKey struct {
	Data string
	Spec string
}

func cacheKey(data []byte, spec []int) CKey {
	return CKey{Data: string(data), Spec: strints(spec)}
}

func strints(ints []int) string {
	sb := strings.Builder{}
	sb.Grow(len(ints))
	for _, n := range ints {
		if _, err := sb.WriteString(strconv.Itoa(n) + " "); err != nil {
			panic("unexpected err")
		}
	}
	return sb.String()
}

func NumArrangements(data []byte, spec []int) int {
	return numArrangements(data, spec, make(map[CKey]int, 64))
}

func numArrangements(data []byte, spec []int, cache map[CKey]int) (total int) {
	if len(spec) == 0 || len(data) < RequiredSize(spec) {
		return 0
	}

	if count, ok := cache[cacheKey(data, spec)]; ok {
		return count
	}

	slen, rest := spec[0], spec[1:]
	maxIndex := MaxIndex(data, slen)

	for index := range data {
		if index > maxIndex {
			break
		}
		if len(data)-index-slen < RequiredSize(rest) {
			break
		}
		if !CheckSegment(data, index, slen) {
			continue // cant use current segment
		}
		if len(rest) == 0 {
			if bytes.IndexByte(data[index+slen:], Damaged) < 0 {
				total++
			}
			continue
		}
		// could just use offset := index + slen + 1
		if offset := FindNextEligibleOffset(data, index+slen+1, rest[0]); offset >= 0 {
			total += numArrangements(data[offset:], rest, cache)
		}
	}

	cache[cacheKey(data, spec)] = total
	return
}

func RequiredSize(spec []int) int {
	return lib.Sum(spec...) + len(spec) - 1
}

func CheckSegment(data []byte, idx, slen int) bool {
	// check segment
	if bytes.IndexByte(data[idx:idx+slen], Empty) >= 0 {
		return false
	}
	// check ahead
	if idx+slen < len(data) && data[idx+slen] == Damaged {
		return false
	}
	return true
}

func MaxIndex(data []byte, slen int) (index int) {
	// segment wont fit into data
	if index = len(data) - slen; index < 0 {
		return -1
	}
	// scan up to the first damaged cell
	if dmg := bytes.IndexByte(data, Damaged); dmg >= 0 {
		index = min(index, dmg)
	}
	return
}

func ParseInput(line string) ([]byte, []int) {
	springs, groups, ok := strings.Cut(line, " ")
	if !ok {
		panic(errors.New("invalid format"))
	}
	intGroups := make([]int, 0)
	for _, s := range strings.Split(groups, ",") {
		intGroups = append(intGroups, lib.AsInt(s))
	}
	return []byte(springs), intGroups
}

func CountArrangements(input []string, repeat int) (total int) {
	for _, line := range input {
		springs, groups := ParseInput(line)
		springs, groups = Unfold(springs, groups, repeat)
		arrangements := NumArrangements(springs, groups)
		total += arrangements
	}
	return
}

func Unfold(springs []byte, groups []int, times int) ([]byte, []int) {
	springBox := make([][]byte, 0, times)
	for k := 0; k < times; k++ {
		springBox = append(springBox, springs)
	}
	return bytes.Join(springBox, []byte{Unknown}), lib.Repeat(groups, times)
}

// FindNextEligibleOffset is an unnecessary cache optimization
func FindNextEligibleOffset(data []byte, offset, slen int) int {
	var (
		seq     int
		seenDmg bool
		datalen = len(data)
	)

	if datalen < slen {
		return -1
	}

	for {
		if offset >= datalen {
			break
		}

		if data[offset] == Empty {
			if seenDmg {
				break
			}
			seq = 0
			offset++
			continue
		}

		seq++

		if seq >= slen {
			return offset - seq + 1
		}

		if data[offset] == Damaged {
			seenDmg = true
		}

		offset++
	}

	return -1
}
