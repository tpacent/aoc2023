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
	return &RangeMapper{mappings: m, name: name}
}

type RangeMapper struct {
	mappings []*Mapping
	name     string
}

func (rm *RangeMapper) Map(value int) int {
	for _, mapper := range rm.mappings {
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
