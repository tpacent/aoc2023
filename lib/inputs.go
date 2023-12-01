package lib

import (
	"bufio"
	"os"
)

func MustReadFile(path string) (lines []string) {
	file, err := os.Open(path)

	if err != nil {
		panic("cannot read input")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
