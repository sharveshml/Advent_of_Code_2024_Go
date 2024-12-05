package main

import (
	"fmt"
	"os"
	"strings"
)

func checkRightUpperDiagonal(grid [][]rune, i int, j int) int {
	count := 0
	if i-2 < 0 || j+2 >= len(grid) {
		return 0
	}

	if grid[i][j] == 'M' && grid[i-1][j+1] == 'A' && grid[i-2][j+2] == 'S' {
		count += 1
	}

	return count
}

func checkRightBottomDiagonal(grid [][]rune, i int, j int) int {
	count := 0

	if i+2 >= len(grid) || j+2 >= len(grid) {
		return 0
	}

	if grid[i][j] == 'M' && grid[i+1][j+1] == 'A' && grid[i+2][j+2] == 'S' {
		count += 1
	}

	return count
}

func checkLeftUpperDiagonal(grid [][]rune, i int, j int) int {
	count := 0

	if i-2 < 0 || j-2 < 0 {
		return 0
	}

	if grid[i][j] == 'M' && grid[i-1][j-1] == 'A' && grid[i-2][j-2] == 'S' {
		count += 1
	}

	return count
}

func checkLeftBottomDiagonal(grid [][]rune, i int, j int) int {
	count := 0

	if i+2 >= len(grid) || j-2 < 0 {
		return 0
	}

	if grid[i][j] == 'M' && grid[i+1][j-1] == 'A' && grid[i+2][j-2] == 'S' {
		count += 1
	}

	return count
}

func checkRight(grid [][]rune, i int, j int) int {
	count := 0

	if j+2 >= len(grid[i]) {
		return 0
	}

	if grid[i][j] == 'M' && grid[i][j+1] == 'A' && grid[i][j+2] == 'S' {
		count += 1
	}

	return count
}

func checkLeft(grid [][]rune, i int, j int) int {
	count := 0

	if j-2 < 0 {
		return 0
	}

	if grid[i][j] == 'M' && grid[i][j-1] == 'A' && grid[i][j-2] == 'S' {
		count += 1
	}

	return count
}

func checkTop(grid [][]rune, i int, j int) int {

	if i-2 < 0 {
		return 0
	}

	if grid[i][j] == 'M' && grid[i-1][j] == 'A' && grid[i-2][j] == 'S' {
		return 1
	}

	return 0
}

func checkBottom(grid [][]rune, i int, j int) int {

	if i+2 >= len(grid) {
		return 0
	}

	if grid[i][j] == 'M' && grid[i+1][j] == 'A' && grid[i+2][j] == 'S' {
		return 1
	}

	return 0
}

func getMasCount2(grid [][]rune, i int, j int) int {
	count := 0

	if i-1 < 0 || j-1 < 0 || i+1 >= len(grid) || j+1 >= len(grid) {
		return 0
	}

	if grid[i-1][j-1] == 'M' && grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S' && grid[i+1][j+1] == 'S' {
		count += 1
	}

	if grid[i-1][j-1] == 'S' && grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M' && grid[i+1][j+1] == 'M' {
		count += 1
	}

	if grid[i-1][j-1] == 'M' && grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M' && grid[i+1][j+1] == 'S' {
		count += 1
	}

	if grid[i-1][j-1] == 'S' && grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S' && grid[i+1][j+1] == 'M' {
		count += 1
	}

	return count
}

func getXmasCount1(xmas_grid [][]rune) int {

	xmas_count := 0
	for i, line := range xmas_grid {
		for j, char := range line {

			if char == 'X' {
				xmas_count += checkRight(xmas_grid, i, j+1)
				xmas_count += checkLeft(xmas_grid, i, j-1)
				xmas_count += checkTop(xmas_grid, i-1, j)
				xmas_count += checkBottom(xmas_grid, i+1, j)
				xmas_count += checkLeftUpperDiagonal(xmas_grid, i-1, j-1)
				xmas_count += checkRightUpperDiagonal(xmas_grid, i-1, j+1)
				xmas_count += checkLeftBottomDiagonal(xmas_grid, i+1, j-1)
				xmas_count += checkRightBottomDiagonal(xmas_grid, i+1, j+1)
			}
		}
	}

	return xmas_count
}

func getXmasCount2(xmas_grid [][]rune) int {

	xmas_count := 0

	for i, line := range xmas_grid {
		for j, char := range line {

			if char == 'A' {
				xmas_count += getMasCount2(xmas_grid, i, j)
			}
		}
	}
	return xmas_count
}

func main() {
	file, err := os.ReadFile("inputs/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(file), "\n")

	var xmas_grid [][]rune
	for _, line := range lines {
		row := []rune(line)
		xmas_grid = append(xmas_grid, row)
	}

	fmt.Printf("Answer for Day 4 Part 1: %d\n", getXmasCount1(xmas_grid))
	fmt.Printf("Answer for Day 4 Part 2: %d\n", getXmasCount2(xmas_grid))
}
