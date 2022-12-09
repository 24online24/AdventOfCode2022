package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	name     string
	children []string
	size     int
	parent   int
}

var filesystem []directory

func main() {

	inputFile, err := os.Open("data.in")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	filesystem = append(filesystem, directory{name: "/", parent: -1})
	directoryIndex := -1
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		lineTextSeparated := strings.Split(lineText, " ")
		switch lineTextSeparated[1] {
		case "cd":
			directoryIndex = changeDirectory(lineTextSeparated[2], directoryIndex)
		case "ls":
			lineInBufferSeparated := listChildren(directoryIndex, fileScanner)
			// fmt.Println(filesystem)
			if len(lineInBufferSeparated) > 0 {
				directoryIndex = changeDirectory(lineInBufferSeparated[2], directoryIndex)
			}
		}
	}
	// fmt.Println(filesystem)
	spaceNeeded := 70000000 - filesystem[0].size
	sizeToDelete := math.MaxInt
	for _, dir := range filesystem {
		if dir.size > spaceNeeded {
			if dir.size < sizeToDelete {
				sizeToDelete = dir.size
			}
		}
	}
	fmt.Println(sizeToDelete)
}

func changeDirectory(directoryName string, directoryIndex int) int {
	if directoryName == ".." {
		return filesystem[directoryIndex].parent
	}
	for index, dir := range filesystem {
		if dir.name == directoryName && dir.parent == directoryIndex {
			return index
		}
	}
	fmt.Println(filesystem[directoryIndex])
	return -1
}

func listChildren(directoryIndex int, fileScanner *bufio.Scanner) []string {
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		lineTextSeparated := strings.Split(lineText, " ")
		switch lineTextSeparated[0] {
		case "$":
			return lineTextSeparated
		case "dir":
			filesystem = append(filesystem, directory{name: lineTextSeparated[1], parent: directoryIndex})
			filesystem[directoryIndex].children = append(filesystem[directoryIndex].children, "")
		default:
			sizeUpdate, err := strconv.Atoi(lineTextSeparated[0])

			if err != nil {
				panic("Something went wrong with the size transformation to int!")
			}

			indexToUpdate := directoryIndex

			for indexToUpdate != -1 {
				filesystem[indexToUpdate].size += sizeUpdate
				indexToUpdate = filesystem[indexToUpdate].parent
			}

			// directoryToUpdate := &filesystem[directoryIndex]
			// for directoryToUpdate != nil {
			// 	(*directoryToUpdate).size += sizeUpdate
			// 	fmt.Println("Updated", directoryToUpdate, "with file", lineTextSeparated, ". Size is now", directoryToUpdate.size)
			// 	directoryToUpdate = directoryToUpdate.parent
			// }
		}
	}
	var emptyArr []string
	return emptyArr
}
