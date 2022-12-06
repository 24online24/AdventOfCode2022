package day1

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

	// var maximumCalories [3]int = {0, 0, 0}

	maximumCalories := [3]int{0, 0, 0}
	currentCalories := 0
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		if lineText != "" {
			lineValue, err := strconv.Atoi(lineText)

			if err != nil {
				fmt.Println(err)
			}

			currentCalories += lineValue
		} else {
			// if currentCalories > maximumCalories[2] {
			// 	if currentCalories > maximumCalories[1] {
			// 		if currentCalories > maximumCalories[0] {
			// 			maximumCalories[2] = maximumCalories[1]
			// 			maximumCalories[1] = maximumCalories[0]
			// 			maximumCalories[0] = currentCalories
			// 		} else {
			// 			maximumCalories[2] = maximumCalories[1]
			// 			maximumCalories[1] = currentCalories
			// 		}
			// 	} else {
			// 		maximumCalories[2] = currentCalories
			// 	}
			// }
			// currentCalories = 0

			// -------------------------------------------------------------

			// if currentCalories <= maximumCalories[2] {
			// 	currentCalories = 0
			// 	continue
			// }
			// if currentCalories <= maximumCalories[1] {
			// 	maximumCalories[2] = currentCalories
			// 	currentCalories = 0
			// 	continue
			// }
			// if currentCalories <= maximumCalories[0] {
			// 	maximumCalories[2] = maximumCalories[1]
			// 	maximumCalories[1] = currentCalories
			// 	currentCalories = 0
			// 	continue
			// }
			// maximumCalories[2] = maximumCalories[1]
			// maximumCalories[1] = maximumCalories[0]
			// maximumCalories[0] = currentCalories
			// currentCalories = 0

			// -------------------------------------------------------------

			// switch {
			// case currentCalories < maximumCalories[2]:
			// case currentCalories < maximumCalories[1]:
			// 	maximumCalories[2] = currentCalories
			// case currentCalories < maximumCalories[0]:
			// 	maximumCalories[2] = maximumCalories[1]
			// 	maximumCalories[1] = currentCalories
			// default:
			// 	maximumCalories[2] = maximumCalories[1]
			// 	maximumCalories[1] = maximumCalories[0]
			// 	maximumCalories[0] = currentCalories
			// }
			// currentCalories = 0

			// -------------------------------------------------------------

			switch {
			case currentCalories > maximumCalories[0]:
				maximumCalories[2] = maximumCalories[1]
				maximumCalories[1] = maximumCalories[0]
				maximumCalories[0] = currentCalories
			case currentCalories > maximumCalories[1]:
				maximumCalories[2] = maximumCalories[1]
				maximumCalories[1] = currentCalories
			case currentCalories > maximumCalories[2]:
				maximumCalories[2] = currentCalories
			}
			currentCalories = 0
		}
	}
	inputFile.Close()

	fmt.Println(maximumCalories)
	sumOfMaximums := 0
	for _, element := range maximumCalories {
		sumOfMaximums += element
	}
	fmt.Println(sumOfMaximums)
}
