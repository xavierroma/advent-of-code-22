package day3

import (
	"advent-of-code/utils"
	"strconv"
)

type challangeResult = int

func resolve1(lines []string) challangeResult {

	repeatedChars := make([]byte, 0)
	for _, line := range lines {
		charMap := map[byte]bool{}
		for i := 0; i < len(line); i++ {
			if i < len(line)/2 {
				charMap[line[i]] = true
			} else {
				if charMap[line[i]] {
					repeatedChars = append(repeatedChars, line[i])
					break
				}
			}
		}
	}

	return sumChars(repeatedChars)
}

func resolve2(lines []string) challangeResult {
	repeatedChars := make([]byte, 0)

	for j := 0; j < len(lines); j += 3 {
		charMap := make([]map[byte]bool, 0)

		for v := 0; v < 3; v++ {
			line := lines[j+v]
			charMap = append(charMap, map[byte]bool{})
			for i := 0; i < len(line); i++ {
				charMap[v][line[i]] = true
				if v == 2 {
					if charMap[0][line[i]] && charMap[1][line[i]] {
						repeatedChars = append(repeatedChars, line[i])
						break
					}
				}
			}
		}
	}
	return sumChars(repeatedChars)
}

func sumChars(chars []byte) int {
	sum := 0
	for _, c := range chars {
		if int(byte('a')) > int(c) {
			sum += ((int(c) - int(byte('A'))) + 27)
		} else {
			sum += ((int(c) - int(byte('a'))) + 1)

		}
	}
	return sum
}

func Solve() {
	solveChallange("day-3/files/test.txt")
	solveChallange("day-3/files/input.txt")
}

func solveChallange(file string) {
	input := ParseExcercise(file)
	inputRes, inputRes2 := resolve1(input), resolve2(input)
	logResults(file, inputRes, inputRes2)
}

func logResults(title string, res1 challangeResult, res2 challangeResult) {
	utils.FormatResults(title, strconv.Itoa(res1), strconv.Itoa(res2))
}
