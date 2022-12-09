package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	name     string
	children []string
	size     int
	parent   *directory
}

var filesystem []directory

func main() {

	inputFile, err := os.Open("testdata.in")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	filesystem = append(filesystem, directory{name: "/"})
	directoryIndex := 0
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		lineTextSeparated := strings.Split(lineText, " ")
		switch lineTextSeparated[1] {
		case "cd":
			directoryIndex = changeDirectory(lineTextSeparated[2])
		case "ls":
			lineInBufferSeparated := listChildren(directoryIndex, fileScanner)
			fmt.Println(filesystem)
			if len(lineInBufferSeparated) > 0 {
				directoryIndex = changeDirectory(lineInBufferSeparated[2])
			}
		}
	}
	fmt.Println(filesystem)
	for index, dir := range filesystem {
		fmt.Println(index, dir)
	}
}

func changeDirectory(directoryName string) int {
	for index, dir := range filesystem {
		if dir.name == directoryName {
			return index
		}
	}
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
			filesystem = append(filesystem, directory{name: lineTextSeparated[1], parent: &filesystem[directoryIndex]})
			filesystem[directoryIndex].children = append(filesystem[directoryIndex].children, "")
		default:
			sizeUpdate, err := strconv.Atoi(lineTextSeparated[0])

			if err != nil {
				panic("Something went wrong with the size transformation to int!")
			}

			directoryToUpdate := &filesystem[directoryIndex]
			for directoryToUpdate != nil {
				(*directoryToUpdate).size += sizeUpdate
				fmt.Println("Updated", directoryToUpdate, "with file", lineTextSeparated, ". Size is now", directoryToUpdate.size)
				directoryToUpdate = directoryToUpdate.parent
			}
		}
	}
	var emptyArr []string
	return emptyArr
}
