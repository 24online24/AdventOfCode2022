package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	matrix := [][]rune{}
	iStart, jStart, iEnd, jEnd := -1, -1, -1, -1
	read(&matrix, &iStart, &jStart, &iEnd, &jEnd)
	matrix[iStart][jStart] = 'a'
	matrix[iEnd][jEnd] = 'z'
	road := lee(matrix, iStart, jStart)
	fmt.Println(road[iEnd][jEnd])

}

func read(matrix *[][]rune, iStart *int, jStart *int, iEnd *int, jEnd *int) {
	inputFile, err := os.Open("data.in")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		matrixLine := []rune{}
		for _, chr := range lineText {
			matrixLine = append(matrixLine, chr)
		}
		*matrix = append(*matrix, matrixLine)
	}
	for i, line := range *matrix {
		for j, element := range line {
			if element == 'S' {
				*iStart, *jStart = i, j
			} else if element == 'E' {
				*iEnd, *jEnd = i, j
			}
		}
	}
}

func lee(matrix [][]rune, iStart int, jStart int) [][]int {
	type coordinates struct {
		x int
		y int
	}

	directions := [4]coordinates{
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: -1, y: 0},
	}

	coordinatesQueue := []coordinates{}
	road := [][]int{}
	for _, line := range matrix {
		newLine := []int{}
		for range line {
			newLine = append(newLine, 0)
		}
		road = append(road, newLine)
	}
	coordinatesQueue = append(coordinatesQueue, coordinates{x: jStart, y: iStart})
	for len(coordinatesQueue) > 0 {
		i := coordinatesQueue[0].y
		j := coordinatesQueue[0].x
		coordinatesQueue = coordinatesQueue[1:]
		for _, direction := range directions {
			next_i := i + direction.y
			next_j := j + direction.x
			if viable(matrix, next_i, next_j) && matrix[next_i][next_j] <= matrix[i][j]+1 && road[next_i][next_j] == 0 {
				road[next_i][next_j] = road[i][j] + 1
				coordinatesQueue = append(coordinatesQueue, coordinates{x: next_j, y: next_i})
			}

		}
	}
	return road
}

func viable(matrix [][]rune, i int, j int) bool {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
		return false
	}
	return true
}
