package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readData(file []byte) (map[int][]int, [][]int) {
	lines := strings.Split(string(file), "\n")
	var section1, section2 []string
	emptyLineIndex := -1

	for i, line := range lines {
		if line == "" {
			emptyLineIndex = i
			break
		}
	}

	if emptyLineIndex != -1 {
		section1 = lines[:emptyLineIndex]
		section2 = lines[emptyLineIndex+1:]
	} else {
		section1 = lines
	}

	ruleMap := make(map[int][]int)

	for _, rule := range section1 {
		ruleParts := strings.Split(rule, "|")

		num1, _ := strconv.Atoi(ruleParts[0])
		num2, _ := strconv.Atoi(ruleParts[1])

		ruleMap[num1] = append(ruleMap[num1], num2)
	}

	updateSet := [][]int{}

	for _, nums := range section2 {
		nums := strings.Split(nums, ",")

		set := []int{}

		for _, n := range nums {
			num, _ := strconv.Atoi(n)
			set = append(set, num)
		}
		updateSet = append(updateSet, set)
	}

	return ruleMap, updateSet
}

func isCorrectUpdate(ruleMap map[int][]int, updateSet []int) bool {

	for i := 1; i < len(updateSet); i++ {
		for j := 0; j < i; j++ {
			if slices.Contains(ruleMap[updateSet[i]], updateSet[j]) {
				return false
			}
		}
	}
	return true
}

func partOne(file []byte) int {
	ruleMap, updateSet := readData(file)
	var pageNumbers int

	for _, pages := range updateSet {
		if isCorrectUpdate(ruleMap, pages) {
			pageNumbers += pages[len(pages)/2]
		}
	}

	return pageNumbers
}

func sortPages(ruleMap map[int][]int, updateSet []int) []int {
	for {
		swapped := false
		for i := 1; i < len(updateSet); i++ {
			for j := 0; j < i; j++ {
				if slices.Contains(ruleMap[updateSet[i]], updateSet[j]) {
					temp := updateSet[j]
					updateSet[j] = updateSet[i]
					updateSet[i] = temp
					swapped = true
				}
			}
		}
		if !swapped {
			break
		}
	}
	return updateSet
}

func parttwo(file []byte) int {
	ruleMap, updateSet := readData(file)
	pageCount := 0

	for _, pages := range updateSet {
		if !isCorrectUpdate(ruleMap, pages) {
			sortedSet := sortPages(ruleMap, pages)
			pageCount += sortedSet[len(sortedSet)/2]
		}
	}
	return pageCount
}

func main() {
	file, err := os.ReadFile("inputs/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Answer for Day 5 Part 1: %d\n", partOne(file))
	fmt.Printf("Answer for Day 5 Part 2: %d\n", parttwo(file))

}
