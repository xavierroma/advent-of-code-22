package day5

import (
	utils "advent-of-code/utils"
	"fmt"
)

func resolve1(input *challangeInput) challangeResult {

	for _, move := range input.moves {
		for i := 0; i < move.moveCount; i++ {
			var last byte
			input.stack[move.fromStack-1], last = pop(input.stack[move.fromStack-1])
			input.stack[move.toStack-1] = append(input.stack[move.toStack-1], last)
		}
	}
	res := ""
	for _, stack := range input.stack {
		res += string(getLast(stack))
	}

	return res
}

func getLast[T any](arr []T) T {
	return arr[len(arr)-1]
}

func pop[T any](arr []T) ([]T, T) {
	last := getLast(arr)
	return arr[0 : len(arr)-1], last
}

func resolve2(input *challangeInput) challangeResult {
	for _, move := range input.moves {
		var popped []byte
		input.stack[move.fromStack-1], popped = popN(input.stack[move.fromStack-1], move.moveCount)
		input.stack[move.toStack-1] = appendN(input.stack[move.toStack-1], popped)
	}
	res := ""
	for _, stack := range input.stack {
		res += string(getLast(stack))
	}

	return res
}

func printStatus(input *challangeInput) {
	maxLen := 0
	for i := 0; i < len(input.stack); i++ {
		if maxLen < len(input.stack[i]) {
			maxLen = len(input.stack[i])
		}
	}

	for i := maxLen - 1; i >= 0; i-- {
		for _, stack := range input.stack {
			if (len(stack) - 1) < i {
				fmt.Print("    ")
			} else {
				fmt.Print(" [")
				fmt.Print(string(stack[i]))
				fmt.Print("]")
			}
		}
		fmt.Println("")
	}
	for i := 0; i < len(input.stack); i++ {
		fmt.Print("  ")
		fmt.Print(i + 1)
		fmt.Print(" ")
	}
	fmt.Println("")
	fmt.Println("")
}

func popN[T any](arr []T, n int) ([]T, []T) {
	return arr[:len(arr)-n], arr[len(arr)-n:]
}

func appendN[T any](destinationSlice []T, toAppend []T) []T {
	for _, el := range toAppend {
		destinationSlice = append(destinationSlice, el)
	}
	return destinationSlice
}

func Solve() {
	solveChallange("day-5/files/test.txt")
	solveChallange("day-5/files/input.txt")
}

func solveChallange(file string) {
	input := ParseExcercise(file)
	input2 := ParseExcercise(file)
	inputRes, inputRes2 := resolve1(&input), resolve2(&input2)
	logResults(file, inputRes, inputRes2)
}

func logResults(title string, res1 challangeResult, res2 challangeResult) {
	utils.FormatResults(title, res1, res2)
}
