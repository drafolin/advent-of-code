package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
)

func main() {
	// data := "125 17"
	data := utils.ReadInput("11th")
	dataList, err := utils.StrToNumberList(data)
	if err != nil {
		panic(err)
	}

	countAfter6Times := stoneCountAfterBlinks(dataList, 6)
	fmt.Printf("After 6 times: %d\n", countAfter6Times)
	countAfter25Times := stoneCountAfterBlinks(dataList, 25)
	fmt.Printf("After 25 times: %d\n", countAfter25Times)
	countAfter75Times := stoneCountAfterBlinks(dataList, 75)
	fmt.Printf("After 75 times: %d\n", countAfter75Times)
}

func stoneCountAfterBlinks(stones []int, blinks int) int {
	count := 0
	for _, stone := range stones {
		count += stoneTransformAfterBlinks(stone, blinks-1)
	}

	return count
}

type StoneTransformInput struct {
	stone      int
	blinkCount int
}

var stoneTransformCache = make(map[StoneTransformInput]int)

func stoneTransformAfterBlinks(stone int, blinkCount int) int {
	if stone, ok := stoneTransformCache[StoneTransformInput{stone, blinkCount}]; ok {
		return stone
	}

	stones := make([]int, 0)

	if stone == 0 {
		stones = append(stones, 1)
	} else if length := intLen(stone); length%2 == 0 {
		digits := length / 2
		multiplicator := intPow(10, digits)
		stones = append(stones, stone/multiplicator, stone%multiplicator)
	} else {
		stones = append(stones, stone*2024)
	}

	count := 0
	if blinkCount == 0 {
		stoneTransformCache[StoneTransformInput{stone, blinkCount}] = len(stones)
		return len(stones)
	} else {
		for _, stone := range stones {
			count += stoneTransformAfterBlinks(stone, blinkCount-1)
		}
	}

	stoneTransformCache[StoneTransformInput{stone, blinkCount}] = count
	return count
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func intLen(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
