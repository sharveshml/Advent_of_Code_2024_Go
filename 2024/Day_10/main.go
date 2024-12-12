package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readData(path string) [][]int {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	grid := [][]int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tmp := []int{}

		for _, i := range line {
			intVal, err := strconv.Atoi(string(i))
			if err != nil {
				fmt.Println(err)
			}
			tmp = append(tmp, intVal)
		}

		grid = append(grid, tmp)
	}
	return grid
}

func getScore1(grid [][]int, i int, j int, slope int, visited map[[2]int]bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}

	if grid[i][j] != slope {
		return
	}

	currentPos := [2]int{i, j}

	if grid[i][j] == 9 {
		visited[currentPos] = true
		return
	}

	getScore1(grid, i+1, j, slope+1, visited)
	getScore1(grid, i-1, j, slope+1, visited)
	getScore1(grid, i, j+1, slope+1, visited)
	getScore1(grid, i, j-1, slope+1, visited)
}

func getScore2(grid [][]int, i int, j int, slope int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return 0
	}

	if grid[i][j] != slope {
		return 0
	}

	if grid[i][j] == 9 {
		return 1
	}

	score := 0

	score += getScore2(grid, i+1, j, slope+1)
	score += getScore2(grid, i-1, j, slope+1)
	score += getScore2(grid, i, j+1, slope+1)
	score += getScore2(grid, i, j-1, slope+1)

	return score

}
func main() {
	grid := readData("inputs/input.txt")
	output := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				visited := make(map[[2]int]bool)
				getScore1(grid, i, j, 0, visited)

				output += len(visited)
			}
		}
	}

	trial := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				trial += getScore2(grid, i, j, 0)
			}
		}
	}

	fmt.Printf("Answer for Day 10 Part 1: %d\n", output)
	fmt.Printf("Answer for Day 10 Part 2: %d\n", trial)
}
