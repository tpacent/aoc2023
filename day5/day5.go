package day5

import (
	"aoc2023/lib"
	"cmp"
	"slices"
	"strings"
)

type Mapping struct {
	Start  int
	End    int
	Offset int
}

var MappingsSorter = func(a, b *Mapping) int {
	return cmp.Compare(a.Start, b.Start)
}

func NewRangeMapper(mappings []*Mapping, name string) *RangeMapper {
	m := append([]*Mapping(nil), mappings...)
	slices.SortFunc(m, MappingsSorter)
	return &RangeMapper{Mappings: m, Name: name}
}

type RangeMapper struct {
	Mappings []*Mapping
	Name     string
}

func (rm *RangeMapper) Map(value int) int {
	for _, mapper := range rm.Mappings {
		if mapper.Start > value {
			break
		}

		if value >= mapper.Start && value < mapper.End {
			return value + mapper.Offset
		}
	}

	return value // numbers that aren't mapped correspond to the same number
}

func ParseMapping(line string) *Mapping {
	fields := strings.Fields(line)
	dst := lib.AsInt(fields[0])
	src := lib.AsInt(fields[1])
	len := lib.AsInt(fields[2])
	return &Mapping{
		Start:  src,
		End:    src + len,
		Offset: dst - src,
	}
}

type Almanac struct {
	Seeds        []int
	RangeMappers []*RangeMapper
}

func MapThrough(value int, mappers []*RangeMapper) int {
	for _, mapper := range mappers {
		value = mapper.Map(value)
	}

	return value
}

func ParseAlmanac(input []string) *Almanac {
	almanac := Almanac{
		Seeds:        ParseSeeds(input[0]),
		RangeMappers: make([]*RangeMapper, 0),
	}

	var mappings []*Mapping
	var name string
	for _, line := range input[1:] {
		if line == "" {
			if len(mappings) > 0 {
				almanac.RangeMappers = append(almanac.RangeMappers, NewRangeMapper(mappings, name))
			}
			mappings = make([]*Mapping, 0)
			continue
		}

		if strings.Contains(line, "map") {
			name = line
			continue
		}

		mappings = append(mappings, ParseMapping(line))
	}
	// last one
	almanac.RangeMappers = append(almanac.RangeMappers, NewRangeMapper(mappings, name))
	return &almanac
}

func ParseSeeds(line string) (seeds []int) {
	strseeds := strings.Fields(line)[1:]
	for _, seed := range strseeds {
		seeds = append(seeds, lib.AsInt(seed))
	}
	return
}

func MakeSeedRanges(seeds []int) (out []*Mapping) {
	for k := 0; k < len(seeds); k += 2 {
		out = append(out, &Mapping{
			Start: seeds[k],
			End:   seeds[k] + seeds[k+1],
		})
	}

	return
}

func BreakRangeAll(a []*Mapping, b []*Mapping) []*Mapping {
	out := make([]*Mapping, 0)

	for _, ar := range a {
		addRanges := BreakRange(ar, b)

		if len(addRanges) > 0 {
			out = append(out, addRanges...)
		} else {
			out = append(out, ar)
		}
	}

	return SortMergeRange(out)
}

func SortMergeRange(a []*Mapping) (out []*Mapping) {
	if len(a) < 2 {
		return a
	}

	slices.SortFunc(a, MappingsSorter)

	curr := a[0]

	for _, chunk := range a[1:] {
		if Intersects(curr, chunk) {
			curr.End = chunk.End
		} else {
			out = append(out, curr)
			curr = chunk
		}
	}

	out = append(out, curr)
	return
}

func Intersects(a, b *Mapping) bool {
	if b.End <= a.Start {
		return false
	}

	if b.Start >= a.End {
		return false
	}

	return true
}

func BreakRange(a *Mapping, ranges []*Mapping) []*Mapping {
	out := make([]*Mapping, 0)

	from := a.Start
	upto := a.End

	for _, b := range ranges {
		if b.End <= from {
			continue
		}
		if b.Start >= upto {
			continue
		}

		if b.Start > from {
			out = append(out, &Mapping{Start: from, End: b.Start})
			from = b.Start
		}

		tail := min(upto, b.End)
		out = append(out, &Mapping{Start: from + b.Offset, End: tail + b.Offset})
		from = tail
	}

	return out
}

//      0   4 5   10   15  20  25  30  35  40
//  A         [----------------]
//  B   [---]     [----]   [--------]  [---]
// out: 5..10, 10..15, 15..20, 20..25, 25..30
