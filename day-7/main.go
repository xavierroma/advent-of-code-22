package day7

import (
	utils "advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func resolve1(input challangeInput) challangeResult {
	root := buildDirectoryTree(input)
	dirs := findDirectoriesWithSizelt(100000, root)
	size := 0
	for _, d := range dirs {
		size += d.size
	}
	return size
}

func buildDirectoryTree(commands []string) *Directory {

	root := Directory{"/", 0, nil, make([]*Directory, 0), make([]*File, 0)}
	currDir := &root
	for _, c := range commands {
		commandParts := strings.Split(c, " ")
		switch commandParts[0] {
		case "$":
			if c[2] == 'c' {
				dirOperation := commandParts[2]
				if dirOperation == "/" {
					currDir = &root
				} else if dirOperation == ".." {
					currDir = currDir.parentDirectory
				} else {
					currDir = findDirectory(dirOperation, currDir.directories)
				}
			}
			break
		case "dir":
			currDir.directories = append(currDir.directories, newDirectory(commandParts[1], currDir))
			break
		default: //File
			size, err := strconv.Atoi(commandParts[0])
			utils.CheckError(err)
			fileName := commandParts[1]
			currDir.files = append(currDir.files, &File{fileName, size})
			break
		}

	}
	computeDirectoriesSize(&root)
	return &root
}

func findDirectory(name string, dirs []*Directory) *Directory {
	for _, d := range dirs {
		if d.name == name {
			return d
		}
	}
	return nil
}

func newDirectory(name string, parent *Directory) *Directory {
	return &Directory{name, 0, parent, make([]*Directory, 0), make([]*File, 0)}
}

func computeDirectoriesSize(root *Directory) {

	for _, d := range root.directories {
		computeDirectoriesSize(d)
	}

	for _, d := range root.directories {
		root.size += d.size
	}
	for _, f := range root.files {
		root.size += f.size
	}

}

func findDirectoriesWithSizelt(lt int, root *Directory) []*Directory {
	dirs := make([]*Directory, 0)

	for _, d := range root.directories {
		matchingDirs := findDirectoriesWithSizelt(lt, d)
		for _, md := range matchingDirs {
			dirs = append(dirs, md)
		}
	}

	if root.size < lt {
		dirs = append(dirs, root)
	}
	return dirs
}

func findDirectoriesWithSizegte(gte int, root *Directory) []*Directory {
	dirs := make([]*Directory, 0)

	for _, d := range root.directories {
		matchingDirs := findDirectoriesWithSizegte(gte, d)
		for _, md := range matchingDirs {
			dirs = append(dirs, md)
		}
	}

	if root.size > gte {
		dirs = append(dirs, root)
	}
	return dirs
}
func debugTree(d *Directory, depth int) {
	padding := ""
	for i := 0; i < depth; i++ {
		padding += "  "
	}
	fmt.Printf("%s- %s (dir, size=%d)\n", padding, d.name, d.size)
	for _, f := range d.files {
		fmt.Printf("%s  - %s (file, size=%d)\n", padding, f.name, f.size)
	}
	for _, dir := range d.directories {
		debugTree(dir, depth+1)
	}
}

func resolve2(input challangeInput) challangeResult {
	totalSize := 70000000
	required := 30000000
	root := buildDirectoryTree(input)
	toFree := required - (totalSize - root.size)
	dirs := findDirectoriesWithSizegte(toFree, root)
	minDirSize := root.size
	for _, d := range dirs {
		if minDirSize > d.size {
			minDirSize = d.size
		}
	}
	return minDirSize
}

func Solve() {
	solveChallange("day-7/files/test.txt")
	solveChallange("day-7/files/input.txt")
}

func solveChallange(file string) {
	input := ParseExcercise(file)
	inputRes, inputRes2 := resolve1(input), resolve2(input)
	logResults(file, inputRes, inputRes2)
}

func logResults(title string, res1 challangeResult, res2 challangeResult) {
	utils.FormatResults(title, strconv.Itoa(res1), strconv.Itoa(res2))
}
