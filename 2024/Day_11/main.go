package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData(path string) []string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content := string(file)
	stones := strings.Fields(content)
	return stones
}

func getStoneCount(stones []string) int {
	for i := 0; i < 25; i++ {
		newStones := []string{}
		for _, stone := range stones {
			if stone == "0" {
				newStones = append(newStones, "1")
			} else if len(stone)%2 == 0 {
				half := len(stone) / 2
				left, _ := strconv.Atoi(stone[:half])
				right, _ := strconv.Atoi(stone[half:])
				newStones = append(newStones, strconv.Itoa(left), strconv.Itoa(right))
			} else {
				val, _ := strconv.Atoi(stone)
				val *= 2024
				newStones = append(newStones, strconv.Itoa(val))
			}
		}
		stones = newStones
	}
	return len(stones)
}

func getStoneCount2(input []string) int {

	stones := map[int]int{}

	for _, stone := range input {
		num, _ := strconv.Atoi(stone)
		stones[num]++
	}

	for i := 0; i < 75; i++ {
		newStones := map[int]int{}
		for stone, count := range stones {
			if stone == 0 {
				newStones[1] += count
			} else if len(strconv.Itoa(stone))%2 == 0 {
				numStr := strconv.Itoa(stone)
				half := len(strconv.Itoa(stone)) / 2
				left, _ := strconv.Atoi(numStr[:half])
				right, _ := strconv.Atoi(numStr[half:])

				newStones[left] += count
				newStones[right] += count
			} else {
				newStones[stone*2024] += count
			}
		}
		stones = newStones
	}
	output := 0
	for _, count := range stones {
		output += count
	}
	return output
}

func main() {
	input := readData("inputs/input.txt")
	fmt.Println("Final count of stones:", getStoneCount(input))
	fmt.Println("Final count of stones:", getStoneCount2(input))
}
