package day6

import (
	utils "advent-of-code/utils"
	"strconv"
)

func resolve1(input challangeInput) challangeResult {
	queue := CircularQueue[byte](4)
	for i := 4; i < len(input); i++ {
		if areAllDifferent(queue.queue) {
			return i
		}
		queue.enqueue(input[i])
	}
	return -1
}

func areAllDifferent(arr []byte) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return false
			}
		}
	}
	return true
}

func resolve2(input challangeInput) challangeResult {
	queue := CircularQueue[byte](14)
	for i := 14; i < len(input); i++ {
		if areAllDifferent(queue.queue) {
			return i
		}
		queue.enqueue(input[i])
	}
	return -1
}

func Solve() {
	solveChallange("day-6/files/test.txt")
	solveChallange("day-6/files/input.txt")
}

func solveChallange(file string) {
	input := ParseExcercise(file)
	inputRes, inputRes2 := resolve1(input), resolve2(input)
	logResults(file, inputRes, inputRes2)
}

func logResults(title string, res1 challangeResult, res2 challangeResult) {
	utils.FormatResults(title, strconv.Itoa(res1), strconv.Itoa(res2))
}
