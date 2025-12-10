package day05

import (
	"strconv"
	"strings"
	"testing"
)

const s = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestPart1(t *testing.T) {
	parts := strings.Split(s, "\n\n")
	intervals := make([]interval, 0)
	for intervalString := range strings.SplitSeq(parts[0], "\n") {
		numbers := strings.Split(intervalString, "-")
		from, _ := strconv.Atoi(numbers[0])
		to, _ := strconv.Atoi(numbers[1])
		intervals = append(intervals, interval{From: from, To: to})
	}

	ingredients := make([]int, 0)
	for ingredientString := range strings.SplitSeq(parts[1], "\n") {
		ingredient, _ := strconv.Atoi(ingredientString)
		ingredients = append(ingredients, ingredient)
	}

	res := firstPart(intervals, ingredients)

	if res != 3 {
		t.Errorf("First part is %d, expected %d", res, 3)
	}
}

func TestPart2(t *testing.T) {
	parts := strings.Split(s, "\n\n")
	intervals := make([]interval, 0)
	for intervalString := range strings.SplitSeq(parts[0], "\n") {
		numbers := strings.Split(intervalString, "-")
		from, _ := strconv.Atoi(numbers[0])
		to, _ := strconv.Atoi(numbers[1])
		intervals = append(intervals, interval{From: from, To: to})
	}

	res := secondPart(intervals)

	if res != 14 {
		t.Errorf("Second part is %d, expected %d", res, 14)
	}
}
