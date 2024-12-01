package main

import (
	"github.com/drafolin/advent-of-code/2024/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//const data = "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"
	data := utils.ReadInput("1st")

	pairs := strings.Split(data, "\n")

	list1 := make([]int, len(pairs))
	list2 := make([]int, len(pairs))

	for i, pair := range pairs {
		splitPair := strings.Split(pair, "   ")
		n1, _ := strconv.Atoi(splitPair[0])
		n2, _ := strconv.Atoi(splitPair[1])
		list1[i] = n1
		list2[i] = n2
	}

	computeSimilarityScore(list1, list2)
}

func computeTotalDistance(list1 []int, list2 []int) {
	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0.0
	for i := 0; i < len(list1); i++ {
		sum += math.Abs(float64(list1[i] - list2[i]))
	}

	println(int(sum))
}

func computeSimilarityScore(list1 []int, list2 []int) {
	score := 0

	appearances := make(map[int]int)

	for _, n := range list2 {
		appearances[n]++
	}

	for _, n := range list1 {
		score += n * appearances[n]
	}

	println(score)
}
