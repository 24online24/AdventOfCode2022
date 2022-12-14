package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	cave := [200][800]rune{}
	read(&cave)
	// printCave(cave, 0, 494, 9, 503)
	// fmt.Println(yFloor)
	spaceLeft := true
	units := -1
	for spaceLeft {
		units++
		spaceLeft = nextSand(&cave)
	}
	fmt.Println(units)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func initCave(cave *[200][800]rune) {
	for i, line := range cave {
		for j := range line {
			cave[i][j] = '.'
		}
	}
	cave[0][500] = '+'
}

func printCave(cave [200][800]rune, iStart int, jStart int, iEnd int, jEnd int) {
	for i := iStart; i <= iEnd; i++ {
		for j := jStart; j <= jEnd; j++ {
			fmt.Print(string(cave[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func read(cave *[200][800]rune) {
	initCave(cave)
	inputFile, err := os.Open("data.in")
	handleError(err)
	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	yFloor := -1
	for fileScanner.Scan() {
		path := fileScanner.Text()
		coordinates := strings.Split(path, " -> ")
		splitCoords := strings.Split(coordinates[0], ",")
		prev_x, err := strconv.Atoi(splitCoords[0])
		handleError(err)
		prev_y, err := strconv.Atoi(splitCoords[1])
		handleError(err)
		if prev_y > yFloor {
			yFloor = prev_y
		}

		for index := 1; index < len(coordinates); index++ {
			splitCoords := strings.Split(coordinates[index], ",")
			x, err := strconv.Atoi(splitCoords[0])
			handleError(err)
			y, err := strconv.Atoi(splitCoords[1])
			handleError(err)
			// fmt.Print(x, y, " ")

			// fmt.Println("delta x:", x-prev_x)
			// fmt.Println("delta y:", y-prev_y)

			if y > yFloor {
				yFloor = y
			}

			x1, x2 := -1, -1
			if prev_x <= x {
				x1, x2 = prev_x, x
			} else {
				x2, x1 = prev_x, x
			}

			y1, y2 := -1, -1
			if prev_y <= y {
				y1, y2 = prev_y, y
			} else {
				y2, y1 = prev_y, y
			}

			for i := y1; i <= y2; i++ {
				for j := x1; j <= x2; j++ {
					cave[i][j] = '#'
				}
			}

			prev_x = x
			prev_y = y
		}
		// fmt.Println()
	}
	inputFile.Close()
	yFloor += 2
	for x := 0; x < 800; x++ {
		cave[yFloor][x] = '#'
	}
}

func nextSand(cave *[200][800]rune) bool {
	xSand, ySand := 500, 0
	cave[ySand][xSand] = 'o'
	canContinue := true
	for canContinue {
		switch {
		case cave[ySand+1][xSand] == '.':
			{
				cave[ySand][xSand] = '.'
				ySand++
				cave[ySand][xSand] = 'o'
			}
		case cave[ySand+1][xSand-1] == '.':
			{
				cave[ySand][xSand] = '.'
				ySand++
				xSand--
				cave[ySand][xSand] = 'o'
			}
		case cave[ySand+1][xSand+1] == '.':
			{
				cave[ySand][xSand] = '.'
				ySand++
				xSand++
				cave[ySand][xSand] = 'o'
			}
		default:
			canContinue = false
		}
		if ySand == 0 {
			return false
		}
	}
	// printCave(*cave, 0, 490, 11, 510)
	// time.Sleep(time.Millisecond * 500)
	return true
}
