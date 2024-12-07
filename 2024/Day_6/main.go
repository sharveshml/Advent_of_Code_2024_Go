package main

import (
	"bufio"
	"fmt"
	"os"
)

func readData(path string) [][]rune {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var labGrid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		labGrid = append(labGrid, row)
	}

	return labGrid
}

func getDistinctPositions(labGrid [][]rune, startPos [2]int) int {
	guardDirection := 0

	guardRow := startPos[0]
	guardCol := startPos[1]

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	visitedMap := make(map[[2]int]bool)

	for {

		if guardRow < 0 || guardRow >= len(labGrid) || guardCol < 0 || guardCol >= len(labGrid[0]) {
			break
		}

		visitedMap[[2]int{guardRow, guardCol}] = true
		nextGuardRow := guardRow + directions[guardDirection][0]
		nextGuardCol := guardCol + directions[guardDirection][1]

		if nextGuardRow < 0 || nextGuardCol < 0 || nextGuardRow >= len(labGrid) || nextGuardCol >= len(labGrid[0]) {
			break
		}

		if labGrid[nextGuardRow][nextGuardCol] == '#' {
			guardDirection = (guardDirection + 1) % 4
			guardRow += directions[guardDirection][0]
			guardCol += directions[guardDirection][1]
		} else {
			guardRow = nextGuardRow
			guardCol = nextGuardCol
		}

	}

	return len(visitedMap)

}

func isCycle(labGrid [][]rune, startPos [2]int) bool {
	guardDirection := 0
	guardRow := startPos[0]
	guardCol := startPos[1]
	// debugGrid := make([][]rune, len(labGrid))
	// for i := range labGrid {
	// 	debugGrid[i] = make([]rune, len(labGrid[i]))
	// 	copy(debugGrid[i], labGrid[i])
	// }

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	visitedMap := make(map[[3]int]bool)

	for {
		if guardRow < 0 || guardRow >= len(labGrid) || guardCol < 0 || guardCol >= len(labGrid[0]) {
			break
		}

		// debugGrid[guardRow][guardCol] = 'X'
		if visitedMap[[3]int{guardRow, guardCol, guardDirection}] {

			// for i := range debugGrid {
			// 	for j := range debugGrid[i] {
			// 		fmt.Printf("%c", debugGrid[i][j])
			// 	}
			// 	fmt.Println()
			// }
			// fmt.Println()
			return true
		}

		visitedMap[[3]int{guardRow, guardCol, guardDirection}] = true
		nextGuardRow := guardRow + directions[guardDirection][0]
		nextGuardCol := guardCol + directions[guardDirection][1]

		if nextGuardRow < 0 || nextGuardCol < 0 || nextGuardRow >= len(labGrid) || nextGuardCol >= len(labGrid[0]) {
			break
		}

		if labGrid[nextGuardRow][nextGuardCol] == '#' {
			guardDirection = (guardDirection + 1) % 4
		} else {
			guardRow = nextGuardRow
			guardCol = nextGuardCol
		}
	}

	return false
}

func getMinObstacleCount(labGrid [][]rune, startPos [2]int) int {
	gridCopy := make([][]rune, len(labGrid))

	for i := range labGrid {
		gridCopy[i] = make([]rune, len(labGrid[i]))
		copy(gridCopy[i], labGrid[i])
	}

	obstacleCount := 0

	for i := range gridCopy {
		for j := range gridCopy[i] {
			if i == startPos[0] && j == startPos[1] {
				fmt.Printf("Start Position: %d %d\n", i, j)
			}
			if gridCopy[i][j] == '.' {
				gridCopy[i][j] = '#'
				if isCycle(gridCopy, startPos) {
					obstacleCount++
				}
				gridCopy[i][j] = '.'
			}
		}
	}
	return obstacleCount
}

func main() {
	labGrid := readData("inputs/input.txt")
	startPos := [2]int{}
	isFound := false

	for row := range labGrid {
		for col := range labGrid[row] {
			if labGrid[row][col] == '^' {
				isFound = true
				startPos[0] = row
				startPos[1] = col
				break
			}
		}
		if isFound {
			break
		}
	}

	fmt.Printf("Answer for Day 6 Part 1: %d\n", getDistinctPositions(labGrid, startPos))
	fmt.Printf("Answer for Day 6 Part 2: %d\n", getMinObstacleCount(labGrid, startPos))
}
