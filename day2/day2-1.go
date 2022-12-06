package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	type rules struct {
		lose byte
		win  byte
	}

	guide := []rules{
		{
			lose: 'A',
			win:  'Y',
		},
		{
			lose: 'B',
			win:  'Z',
		},
		{
			lose: 'C',
			win:  'X',
		},
	}

	inputFile, err := os.Open("data.in")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	score := 0

	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		p1 := lineText[0]
		p2 := lineText[2]
		if guide[p1-'A'].win == p2 {
			score += 6
		} else if p1 == p2-'X'+'A' {
			score += 3
		}
		score += int(p2 - 'X' + 1)
	}
	fmt.Println(score)
}
