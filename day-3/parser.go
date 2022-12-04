package day3

import (
	"advent-of-code/utils"
	"bufio"
	"os"
)

type challangeInput = []string

func ParseExcercise(file string) challangeInput {
	readFile, err := os.Open(file)

	utils.CheckError(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var lines []string = make([]string, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)

	}
	readFile.Close()
	return lines
}
