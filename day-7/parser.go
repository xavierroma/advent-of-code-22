package day7

import (
	"advent-of-code/utils"
	"bufio"
	"os"
)

type challangeInput = []string

type challangeResult = int

func ParseExcercise(file string) challangeInput {
	readFile, err := os.Open(file)

	utils.CheckError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	input := []string{}
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	fileScanner.Scan()
	readFile.Close()
	return input
}
