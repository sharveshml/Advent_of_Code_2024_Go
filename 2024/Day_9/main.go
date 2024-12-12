package main

import (
	"fmt"
	"os"
	"strconv"
)

func processDiskMap(input []rune) ([]rune, int) {
	isBlock := true
	blockIndex := 0
	i := 0
	var diskMap []rune

	for i < len(input) {
		if isBlock {
			num, err := strconv.Atoi(string(input[i]))
			if err != nil {
				fmt.Println(err)
			}
			for j := 0; j < num; j++ {
				diskMap = append(diskMap, rune('0'+blockIndex))
			}
			i++
			blockIndex++
		} else {
			num, err := strconv.Atoi(string(input[i]))
			if err != nil {
				fmt.Println(err)
			}
			for j := 0; j < num; j++ {
				diskMap = append(diskMap, '.')
			}
			i++
		}
		isBlock = !isBlock
	}
	fmt.Println(string(diskMap))
	return diskMap, blockIndex
}

func moveFileBlocks1(diskMap []rune) []rune {
	emptyIndex := 0
	lastValidIndex := len(diskMap) - 1

	for {
		for emptyIndex < len(diskMap) && diskMap[emptyIndex] != '.' {
			emptyIndex++
		}
		for lastValidIndex >= 0 && diskMap[lastValidIndex] == '.' {
			lastValidIndex--
		}

		if emptyIndex < lastValidIndex {
			diskMap[emptyIndex] = diskMap[lastValidIndex]
			diskMap[lastValidIndex] = '.'
		}
		if emptyIndex >= lastValidIndex {
			break
		}
	}
	return diskMap
}

func moveFileBlocks2(diskMap []rune, fileId int) []rune {

	for currentFile := fileId - 1; currentFile >= 0; currentFile-- {
		fileBlocks := []int{}

		for i := 0; i < len(diskMap); i++ {
			if diskMap[i] == rune('0'+currentFile) {
				fileBlocks = append(fileBlocks, i)
			}
		}

		start := -1
		length := 0

		for i := 0; i < fileBlocks[0]; i++ {

			if diskMap[i] != '.' {
				start = -1
				length = 0
				continue
			}

			if start == -1 {
				start = i
			}

			length++

			if length == len(fileBlocks) {
				break
			}
		}
		// fmt.Printf("Swapping %v with %d and with length %d\n", fileBlocks, start, length)
		if length == len(fileBlocks) {
			for i := 0; i < length; i++ {
				diskMap[start+i] = rune('0' + currentFile)
				diskMap[fileBlocks[i]] = '.'
			}
		}
	}
	// fmt.Printf("After moving %s\n", string(diskMap))
	return diskMap
}

func processCheckSum(diskMap []rune) int {
	sum := 0
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] != '.' {
			sum += i * int(diskMap[i]-'0')
		}
	}
	return sum
}

func main() {
	file, err := os.ReadFile("inputs/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var input []rune
	content := string(file)

	for _, i := range content {
		input = append(input, i)
	}

	diskMap, blockIndex := processDiskMap(input)
	diskMap = moveFileBlocks2(diskMap, blockIndex)
	checkSum := processCheckSum(diskMap)

	fmt.Println(checkSum)
}
