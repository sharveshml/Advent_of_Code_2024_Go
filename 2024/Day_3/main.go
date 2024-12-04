package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getMuls(nums string) int {
	pattern := regexp.MustCompile(`(\d{1,3}),(\d{1,3})`)
	matches := pattern.FindAllStringSubmatch(nums, -1)

	if len(matches) == 0 {
		fmt.Println("No matches found")
		return 0
	}

	num1Str := matches[0][1]
	num2Str := matches[0][2]

	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)

	if err1 != nil || err2 != nil {
		fmt.Println("Error converting string to integer")
		return 0
	}

	return num1 * num2
}

func unCorruptDoDonts(str string) []string {
	pattern := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don\'t\(\))`)
	match := pattern.FindAllString(str, -1)

	return match
}

func unCorrupt(str string) []string {
	pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	match := pattern.FindAllString(str, -1)

	return match
}

func main() {
	file, err := os.ReadFile("inputs/inputs.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	unCorruptedString := unCorrupt(string(file))
	unCorruptedStringWithDoDonts := unCorruptDoDonts(string(file))

	sum := 0
	for _, str := range unCorruptedString {
		sum += getMuls(str)
	}

	sumDoDonts := 0
	active := true

	for _, str1 := range unCorruptedStringWithDoDonts {
		if strings.Contains(str1, "don't()") {
			active = false
			continue
		} else if strings.Contains(str1, "do()") {
			active = true
			continue
		} else {
			if active {
				sumDoDonts += getMuls(str1)
			}
		}
	}

	fmt.Printf("Answer for Day 3 Part 1: %d\n", sum)
	fmt.Printf("Answer for Day 3 Part 2: %d\n", sumDoDonts)
}
