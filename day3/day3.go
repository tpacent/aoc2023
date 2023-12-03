package day3

import (
	"aoc2023/lib"
	"regexp"
	"strconv"
	"strings"
)

var PartNumRe = regexp.MustCompile(`\d+`)
var GearsRe = regexp.MustCompile(`\*`)

// IsSymbol matches anything not a dot (.) or a digit
func IsSymbol(r rune) bool {
	if r == '.' {
		return false
	}

	if r >= '0' && r <= '9' {
		return false
	}

	return true
}

// IsPart tests if any of the chunks contain symbol
func IsPart(chunks []string) bool {
	for _, chunk := range chunks {
		if strings.ContainsFunc(chunk, IsSymbol) {
			return true
		}
	}

	return false
}

func FindPartNumbers(lines []string) (parts []int) {
	for index, line := range lines {
		matches := PartNumRe.FindAllStringIndex(line, -1)
		for _, match := range matches {
			fromIndex := max(0, match[0]-1)
			uptoIndex := min(len(line)-1, match[1]+1)
			chunks := make([]string, 0, 3)
			chunks = append(chunks, line[fromIndex:uptoIndex])
			if index > 0 {
				chunks = append(chunks, lines[index-1][fromIndex:uptoIndex])
			}
			if index < len(lines)-1 {
				chunks = append(chunks, lines[index+1][fromIndex:uptoIndex])
			}
			if IsPart(chunks) {
				if n, err := strconv.Atoi(line[match[0]:match[1]]); err == nil {
					parts = append(parts, n)
				} else {
					panic("unreachable")
				}
			}
		}
	}

	return
}

func ExtractNumbersAroundGear(lines []string, gearCol, gearRow int) (partNumbers []int) {
	fromIndex := max(0, gearCol-1)
	uptoIndex := min(len(lines[0])-1, gearCol+2)
	vicinity := lines[max(0, gearRow-1):min(gearRow+2, len(lines))]
	for _, line := range vicinity {
		matches := PartNumRe.FindAllStringIndex(line, -1)
		for _, match := range matches {
			if intersects(match[0], match[1], fromIndex, uptoIndex) {
				if n, err := strconv.Atoi(line[match[0]:match[1]]); err == nil {
					partNumbers = append(partNumbers, n)
				} else {
					panic("unreachable")
				}
			}
		}
	}

	return
}

func FindGearsRatios(lines []string) (ratios []int) {
	for index, line := range lines {
		matches := GearsRe.FindAllStringIndex(line, -1) // [[3, 4]]
		for _, match := range matches {
			numbers := ExtractNumbersAroundGear(lines, match[0], index)
			if len(numbers) == 2 {
				ratios = append(ratios, lib.Mul(numbers...))
			}
		}
	}

	return
}

func intersects(aFrom, aUpto, bFrom, bUpto int) bool {
	if aUpto <= bFrom {
		return false
	}

	if aFrom >= bUpto {
		return false
	}

	return true
}
