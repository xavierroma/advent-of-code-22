package day5

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type challangeInput = struct {
	stack [][]byte
	moves []struct {
		moveCount int
		fromStack int
		toStack   int
	}
}
type challangeResult = string

func ParseExcercise(file string) challangeInput {
	readFile, err := os.Open(file)

	utils.CheckError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var input challangeInput

	var strStacks []string
	var strLen int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line[1] == '1' {
			strLen, err = strconv.Atoi(string(line[len(line)-2]))
			break
		}
		strStacks = append(strStacks, line)
	}
	var stacks [][]byte = make([][]byte, strLen)

	for i := 0; i < len(stacks); i++ {
		stacks[i] = make([]byte, 0)
	}
	for i := len(strStacks) - 1; i >= 0; i-- {
		stackNum := 0
		for j := 1; j < len(strStacks[i]); j += 4 {
			if strStacks[i][j] != ' ' && stackNum < len(stacks) {
				stacks[stackNum] = append(stacks[stackNum], strStacks[i][j])
			}
			stackNum++
		}
	}

	var moves []struct {
		moveCount int
		fromStack int
		toStack   int
	}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}
		var moveCount, fromStack, toStack int
		fmt.Sscanf(line, "move %d from %d to %d", &moveCount, &fromStack, &toStack)
		moves = append(moves, struct {
			moveCount int
			fromStack int
			toStack   int
		}{moveCount, fromStack, toStack})
	}
	input.stack = stacks
	input.moves = moves
	readFile.Close()
	return input
}
