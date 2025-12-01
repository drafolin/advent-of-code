package day_01

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Main() {
	f, err := os.ReadFile("day_01/input")
	if err != nil {
		panic(err)
	}

	s := string(f)

	s = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

	firstPart(s)
	secondPart(s)
}

func firstPart(input string) {
	lines := strings.SplitSeq(input[:len(input)-1], "\n")
	val := 50
	turns := 0
	for line := range lines {
		operation := line[0]
		amount, _ := strconv.Atoi(line[1:])

		switch operation {
		case 'R':
			val += amount % 100
		case 'L':
			val -= amount % 100
		}

		if val < 0 || val >= 100 {
			val = (val%100 + 100) % 100
		}

		if val == 0 {
			turns++
		}
	}

	fmt.Println(turns)
}

func secondPart(input string) {

	lines := strings.SplitSeq(input[:len(input)-1], "\n")
	val := 50
	turns := 0
	for line := range lines {
		operation := line[0]
		amount, _ := strconv.Atoi(line[1:])

		prevVal := val

		switch operation {
		case 'R':
			val += amount
			if val >= 100 {
				if prevVal < 100 {
					turns += int(math.Abs(math.Floor(float64(val) / 100)))
				}
			}
		case 'L':
			val -= amount
			if val <= 0 {
				if prevVal > 0 {
					turns += int(math.Abs(math.Ceil(float64(val)/100)) + 1)
				} else if prevVal == 0 {
					turns += int(math.Abs(math.Ceil(float64(val) / 100)))
				}
			}
		}

		if val < 0 || val >= 100 {
			val = (val%100 + 100) % 100
		}
	}

	fmt.Println(turns)
}
