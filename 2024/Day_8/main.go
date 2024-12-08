package main

import (
	"bufio"
	"fmt"
	"os"
)

func readData(path string) [][]rune {
	var grid [][]rune
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		var row []rune
		for _, i := range line {
			row = append(row, i)
		}
		grid = append(grid, row)
	}

	return grid
}

func getAntinodesCount(grid [][]rune) int {
	antenaMap := make(map[rune][][2]int)
	antiNodes := make(map[[2]int]bool)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != '.' {
				antenaMap[grid[row][col]] = append(antenaMap[grid[row][col]], [2]int{row, col})
			}
		}
	}

	for _, positions := range antenaMap {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {

				pos1 := positions[i]
				pos2 := positions[j]

				delRow := pos2[0] - pos1[0]
				delCol := pos2[1] - pos1[1]

				antiNode1 := [2]int{pos1[0] - delRow, pos1[1] - delCol}
				antiNode2 := [2]int{pos2[0] + delRow, pos2[1] + delCol}

				if antiNode1[0] >= 0 && antiNode1[0] < len(grid) && antiNode1[1] >= 0 && antiNode1[1] < len(grid[0]) {
					antiNodes[antiNode1] = true
				}
				if antiNode2[0] >= 0 && antiNode2[0] < len(grid) && antiNode2[1] >= 0 && antiNode2[1] < len(grid[0]) {
					antiNodes[antiNode2] = true
				}
			}
		}
	}
	return len(antiNodes)

}

func getAntiNodeCount2(grid [][]rune) int {
	antenaMap := make(map[rune][][2]int)
	antiNodes := make(map[[2]int]bool)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != '.' {
				antenaMap[grid[row][col]] = append(antenaMap[grid[row][col]], [2]int{row, col})
			}
		}
	}

	for _, positions := range antenaMap {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {

				pos1 := positions[i]
				pos2 := positions[j]

				delRow := pos2[0] - pos1[0]
				delCol := pos2[1] - pos1[1]

				antiNode1Row := pos1[0]
				antiNode1Col := pos1[1]

				for antiNode1Row >= 0 && antiNode1Row < len(grid) && antiNode1Col >= 0 && antiNode1Col < len(grid) {
					antiNodes[[2]int{antiNode1Row, antiNode1Col}] = true

					antiNode1Row -= delRow
					antiNode1Col -= delCol
				}

				antiNode2Row := pos2[0]
				antiNode2Col := pos2[1]

				for antiNode2Row >= 0 && antiNode2Row < len(grid) && antiNode2Col >= 0 && antiNode2Col < len(grid) {
					antiNodes[[2]int{antiNode2Row, antiNode2Col}] = true

					antiNode2Row += delRow
					antiNode2Col += delCol
				}
			}
		}
	}
	return len(antiNodes)
}

func main() {
	grid := readData("inputs/input.txt")

	fmt.Printf("Answer for Day 8 Part 1: %d\n", getAntinodesCount(grid))
	fmt.Printf("Answer for Day 8 Part 2: %d\n", getAntiNodeCount2(grid))
}
