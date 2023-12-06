package day6

import (
	"aoc2023/lib"
	"strings"
)

type RaceInfo struct {
	Time     int
	Distance int
}

func ParseRaces(input []string) (races []*RaceInfo) {
	times := strings.Fields(input[0])[1:]
	dists := strings.Fields(input[1])[1:]

	for k := 0; k < len(times); k++ {
		races = append(races, &RaceInfo{
			Time:     lib.AsInt(times[k]),
			Distance: lib.AsInt(dists[k]),
		})
	}

	return
}

func CountWaysToWin(race *RaceInfo) (count int) {
	for hold := 1; hold < race.Time; hold++ {
		speed := hold
		moveTime := race.Time - hold
		dist := moveTime * speed
		if dist > race.Distance {
			count++
		}
	}

	return
}
