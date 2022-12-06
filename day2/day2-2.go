package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("data.in")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	score := 0
	var valuePlayed int
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		p1 := int(lineText[0] - 'A')
		p2 := lineText[2]
		switch p2 {
		case 'X':
			valuePlayed = (p1-1)%3 + 1
		case 'Y':
			valuePlayed = p1 + 1
		case 'Z':
			valuePlayed = (p1+1)%3 + 1
		}
		score += valuePlayed
		score += int(p2-'X') * 3
	}
	fmt.Println(score)
}
