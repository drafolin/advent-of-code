package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
	"strconv"
)

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
	}

	fmt.Println(calcChecksum(compactDisk(dataArr)))

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
	for i := 0; disk[i] != -1; i++ {
		sum += disk[i] * i
	}

	return sum
}
