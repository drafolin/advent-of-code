package day06

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
	s = s[:len(s)-1]

	lines := strings.Split(s, "\n")

	timeStart := time.Now()
	res := firstPart(lines)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))
	fmt.Println("First part result: ", res)

	timeStart = time.Now()
	res = secondPart(lines)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
	fmt.Println("Second part result: ", res)
}

func firstPart(lines []string) (grandTotal int) {
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

	return
}

func secondPart(lines []string) (grandTotal int) {
	rangeStart := 0
	rangeIndex := 0

	operatorsLine := lines[len(lines)-1]
	operatorsLine += " " + string(rune(0))
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

	return
}
