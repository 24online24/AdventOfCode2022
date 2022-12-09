package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("data.in")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	containedPairs := 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
		// fmt.Println(text)
		textSeparated := strings.FieldsFunc(text, separatorsFunction)
		sectionIDs := []int{}
		for _, IDstring := range textSeparated {
			ID, err := strconv.Atoi(IDstring)

			if err != nil {
				fmt.Println(err)
			}

			sectionIDs = append(sectionIDs, ID)
		}
		if contains(sectionIDs[0], sectionIDs[1], sectionIDs[2], sectionIDs[3]) {
			fmt.Println(sectionIDs)
			containedPairs++
		}
	}
	fmt.Println(containedPairs)
}

func separatorsFunction(character rune) bool {
	return character == ',' || character == '-'
}

func contains(start1 int, end1 int, start2 int, end2 int) bool {
	return (start1 <= start2 && start2 <= end1) || (start2 <= start1 && start1 <= end2)
}
