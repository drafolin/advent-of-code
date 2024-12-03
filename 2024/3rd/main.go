package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"regexp"
	"strconv"
)

func main() {
	// data := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	// data := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	data := utils.ReadInput("3rd")

	mults := getConditionalMults(data)

	sum := 0

	for _, mult := range mults {
		sum += mult[0] * mult[1]
	}

	fmt.Println(sum)
}

func getMults(input string) [][2]int {
	r, err := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	if err != nil {
		panic(err)
	}

	matches := r.FindAllStringSubmatch(input, -1)

	matchesValues := make([][2]int, 0)

	for _, match := range matches {
		match1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		match2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		matchesValues = append(matchesValues, [2]int{match1, match2})
	}

	return matchesValues
}

func getConditionalMults(input string) [][2]int {
	r, err := regexp.Compile("(mul|don't|do)\\((?:(\\d{1,3}),(\\d{1,3}))?\\)")

	if err != nil {
		panic(err)
	}

	matches := r.FindAllStringSubmatch(input, -1)

	shouldDo := true

	corrects := make([][2]int, 0)

	for _, match := range matches {
		switch match[1] {
		case "mul":
			if !shouldDo {
				break
			}

			match1, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			match2, err := strconv.Atoi(match[3])
			if err != nil {
				panic(err)
			}

			corrects = append(corrects, [2]int{match1, match2})
		case "don't":
			shouldDo = false
		case "do":
			shouldDo = true
		}
	}

	return corrects
}
