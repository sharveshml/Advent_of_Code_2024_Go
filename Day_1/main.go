package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fileContent, err := os.ReadFile("inputs/day1.txt")

	if err != nil {
		fmt.Println(err)
	}

	str := strings.Fields(string(fileContent))

	left := []int{}
	right := []int{}

	for i, s := range str {
		num, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println(err)
		}

		if i%2 == 0 {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		cur_sum := Abs(left[i] - right[i])

		sum += cur_sum

	}

	fmt.Printf("Day1 Answer 1: %d\n", sum)
	day1_2()
}

func day1_2() {

	fileContent, err := os.ReadFile("inputs/day1.txt")

	if err != nil {
		fmt.Println(err)
	}

	str := strings.Fields(string(fileContent))

	left := []int{}
	right := []int{}

	for i, s := range str {
		num, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println(err)
		}

		if i%2 == 0 {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	rightCount := make(map[int]int)

	for _, v := range right {
		rightCount[v]++
	}

	similarityScore := 0
	for _, num := range left {
		if count, exists := rightCount[num]; exists {
			similarityScore += num * count
		}
	}

	fmt.Printf("Day1 Answer 2: %d\n", similarityScore)
}
