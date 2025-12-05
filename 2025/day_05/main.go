package day_05

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type interval struct {
	From int
	To   int
}

func (i1 interval) merge(i2 interval) (ok bool, result interval) {
	if i2.From <= i1.To &&
		i2.From >= i1.From &&
		i2.To >= i1.To {
		ok = true
		result = interval{
			From: i1.From,
			To:   i2.To,
		}
		return
	}

	if i2.To <= i1.To &&
		i2.To >= i1.From &&
		i2.From <= i1.From {
		ok = true
		result = interval{
			From: i2.From,
			To:   i1.To,
		}
		return
	}

	if i2.From <= i1.From &&
		i2.To >= i1.To {
		result = i2
		ok = true
		return
	}

	if i1.From <= i2.From &&
		i1.To >= i2.To {
		result = i1
		ok = true
		return
	}

	return
}

func Main() {
	f, _ := os.ReadFile("day_05/input")
	s := string(f)

	// 	s = `3-5
	// 10-14
	// 16-20
	// 12-18

	// 1
	// 5
	// 8
	// 11
	// 17
	// 32
	// `

	s = s[:len(s)-1]

	input := strings.Split(s, "\n\n")
	intervalsInput := input[0]
	intervals := make([]interval, 0)
	for intervalString := range strings.SplitSeq(intervalsInput, "\n") {
		numbers := strings.Split(intervalString, "-")
		from, _ := strconv.Atoi(numbers[0])
		to, _ := strconv.Atoi(numbers[1])
		intervals = append(intervals, interval{From: from, To: to})
	}

	ingredientsInput := input[1]
	ingredients := make([]int, 0)
	for ingredientString := range strings.SplitSeq(ingredientsInput, "\n") {
		ingredient, _ := strconv.Atoi(ingredientString)
		ingredients = append(ingredients, ingredient)
	}

	intervalsCpy := make([]interval, len(intervals))
	copy(intervalsCpy, intervals)

	ingredientsCpy := make([]int, len(ingredients))
	copy(ingredientsCpy, ingredients)

	timeStart := time.Now()
	partOne(intervalsCpy, ingredientsCpy)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))

	timeStart = time.Now()
	partTwo(intervals)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
}

func partOne(freshIngredients []interval, ingredients []int) {
	slices.SortFunc(freshIngredients, func(a interval, b interval) int {
		return a.From - b.From
	})

	finalFreshIngredients := make([]interval, 1)
	finalFreshIngredients[0] = freshIngredients[0]

	for _, freshIngredient := range freshIngredients[1:] {
		if ok, mergedInterval := finalFreshIngredients[len(finalFreshIngredients)-1].merge(freshIngredient); ok {
			finalFreshIngredients[len(finalFreshIngredients)-1] = mergedInterval
		} else {
			finalFreshIngredients = append(finalFreshIngredients, freshIngredient)
		}
	}

	total := 0
	for _, ingredient := range ingredients {
		if slices.ContainsFunc(finalFreshIngredients, func(i interval) bool {
			return ingredient >= i.From && ingredient <= i.To
		}) {
			total++
		}
	}

	fmt.Println(total)
}

func partTwo(freshIngredients []interval) {
	slices.SortFunc(freshIngredients, func(a interval, b interval) int {
		return a.From - b.From
	})

	finalFreshIngredients := make([]interval, 1)
	finalFreshIngredients[0] = freshIngredients[0]

	for _, freshIngredient := range freshIngredients[1:] {
		if ok, mergedInterval := finalFreshIngredients[len(finalFreshIngredients)-1].merge(freshIngredient); ok {
			finalFreshIngredients[len(finalFreshIngredients)-1] = mergedInterval
		} else {
			finalFreshIngredients = append(finalFreshIngredients, freshIngredient)
		}
	}

	total := 0
	for _, interval := range finalFreshIngredients {
		total += interval.To - interval.From + 1
	}

	fmt.Println(total)
}
