package day4

import (
	utils "advent-of-code/utils"
	"strconv"
)

func resolve1(pairs challangeInput) challangeResult {

	repeatedPairs := 0
	for _, p := range pairs {
		firstPair, secondPair := p[0], p[1]
		if firstPair.from >= secondPair.from && firstPair.to <= secondPair.to {
			repeatedPairs++
		} else if secondPair.from >= firstPair.from && secondPair.to <= firstPair.to {
			repeatedPairs++
		}
	}
	return repeatedPairs
}

func resolve2(pairs challangeInput) challangeResult {
	repeatedPairs := 0
	for _, p := range pairs {
		firstPair, secondPair := p[0], p[1]
		if firstPair.from >= secondPair.from && firstPair.from <= secondPair.to || secondPair.from >= firstPair.from && secondPair.from <= firstPair.to {
			repeatedPairs++
		}
	}
	return repeatedPairs
}

func Solve() {
	solveChallange("day-4/files/test.txt")
	solveChallange("day-4/files/input.txt")
}

func solveChallange(file string) {
	input := ParseExcercise(file)
	inputRes, inputRes2 := resolve1(input), resolve2(input)
	logResults(file, inputRes, inputRes2)
}

func logResults(title string, res1 challangeResult, res2 challangeResult) {
	utils.FormatResults(title, strconv.Itoa(res1), strconv.Itoa(res2))
}
