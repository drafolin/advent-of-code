package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
	"strconv"
)

var smallestSortedId = 0

func main() {
	//data := "2333133121414131402"
	data := utils.ReadInput("9th")
	dataArr := make([]int, 0)

	for i, char := range data {
		isEmptySpace := i%2 == 1
		cnt, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

		if isEmptySpace {
			dataArr = append(dataArr, slices.Repeat([]int{-1}, cnt)...)
		} else {
			dataArr = append(dataArr, slices.Repeat([]int{i / 2}, cnt)...)
		}

		smallestSortedId = i / 2
	}

	fmt.Println(calcChecksum(compactDisk(dataArr)))
	defragmented := defragment(dataArr)
	fmt.Println(calcChecksum(defragmented))
}

func compactDisk(disk []int) []int {
	diskCpy := make([]int, len(disk))
	copy(diskCpy, disk)

	i := len(diskCpy)

	smallestEmptyIndex := 0

	for !isDiskCompacted(diskCpy) {
		i--

		if diskCpy[i] == -1 {
			continue
		}

		for diskCpy[smallestEmptyIndex] != -1 {
			smallestEmptyIndex++
		}

		diskCpy[smallestEmptyIndex] = diskCpy[i]
		diskCpy[i] = -1
	}

	return diskCpy
}

func countBlockLength(disk []int, atIndex int) int {
	cnt := 0

	target := disk[atIndex]

	for atIndex < len(disk) && disk[atIndex] == target {
		cnt++
		atIndex++
	}

	return cnt
}

func indexAtStartOfBlock(disk []int, atIndex int) int {
	target := disk[atIndex]

	for atIndex >= 0 && disk[atIndex] == target {
		atIndex--
	}

	return atIndex + 1
}

func defragment(disk []int) []int {
	defragmentedDisk := make([]int, len(disk))
	copy(defragmentedDisk, disk)

	for smallestSortedId >= 0 {
		startOfBlock := slices.Index(defragmentedDisk, smallestSortedId)
		blockLength := countBlockLength(defragmentedDisk, startOfBlock)

		i := 0

		for i < startOfBlock {
			length := countBlockLength(defragmentedDisk, i)

			if defragmentedDisk[i] != -1 {
				i += length
				continue
			}

			if length >= blockLength {
				break
			}

			i += length
		}

		if i == startOfBlock {
			smallestSortedId--
			continue
		}

		startOfNewBlock := indexAtStartOfBlock(defragmentedDisk, i)

		for i := 0; i < blockLength; i++ {
			defragmentedDisk[startOfNewBlock+i] = smallestSortedId
			defragmentedDisk[startOfBlock+i] = -1
		}

		smallestSortedId--
	}

	return defragmentedDisk
}

func isDiskCompacted(disk []int) bool {
	isAtEnd := false

	for _, block := range disk {
		if block == -1 {
			isAtEnd = true
		}

		if block != -1 && isAtEnd {
			return false
		}
	}

	return true
}

func calcChecksum(disk []int) int {
	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}
		sum += disk[i] * i
	}

	return sum
}
