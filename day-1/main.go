package day1

import (
	"advent-of-code/utils"
	"sort"
	"strconv"
)

type challangeResult = int

func resolve1(elvesCalories [][]int) int {
	max := 0
	for _, elve := range elvesCalories {
		elveCaloriesSum := 0
		for _, calories := range elve {
			elveCaloriesSum += calories
		}
		if elveCaloriesSum > max {
			max = elveCaloriesSum
		}
	}

	return max
}

func resolve2(elvesCalories [][]int) int {
	elveSnacks := sumAndFlat(elvesCalories)
	sort.Slice(elveSnacks, func(i, j int) bool { return elveSnacks[i] > elveSnacks[j] })

	return elveSnacks[0] + elveSnacks[1] + elveSnacks[2]
}

func sumAndFlat(elvesCalories [][]int) []int {
	flatSum := make([]int, 0)
	for _, elve := range elvesCalories {
		elveCaloriesSum := 0
		for _, calories := range elve {
			elveCaloriesSum += calories
		}
		flatSum = append(flatSum, elveCaloriesSum)
	}
	return flatSum
}

func Solve() {
	solveChallange("day-1/files/test.txt")
	solveChallange("day-1/files/input.txt")
}

func solveChallange(file string) {
	input := ParseExcercise(file)
	inputRes, inputRes2 := resolve1(input), resolve2(input)
	logResults(file, inputRes, inputRes2)
}

func logResults(title string, res1 challangeResult, res2 challangeResult) {
	utils.FormatResults(title, strconv.Itoa(res1), strconv.Itoa(res2))
}
