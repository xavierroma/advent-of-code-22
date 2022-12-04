package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func FormatResults(title string, part1Result string, part2Result string) {
	fmt.Printf("\n--- %s ---\n", title)
	fmt.Printf("Part 1: %s\n", part1Result)
	fmt.Printf("Part 2: %s\n", part2Result)
	fmt.Printf("--------%s\n\n", nDashes(len(title)))
}

func nDashes(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += "-"
	}
	return str
}

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
