package lib

import (
	"bufio"
	"os"
)

func MustReadFileBytes(path string) (lines [][]byte) {
	file, err := os.Open(path)

	if err != nil {
		panic("cannot read input")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// gotcha here: without copy, lines are ordered randomly
		lines = append(lines, append([]byte(nil), scanner.Bytes()...))
	}

	return
}

func MustReadFile(path string) (lines []string) {
	file, err := os.Open(path)

	if err != nil {
		panic("cannot read input")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}
