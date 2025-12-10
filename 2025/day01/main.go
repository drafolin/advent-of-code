package day01

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Main() {
	f, err := os.ReadFile("day_01/input")
	if err != nil {
		panic(err)
	}

	s := string(f)
	s = s[:len(s)-1]
	lines := strings.Split(s, "\n")

	timeStart := time.Now()
	res := firstPart(lines)
	timeEnd := time.Now()
	fmt.Println("Part 1 took ", timeEnd.Sub(timeStart))
	fmt.Println("Part 1 solution: ", res)

	timeStart = time.Now()
	res = secondPart(lines)
	timeEnd = time.Now()
	fmt.Println("Part 2 took ", timeEnd.Sub(timeStart))
	fmt.Println("Part 2 solution: ", res)
}

func firstPart(lines []string) (turns int) {
	val := 50
	for _, line := range lines {
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
	return
}

func secondPart(lines []string) (turns int) {
	val := 50

	for _, line := range lines {
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

	return
}
