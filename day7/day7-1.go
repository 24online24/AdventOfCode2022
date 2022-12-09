package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println(lineTextSeparated)
		switch lineTextSeparated[1] {
		case "cd":
			directoryIndex = changeDirectory(lineTextSeparated[2])
			fmt.Println("\t\t\t\t", filesystem[directoryIndex])
		case "ls":
			lineInBuffer := listChildren(directoryIndex, fileScanner)
			if lineInBuffer != "" {
				lineInBufferSeparated := strings.Split(lineInBuffer, " ")
				directoryIndex = changeDirectory(lineInBufferSeparated[2])
			}
		}
	}
	fmt.Println(filesystem)
}
func changeDirectory(directoryName string) int {
	for index, dir := range filesystem {
		if dir.name == directoryName {
			return index
		}
	}
	return -1
}

func listChildren(directoryIndex int, fileScanner *bufio.Scanner) string {
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		if lineText[0] == '$' {
			return lineText
		}
	}
}
