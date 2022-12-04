package day2

import (
	"advent-of-code/utils"
	"bufio"
	"os"
	"strings"
)

type challangeInput = []rockPaperSisorsRound
type rockPaperSisorsRound struct {
	elveMove   byte
	playerMove byte
}

func ParseExcercise(file string) challangeInput {
	readFile, err := os.Open(file)

	utils.CheckError(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var rounds []rockPaperSisorsRound = make([]rockPaperSisorsRound, 1)

	for fileScanner.Scan() {
		content := fileScanner.Text()
		value := strings.Split(content, " ")
		rounds = append(rounds, rockPaperSisorsRound{
			elveMove:   byte(value[0][0]),
			playerMove: byte(value[1][0]),
		})
	}
	readFile.Close()
	return rounds
}
