package day2

import (
	"advent-of-code/utils"
	"strconv"
)

type challangeResult = int

var gameScore = map[string]map[string]int{
	"Rock": {
		"Rock":     3,
		"Paper":    0,
		"Scissors": 6,
	},
	"Paper": {
		"Rock":     6,
		"Paper":    3,
		"Scissors": 0,
	},
	"Scissors": {
		"Rock":     0,
		"Paper":    6,
		"Scissors": 3,
	},
}

var figuresScore = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

var game2Score = map[byte]map[string]int{
	'X': {
		"Rock":     figuresScore["Scissors"],
		"Paper":    figuresScore["Rock"],
		"Scissors": figuresScore["Paper"],
	},
	'Y': {
		"Rock":     figuresScore["Rock"],
		"Paper":    figuresScore["Paper"],
		"Scissors": figuresScore["Scissors"],
	},
	'Z': {
		"Rock":     figuresScore["Paper"],
		"Paper":    figuresScore["Scissors"],
		"Scissors": figuresScore["Rock"],
	},
}

var Moves = map[byte]string{
	'X': "Rock",
	'Y': "Paper",
	'Z': "Scissors",

	'A': "Rock",
	'B': "Paper",
	'C': "Scissors",
}

func resolve1(rounds []rockPaperSisorsRound) int {
	score := 0

	for _, round := range rounds {
		elveMove := Moves[round.elveMove]
		playerMove := Moves[round.playerMove]
		score += figuresScore[playerMove]
		score += gameScore[playerMove][elveMove]
	}

	return score
}

func resolve2(rounds []rockPaperSisorsRound) int {
	score := 0
	resultScore := map[byte]int{
		'X': 0,
		'Y': 3,
		'Z': 6,
	}

	for _, round := range rounds {
		elveMove := Moves[round.elveMove]
		score += resultScore[round.playerMove]
		score += game2Score[round.playerMove][elveMove]
	}

	return score
}

func Solve() {
	solveChallange("day-2/files/test.txt")
	solveChallange("day-2/files/input.txt")
}

func solveChallange(file string) {
	input := ParseExcercise(file)
	inputRes, inputRes2 := resolve1(input), resolve2(input)
	logResults(file, inputRes, inputRes2)
}

func logResults(title string, res1 challangeResult, res2 challangeResult) {
	utils.FormatResults(title, strconv.Itoa(res1), strconv.Itoa(res2))
}
