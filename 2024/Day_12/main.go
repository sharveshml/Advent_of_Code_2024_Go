package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readData(path string) [][]rune {
	grid := [][]rune{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tmp := []rune{}
		for _, i := range line {
			tmp = append(tmp, i)
		}
		grid = append(grid, tmp)
	}

	return grid
}

func getCost(grid [][]rune, visited [][]bool, i int, j int, plant rune, area int, perimeter int) (int, int) {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
		if grid[i][j] == plant && !visited[i][j] {
			visited[i][j] = true
			area++

			area, perimeter = getCost(grid, visited, i-1, j, plant, area, perimeter)
			area, perimeter = getCost(grid, visited, i+1, j, plant, area, perimeter)
			area, perimeter = getCost(grid, visited, i, j-1, plant, area, perimeter)
			area, perimeter = getCost(grid, visited, i, j+1, plant, area, perimeter)
		} else if grid[i][j] != plant {
			perimeter++
		}
	} else {
		perimeter++
	}
	return area, perimeter
}

func inGrid(grid [][]rune, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func rotate(direction [2]int) [2]int {
	return [2]int{direction[1], -direction[0]}
}

func cVal(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}
	return c
}

func corners(grid [][]rune, row, col int, c rune) int {
	upLeft := [2]int{-1, -1}
	upRight := [2]int{-1, 0}
	left := [2]int{0, -1}

	cornerCount := 0

	for d := 0; d < 4; d++ {
		upperLeft := [2]int{row + upLeft[0], col + upLeft[1]}
		upperRight := [2]int{row + upRight[0], col + upRight[1]}
		lowerLeft := [2]int{row + left[0], col + left[1]}

		isUpperLeft := inGrid(grid, upperLeft[0], upperLeft[1]) && cVal(grid[upperLeft[0]][upperLeft[1]]) == c
		isUpperRight := inGrid(grid, upperRight[0], upperRight[1]) && cVal(grid[upperRight[0]][upperRight[1]]) == c
		isLowerLeft := inGrid(grid, lowerLeft[0], lowerLeft[1]) && cVal(grid[lowerLeft[0]][lowerLeft[1]]) == c

		cornerCount += boolToInt(!isUpperRight && !isLowerLeft)
		cornerCount += boolToInt(isUpperRight && isLowerLeft && !isUpperLeft)

		upLeft = rotate(upLeft)
		upRight = rotate(upRight)
		left = rotate(left)
	}

	return cornerCount
}

func fill(grid [][]rune, row, col int, c rune, plot *map[string]int) {
	if !inGrid(grid, row, col) || grid[row][col] != c {
		return
	}

	(*plot)["edges"] += corners(grid, row, col, c)
	(*plot)["area"]++
	grid[row][col] = rune(strings.ToLower(string(grid[row][col]))[0])

	fill(grid, row-1, col, c, plot)
	fill(grid, row+1, col, c, plot)
	fill(grid, row, col-1, c, plot)
	fill(grid, row, col+1, c, plot)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func solve(grid [][]rune) int {
	output := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] >= 'A' && grid[row][col] <= 'Z' {
				plot := map[string]int{"area": 0, "edges": 0}
				fill(grid, row, col, grid[row][col], &plot)
				output += plot["area"] * plot["edges"]
			}
		}
	}

	return output
}

func main() {
	input := readData("inputs/input.txt")

	visited := make([][]bool, len(input))
	for i := range visited {
		visited[i] = make([]bool, len(input[i]))
	}

	out := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if !visited[i][j] {
				area, perimeter := getCost(input, visited, i, j, input[i][j], 0, 0)
				out += area * perimeter
			}
		}
	}
	fmt.Printf("Answer for Day 12 Part 1: %d\n", out)

	result := solve(input)

	fmt.Printf("Answer for Day 12 Part 2: %d\n", result)
}
