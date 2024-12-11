package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"math"
	"strconv"
)

func main() {
	//data := "125 17"
	data := utils.ReadInput("11th")
	dataList, err := utils.StrToNumberList(data)
	if err != nil {
		panic(err)
	}

	listAfter25Times := applyRulesNTimes(dataList, 25)
	fmt.Printf("After 25 times: %d", len(listAfter25Times))
}

func applyRulesNTimes(list []int, blinks int) []int {
	listCpy := make([]int, len(list))
	copy(listCpy, list)
	for i := 0; i < blinks; i++ {
		newList := make([]int, 0, len(list))
		for _, stone := range listCpy {
			newList = append(newList, applyRules(stone)...)
		}
		listCpy = newList
	}

	return listCpy
}

func applyRules(n int) []int {
	switch true {
	case n == 0:
		return []int{1}
	case len(strconv.Itoa(n))%2 == 0:
		digits := len(strconv.Itoa(n)) / 2
		left := n / int(math.Pow10(digits))
		right := n % int(math.Pow10(digits))
		return []int{left, right}
	default:
		return []int{n * 2024}
	}
}
