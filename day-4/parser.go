package day4

import (
	"advent-of-code/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type challangeInput = [][2]pair
type challangeResult = int
type pair struct {
	from int
	to   int
}

func ParseExcercise(file string) challangeInput {
	readFile, err := os.Open(file)

	utils.CheckError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var pairs challangeInput

	for fileScanner.Scan() {
		line := fileScanner.Text()
		rawPairs := strings.Split(line, ",")
		rawPair0 := strings.Split(rawPairs[0], "-")
		from0, err := strconv.Atoi(rawPair0[0])
		utils.CheckError(err)
		to0, err := strconv.Atoi(rawPair0[1])
		utils.CheckError(err)
		rawPair1 := strings.Split(rawPairs[1], "-")
		from1, err := strconv.Atoi(rawPair1[0])
		utils.CheckError(err)
		to1, err := strconv.Atoi(rawPair1[1])
		utils.CheckError(err)
		newPair := [2]pair{{from0, to0}, {from1, to1}}
		pairs = append(pairs, newPair)
	}
	readFile.Close()
	return pairs
}
