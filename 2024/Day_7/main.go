package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData(path string) map[int][]int {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	equationMap := make(map[int][]int)

	for scanner.Scan() {
		line := scanner.Text()
		chunk := strings.Split(line, ":")

		key, err := strconv.Atoi(chunk[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		values := strings.Fields(chunk[1])
		intValues := make([]int, len(values))

		for i, val := range values {
			intValues[i], err = strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
			}
		}
		equationMap[key] = intValues
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return equationMap

}

func caluateEquations1(target int, values []int, ind int, currentResult int) bool {
	if ind >= len(values) {
		return currentResult == target
	}

	currentNumber := values[ind]
	newResult := currentResult + currentNumber

	if newResult <= target {
		if caluateEquations1(target, values, ind+1, newResult) {
			return true
		}
	}

	newResult = currentResult * currentNumber

	if newResult <= target {
		if caluateEquations1(target, values, ind+1, newResult) {
			return true
		}
	}

	return false
}

func caluateEquations2(target int, values []int, ind int, currentValue int) bool {
	if ind >= len(values) {
		return currentValue == target
	}

	currentNumber := values[ind]
	newResult := currentValue + currentNumber

	if newResult <= target {
		if caluateEquations2(target, values, ind+1, newResult) {
			return true
		}
	}

	newResult = currentValue * values[ind]

	if newResult <= target {
		if caluateEquations2(target, values, ind+1, newResult) {
			return true
		}
	}

	concatenatedValue, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentValue, currentNumber))

	if concatenatedValue <= target {
		if caluateEquations2(target, values, ind+1, concatenatedValue) {
			return true
		}
	}

	return false
}

func main() {
	equationMap := readData("inputs/input.txt")
	calibrationResult1 := 0
	calibrationResult2 := 0

	for k, v := range equationMap {
		if caluateEquations1(k, v, 0, 0) {
			calibrationResult1 += k
		}
	}

	for k, v := range equationMap {
		if caluateEquations2(k, v, 0, 0) {
			calibrationResult2 += k
		}
	}

	fmt.Printf("Answer for Day 7 Part 1: %d\n", calibrationResult1)
	fmt.Printf("Answer for Day 7 Part 2: %d\n", calibrationResult2)
}
