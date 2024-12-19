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

	fmt.Println(availableTowels, patterns)
	sum := 0
	for _, pattern := range patterns {
		if solvePattern(pattern, availableTowels) {
			sum++
		}
	}
	fmt.Println(sum)
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

func solvePattern(pattern string, availableTowels []Towel) bool {
	if len(pattern) == 0 {
		return true
	}

	for _, availableTowel := range availableTowels {
		if len(availableTowel) > len(pattern) {
			continue
		}

		if string(availableTowel) == pattern[:len(availableTowel)] {
			if solvePattern(pattern[len(availableTowel):], availableTowels) {
				return true
			}
		}
	}

	return false
}
