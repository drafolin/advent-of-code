package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"strings"
)

type Towel []rune

func main() {
	//availableTowels, patterns := parseInput("r, wr, b, g, bwu, rb, gb, br\n\nbrwrr\nbggr\ngbbr\nrrbgbr\nubwu\nbwurrg\nbrgr\nbbrgwb")
	availableTowels, patterns := parseInput(utils.ReadInput("19th"))

	sumPart1 := 0
	sumPart2 := 0
	for _, pattern := range patterns {
		if part1(pattern, availableTowels) {
			sumPart1++
		}
		sumPart2 += part2(pattern, availableTowels)
	}
	fmt.Printf("Part 1: %d, Part 2: %d", sumPart1, sumPart2)
}

func parseInput(in string) (availableTowels []Towel, patterns []string) {
	splittedInput := strings.Split(in, "\n\n")
	towelsString := splittedInput[0]
	for _, towelString := range strings.Split(towelsString, ", ") {
		availableTowels = append(availableTowels, []rune(towelString))
	}

	patterns = strings.Split(splittedInput[1], "\n")
	return
}

func part1(pattern string, availableTowels []Towel) bool {
	if len(pattern) == 0 {
		return true
	}

	for _, availableTowel := range availableTowels {
		if len(availableTowel) > len(pattern) {
			continue
		}

		if string(availableTowel) == pattern[:len(availableTowel)] {
			if part1(pattern[len(availableTowel):], availableTowels) {
				return true
			}
		}
	}

	return false
}

var cache = make(map[string]int)

func part2(pattern string, availableTowels []Towel) (sum int) {
	if res, ok := cache[pattern]; ok {
		return res
	}

	if len(pattern) == 0 {
		cache[pattern] = 1
		return 1
	}

	for _, availableTowel := range availableTowels {
		if len(availableTowel) > len(pattern) {
			continue
		}

		if string(availableTowel) == pattern[:len(availableTowel)] {
			sum += part2(pattern[len(availableTowel):], availableTowels)
		}
	}

	cache[pattern] = sum
	return
}
