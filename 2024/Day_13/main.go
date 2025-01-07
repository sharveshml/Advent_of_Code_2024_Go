package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	ButtonAX int
	ButtonAY int
	ButtonBX int
	ButtonBY int
	PrizeX   int
	PrizeY   int
}

func readData(path string) []Input {
	var inputs []Input

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var temp Input
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		} else if strings.HasPrefix(line, "Button A") {
			parts := strings.Split(line, ",")
			ax, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], "+")[1]))
			ay, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[1], "+")[1]))
			temp.ButtonAX = ax
			temp.ButtonAY = ay
		} else if strings.HasPrefix(line, "Button B") {
			parts := strings.Split(line, ",")
			bx, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], "+")[1]))
			by, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[1], "+")[1]))
			temp.ButtonBX = bx
			temp.ButtonBY = by
		} else if strings.HasPrefix(line, "Prize") {
			parts := strings.Split(line, ",")
			px, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], "=")[1]))
			py, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[1], "=")[1]))
			temp.PrizeX = px
			temp.PrizeY = py

			inputs = append(inputs, temp)

			temp = Input{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return inputs
}

func solveForMachine(input Input) (bool, int) {
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			if a*input.ButtonAX+b*input.ButtonBX == input.PrizeX &&
				a*input.ButtonAY+b*input.ButtonBY == input.PrizeY {
				return true, a*3 + b*1
			}
		}
	}
	return false, 0
}

func solveForMachine1(input Input) (bool, int) {
	// Using Cramer's Rule to solve for a and b

	det := input.ButtonAX*input.ButtonBY - input.ButtonBX*input.ButtonAY
	if det == 0 {
		return false, 0
	}

	targetX := input.PrizeX + 10000000000000
	targetY := input.PrizeY + 10000000000000

	a := float64(targetX*input.ButtonBY-input.ButtonBX*targetY) / float64(det)
	b := float64(input.ButtonAX*targetY-targetX*input.ButtonAY) / float64(det)

	if a < 0 || b < 0 {
		return false, 0
	}

	if math.Abs(a-math.Round(a)) > 1e-10 || math.Abs(b-math.Round(b)) > 1e-10 {
		return false, 0
	}

	aInt := int(math.Round(a))
	bInt := int(math.Round(b))

	if aInt*input.ButtonAX+bInt*input.ButtonBX != targetX ||
		aInt*input.ButtonAY+bInt*input.ButtonBY != targetY {
		return false, 0
	}

	tokens := aInt*3 + bInt*1
	return true, tokens
}

func main() {
	inputs := readData("inputs/input.txt")
	totalTokens := 0
	totalTokens1 := 0

	for _, input := range inputs {
		if solvable, tokens := solveForMachine(input); solvable {
			totalTokens += tokens
		}
	}

	for _, input := range inputs {
		if solvable, tokens := solveForMachine1(input); solvable {
			totalTokens1 += tokens
		}
	}

	fmt.Printf("Answer for Day 13 Part 1: %d\n", totalTokens)
	fmt.Printf("Answer for Day 13 Part 2: %d\n", totalTokens1)
}
