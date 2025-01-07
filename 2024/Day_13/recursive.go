// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type Input struct {
// 	ButtonAX int
// 	ButtonAY int
// 	ButtonBX int
// 	ButtonBY int
// 	PrizeX   int
// 	PrizeY   int
// }

// func readData(path string) []Input {
// 	var inputs []Input

// 	file, err := os.Open(path)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	var temp Input
// 	for scanner.Scan() {
// 		line := strings.TrimSpace(scanner.Text())

// 		if line == "" {
// 			continue
// 		} else if strings.HasPrefix(line, "Button A") {
// 			parts := strings.Split(line, ",")
// 			ax, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], "+")[1]))
// 			ay, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[1], "+")[1]))
// 			temp.ButtonAX = ax
// 			temp.ButtonAY = ay
// 		} else if strings.HasPrefix(line, "Button B") {
// 			parts := strings.Split(line, ",")
// 			bx, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], "+")[1]))
// 			by, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[1], "+")[1]))
// 			temp.ButtonBX = bx
// 			temp.ButtonBY = by
// 		} else if strings.HasPrefix(line, "Prize") {
// 			parts := strings.Split(line, ",")
// 			px, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], "=")[1]))
// 			py, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[1], "=")[1]))
// 			temp.PrizeX = px
// 			temp.PrizeY = py

// 			inputs = append(inputs, temp)

// 			temp = Input{}
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading file:", err)
// 		os.Exit(1)
// 	}

// 	return inputs
// }

// type State struct {
// 	x, y   int
// 	aCount int
// 	bCount int
// }

// func findMinTokens(curr State, input Input, memo map[State]int) int {
// 	if curr.aCount > 100 || curr.bCount > 100 {
// 		return -1
// 	}

// 	if curr.x > input.PrizeX || curr.y > input.PrizeY {
// 		return -1
// 	}

// 	if curr.x == input.PrizeX && curr.y == input.PrizeY {
// 		return curr.aCount*3 + curr.bCount
// 	}

// 	if val, exists := memo[curr]; exists {
// 		return val
// 	}

// 	minTokens := -1
// 	newStateA := State{
// 		x:      curr.x + input.ButtonAX,
// 		y:      curr.y + input.ButtonAY,
// 		aCount: curr.aCount + 1,
// 		bCount: curr.bCount,
// 	}
// 	tokensA := findMinTokens(newStateA, input, memo)

// 	newStateB := State{
// 		x:      curr.x + input.ButtonBX,
// 		y:      curr.y + input.ButtonBY,
// 		aCount: curr.aCount,
// 		bCount: curr.bCount + 1,
// 	}
// 	tokensB := findMinTokens(newStateB, input, memo)

// 	if tokensA != -1 && (minTokens == -1 || tokensA < minTokens) {
// 		minTokens = tokensA
// 	}
// 	if tokensB != -1 && (minTokens == -1 || tokensB < minTokens) {
// 		minTokens = tokensB
// 	}

// 	memo[curr] = minTokens
// 	return minTokens
// }

// func main() {
// 	inputs := readData("inputs/input.txt")
// 	totalTokens := 0

// 	for _, input := range inputs {
// 		memo := make(map[State]int)
// 		initialState := State{0, 0, 0, 0}
// 		tokens := findMinTokens(initialState, input, memo)
// 		if tokens != -1 {
// 			totalTokens += tokens
// 		}
// 	}

// 	fmt.Println(totalTokens)
// }
