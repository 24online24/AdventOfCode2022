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

	sumPriorities := 0
	for fileScanner.Scan() {
		rucksack := fileScanner.Text()
		firstHalf := rucksack[:len(rucksack)/2]
		secondHalf := rucksack[len(rucksack)/2:]
		found := false
		for _, itemTypeFirst := range firstHalf {
			for _, itemTypeSecond := range secondHalf {
				if itemTypeFirst == itemTypeSecond {
					if itemTypeFirst >= 'a' {
						itemTypeFirst = itemTypeFirst - 'a' + 1
					} else {
						itemTypeFirst = itemTypeFirst - 'A' + 27
					}
					sumPriorities += int(itemTypeFirst)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	fmt.Println(sumPriorities)
}
