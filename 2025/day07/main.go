package day07

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func Main() {
	f, err := os.ReadFile("day_07/input")
	if err != nil {
		panic(err)
	}

	s := string(f)
	s = s[:len(s)-1]

	lines := strings.Split(s, "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
		}
	}

	timeStart := time.Now()
	res := firstPart(grid)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))
	fmt.Println("First part result: ", res)

	timeStart = time.Now()
	res = secondPart(grid)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
	fmt.Println("Second part result: ", res)
}

func firstPart(input [][]rune) (total int) {
	tachions := make([]int, 0)
	for i, char := range input[0] {
		if char == 'S' {
			tachions = append(tachions, i)
		}
	}
	input = input[1:]

	for _, line := range input {
		newTachions := make([]int, 0)
		for _, tachion := range tachions {
			if line[tachion] == '^' {
				newTachions = append(newTachions, tachion-1, tachion+1)
				total++
			} else {
				newTachions = append(newTachions, tachion)
			}
		}

		tachions = make([]int, 0, len(newTachions))
		for _, tachion := range newTachions {
			if !slices.Contains(tachions, tachion) {
				tachions = append(tachions, tachion)
			}
		}
	}

	return
}

type coordinate struct {
	X int
	Y int
}

var cache = make(map[coordinate]int)

/* dynamic programming my beloved */
func parseFromPosition(input [][]rune, coord coordinate) int {
	if val, ok := cache[coord]; ok {
		return val
	}

	if coord.Y >= len(input) {
		return 1
	}

	switch input[coord.Y][coord.X] {
	case '.', 'S':
		result := parseFromPosition(
			input,
			coordinate{X: coord.X, Y: coord.Y + 1},
		)

		cache[coord] = result
		return result
	case '^':
		result := parseFromPosition(
			input,
			coordinate{X: coord.X - 1, Y: coord.Y + 1},
		) + parseFromPosition(
			input,
			coordinate{X: coord.X + 1, Y: coord.Y + 1},
		)
		cache[coord] = result
		return result
	default:
		panic("invalid character")
	}
}

func secondPart(input [][]rune) (total int) {
	tachions := make([]int, 0)
	for i, char := range input[0] {
		if char == 'S' {
			tachions = append(tachions, i)
		}
	}
	input = input[1:]

	for _, tachion := range tachions {
		total += parseFromPosition(input, coordinate{X: tachion, Y: 0})
	}

	return
}
