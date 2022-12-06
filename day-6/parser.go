package day6

import (
	"advent-of-code/utils"
	"bufio"
	"os"
)

type challangeInput = string
type challangeResult = int

func ParseExcercise(file string) challangeInput {
	readFile, err := os.Open(file)

	utils.CheckError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	input := fileScanner.Text()
	readFile.Close()
	return input
}
