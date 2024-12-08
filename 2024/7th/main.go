package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"strconv"
	"strings"
)

type operator int

const (
	mul operator = iota
	add
)

type solution struct {
	val int
	ops []operator
}

func main() {
	// data := "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20"
	data := utils.ReadInput("7th")

	lines := utils.StrToLineList(data)

	cnt := 0

	for _, line := range lines {
		nbrsStr := strings.Split(line, ": ")[1]
		nbrsStrs := strings.Split(nbrsStr, " ")
		target, err := strconv.Atoi(strings.Split(line, ": ")[0])

		if err != nil {
			panic(err)
		}

		nbrs := make([]int, len(nbrsStrs))
		for i, nbrStr := range nbrsStrs {
			nbrs[i], err = strconv.Atoi(nbrStr)
			if err != nil {
				panic(err)
			}
		}

		if compute(nbrs, mul, target) || compute(nbrs, add, target) {
			cnt += target
		}
	}

	fmt.Println(cnt)
}

func compute(nbrs []int, with operator, target int) bool {
	if len(nbrs) == 1 {
		return nbrs[0] == target
	}

	firstNr := nbrs[0]

	if with == mul {
		firstNr *= nbrs[1]
	} else {
		firstNr += nbrs[1]
	}

	newNbrs := make([]int, len(nbrs)-1)
	copy(newNbrs, nbrs[1:])
	newNbrs[0] = firstNr

	if compute(newNbrs, mul, target) || compute(newNbrs, add, target) {
		return true
	}

	return false
}
