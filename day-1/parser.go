package day1

import (
	"advent-of-code/utils"
	"bufio"
	"os"
	"strconv"
)

type challangeInput = [][]int

func ParseExcercise(file string) challangeInput {
	readFile, err := os.Open(file)

	utils.CheckError(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var elves [][]int = make([][]int, 1)
	elveNum := 0

	for fileScanner.Scan() {
		content := fileScanner.Text()
		if len(content) > 0 {
			value, err := strconv.ParseInt(content, 10, 32)
			utils.CheckError(err)
			elves[elveNum] = append(elves[elveNum], int(value))
		} else {
			elves = append(elves, make([]int, 0))
			elveNum++
		}
	}
	readFile.Close()
	return elves
}
