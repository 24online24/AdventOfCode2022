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
		elf1 := ""
		elf2 := ""
		elf3 := ""
		elf1 = fileScanner.Text()
		fileScanner.Scan()
		elf2 = fileScanner.Text()
		fileScanner.Scan()
		elf3 = fileScanner.Text()
		found := false
		for _, itemType1 := range elf1 {
			for _, itemType2 := range elf2 {
				if itemType1 == itemType2 {
					for _, itemType3 := range elf3 {
						if itemType1 == itemType3 {
							if itemType1 >= 'a' {
								itemType1 = itemType1 - 'a' + 1
							} else {
								itemType1 = itemType1 - 'A' + 27
							}
							sumPriorities += int(itemType1)
							found = true
							break
						}
					}
					if found {
						break
					}
				}
			}
			if found {
				break
			}
		}
	}
	fmt.Println(sumPriorities)
}
