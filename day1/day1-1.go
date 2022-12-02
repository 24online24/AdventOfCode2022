package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("data.in")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(inputFile)

	fileScanner.Split(bufio.ScanLines)

	maximumSum := 0
	currentSum := 0
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		if lineText != "" {
			lineValue, err := strconv.Atoi(lineText)

			if err != nil {
				fmt.Println(err)
			}

			currentSum += lineValue
		} else {
			if currentSum > maximumSum {
				maximumSum = currentSum
			}
			// fmt.Println(currentSum)
			currentSum = 0
		}
	}
	inputFile.Close()

	// fmt.Println()
	fmt.Println(maximumSum)
}
