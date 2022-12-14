package day8

import (
	"advent-of-code/utils"
	"bufio"
	"os"
)

type challangeInput = [][]int

type challangeResult = int

func ParseExcercise(file string) challangeInput {
	readFile, err := os.Open(file)

	utils.CheckError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	input := [][]int{}
	for fileScanner.Scan() {
		inputLine := make([]int, 0)
		line := fileScanner.Text()
		for _, v := range line {
			inputLine = append(inputLine, int(v))
		}
		input = append(input, inputLine)
	}
	fileScanner.Scan()
	readFile.Close()
	return input
}
