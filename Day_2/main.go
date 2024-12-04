package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isSafeReport(arr []int) bool {
	isIncreasing := true
	isDecreasing := true

	for j := 1; j < len(arr); j++ {
		del := arr[j] - arr[j-1]
		absDiff := int(math.Abs(float64(del)))

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if del > 0 {
			isDecreasing = false
		} else if del < 0 {
			isIncreasing = false
		}
	}
	return isDecreasing || isIncreasing
}

func getSafeReportCount(reports [][]int) int {
	safeReportCount := 0

	for _, arr := range reports {
		if isSafeReport(arr) {
			safeReportCount++
		}
	}

	return safeReportCount
}

func getSafetyReportCountWithDampener(reports [][]int) int {
	safeReportCountWithDampener := 0

	for _, arr := range reports {
		if isSafeReport(arr) {
			safeReportCountWithDampener++
			continue
		}

		isSafeWithDampener := false
		for i := 0; i < len(arr); i++ {
			copyArr := make([]int, len(arr))
			copy(copyArr, arr)

			var slicedReport []int
			if i == len(copyArr)-1 {
				slicedReport = copyArr[:len(copyArr)-1]
			} else {
				slicedReport = append(copyArr[:i], copyArr[i+1:]...)
			}

			if isSafeReport(slicedReport) {
				isSafeWithDampener = true
				break
			}
		}

		if isSafeWithDampener {
			safeReportCountWithDampener++
		}
	}
	return safeReportCountWithDampener
}

func main() {
	file, err := os.ReadFile("inputs/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	var reports [][]int

	reader := bytes.NewReader(file)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var row []int

		for _, field := range fields {

			number, err := strconv.Atoi(field)

			if err != nil {
				fmt.Println(err)
			}

			row = append(row, number)
		}

		reports = append(reports, row)
	}

	fmt.Printf("Answer for Day 2 Part 1: %d\n", getSafeReportCount(reports))
	fmt.Printf("Answer for Day 2 Part 2: %d\n", getSafetyReportCountWithDampener(reports))
}
