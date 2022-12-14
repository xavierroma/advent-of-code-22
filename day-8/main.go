package day8

import (
	utils "advent-of-code/utils"
	"fmt"
	"strconv"
)

func resolve1(input challangeInput) challangeResult {
	visibility := computeVisibility(input)
	count := 0
	for i := 0; i < len(visibility); i++ {
		for j := 0; j < len(visibility[i]); j++ {
			if visibility[i][j] > 0 {
				count++
			}
		}
	}
	return count
}

func computeVisibility(input challangeInput) [][]int {

	visibility := make([][]int, len(input))

	for i := 0; i < len(input); i++ {
		visibility[i] = make([]int, len(input[0]))
		for j := 0; j < len(input[i]); j++ {
			visibility[i][j] = 4
		}
	}

	for i := 0; i < len(input); i++ {
		max := -1
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] <= max {
				visibility[i][j]--
			} else {
				max = input[i][j]
			}
		}
		max = -1
		for j := len(input[i]) - 1; j >= 0; j-- {
			if input[i][j] <= max {
				visibility[i][j]--
			} else {
				max = input[i][j]
			}
		}
	}

	for i := 0; i < len(input); i++ {
		max := -1
		for j := 0; j < len(input[i]); j++ {
			if input[j][i] <= max {
				visibility[j][i]--
			} else {
				max = input[j][i]
			}
		}
		max = -1
		for j := len(input[i]) - 1; j >= 0; j-- {
			if input[j][i] <= max {
				visibility[j][i]--
			} else {
				max = input[j][i]
			}
		}
	}

	return visibility
}

func printVisibility(visibility challangeInput) {
	for i := 0; i < len(visibility); i++ {
		for j := 0; j < len(visibility[i]); j++ {
			fmt.Print(visibility[i][j])
		}
		fmt.Println()
	}
}

func resolve2(input challangeInput) challangeResult {
	maxScenicScore := 0
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[0])-1; j++ {
			scenicScore := computeScenicScoreAt(i, j, input)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return maxScenicScore
}

func computeScenicScoreAt(i int, j int, input challangeInput) int {
	currTreeHeight := input[i][j]

	left := look(invertArray(input[i][:j]), currTreeHeight)
	right := look(input[i][j+1:], currTreeHeight)

	col := columnToArray(input, j)

	top := look(invertArray(col[:i]), currTreeHeight)
	bottom := look(col[i+1:], currTreeHeight)

	return left * right * top * bottom
}

func look(arr []int, treeHeight int) int {
	tallerThan := 1
	for _, v := range arr {
		if treeHeight <= v {
			return tallerThan
		}
		tallerThan++
	}
	return tallerThan - 1
}

func invertArray(arr []int) []int {
	arrClone := make([]int, len(arr))
	copy(arrClone, arr)
	for i, j := 0, len(arrClone)-1; i < j; i, j = i+1, j-1 {
		arrClone[i], arrClone[j] = arrClone[j], arrClone[i]
	}
	return arrClone
}

func columnToArray(input challangeInput, column int) []int {
	col := make([]int, len(input[0]))
	for c := range input {
		col[c] = input[c][column]
	}
	return col
}

func Solve() {
	solveChallange("day-8/files/test.txt")
	solveChallange("day-8/files/input.txt")
}

func solveChallange(file string) {
	input := ParseExcercise(file)
	inputRes, inputRes2 := resolve1(input), resolve2(input)
	logResults(file, inputRes, inputRes2)
}

func logResults(title string, res1 challangeResult, res2 challangeResult) {
	utils.FormatResults(title, strconv.Itoa(res1), strconv.Itoa(res2))
}
