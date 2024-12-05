package main

import (
	"errors"
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
	"strconv"
	"strings"
)

type pair [2]int

func main() {
	// data := "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47"

	data := utils.ReadInput("5th")

	rules, updates, err := parseInput(data)
	if err != nil {
		panic(err)
	}

	correctSum := 0
	incorrectSum := 0

	for _, update := range updates {
		valid := true
		applyingRules := make([]pair, 0)
		for _, rule := range rules {
			if !slices.Contains(update, rule[0]) || !slices.Contains(update, rule[1]) {
				continue
			}

			applyingRules = append(applyingRules, rule)

			index1 := slices.Index(update, rule[0])
			index2 := slices.Index(update, rule[1])

			if index1 >= index2 {
				valid = false
			}
		}

		if valid {
			middle, err := getMiddle(update)
			if err != nil {
				panic(err)
			}
			correctSum += middle
		} else {
			slices.SortStableFunc(update, func(i, j int) int {
				if (slices.Contains(applyingRules, pair{i, j})) {
					return -1
				} else if (slices.Contains(applyingRules, pair{j, i})) {
					return 1
				} else {
					return 0
				}
			})

			middle, err := getMiddle(update)
			if err != nil {
				panic(err)
			}
			incorrectSum += middle
		}
	}

	fmt.Println("Correct sum :", correctSum)
	fmt.Println("Incorrect sum :", incorrectSum)
}

func parseInput(in string) ([]pair, [][]int, error) {
	input := strings.Split(in, "\n\n")
	rules := strings.Split(input[0], "\n")
	updates := strings.Split(input[1], "\n")

	parsedRules := make([]pair, len(rules))

	for i, rule := range rules {
		ruleSplit := strings.Split(rule, "|")
		rule1, err := strconv.Atoi(ruleSplit[0])
		if err != nil {
			return nil, nil, err
		}
		rule2, err := strconv.Atoi(ruleSplit[1])
		if err != nil {
			return nil, nil, err
		}
		parsedRules[i] = pair{rule1, rule2}
	}

	parsedUpdates := make([][]int, len(updates))

	for i, update := range updates {
		updateSplit := strings.Split(update, ",")
		parsedUpdates[i] = make([]int, len(updateSplit))
		for j, num := range updateSplit {
			var err error
			parsedUpdates[i][j], err = strconv.Atoi(num)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return parsedRules, parsedUpdates, nil
}

func getMiddle(array []int) (int, error) {
	if len(array)%2 == 0 {
		return 0, errors.New("the array has an even number of elements")
	}

	return array[len(array)/2], nil
}
