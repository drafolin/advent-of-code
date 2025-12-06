package day_06

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Main() {
	f, err := os.ReadFile("day_06/input")
	if err != nil {
		panic(err)
	}

	s := string(f)
	// 	s = `123 328  51 64
	//  45 64  387 23
	//   6 98  215 314
	// *   +   *   +
	// `

	s = s[:len(s)-1]

	timeStart := time.Now()
	firstPart(s)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))

	timeStart = time.Now()
	secondPart(s)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
}

func firstPart(input string) {
	lines := strings.Split(input, "\n")

	data := make([][]string, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")

		x := 0

		for _, part := range parts {
			if part == "" {
				continue
			}

			if len(data) <= x {
				data = append(data, make([]string, 0))
			}

			data[x] = append(data[x], part)

			x++
		}
	}

	grandTotal := 0
	for _, column := range data {
		operator := column[len(column)-1]
		numbers := column[:len(column)-1]

		total, _ := strconv.Atoi(numbers[0])

		for _, number := range numbers[1:] {
			n, _ := strconv.Atoi(number)
			switch operator {
			case "+":
				total += n
			case "*":
				total *= n
			}
		}

		grandTotal += total
	}

	fmt.Println(grandTotal)
}

func secondPart(input string) {
	lines := strings.Split(input, "\n")

	// data := make([][]string, 0)

	rangeStart := 0
	rangeIndex := 0

	operatorsLine := lines[len(lines)-1]
	operatorsLine += " " + string(rune(0))
	grandTotal := 0
	for i, char := range operatorsLine {
		if i == 0 {
			continue
		}

		if char == ' ' {
			continue
		}

		number := make([]string, i-rangeStart-1)
		for _, line := range lines[:len(lines)-1] {
			for j, char := range line[rangeStart : i-1] {
				if char == ' ' {
					continue
				}

				number[j] += string(char)
			}
		}

		total, _ := strconv.Atoi(number[0])
		operator := operatorsLine[rangeStart]
		for _, number := range number[1:] {
			n, _ := strconv.Atoi(number)
			switch operator {
			case '+':
				total += n
			case '*':
				total *= n
			}
		}

		grandTotal += total

		rangeStart = i
		rangeIndex++
	}

	fmt.Println(grandTotal)
}
