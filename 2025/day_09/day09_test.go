package day_09

import (
	"strconv"
	"strings"
	"testing"
)

const testInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

func TestPartOne(t *testing.T) {
	input := testInput[:len(testInput)-1]

	lines := strings.Split(input, "\n")

	coordinates := make([]coordinate, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		coordinates = append(coordinates, coordinate{X: x, Y: y})
	}

	res := firstPart(coordinates)

	if res != 50 {
		t.Errorf("First part result is %d, expected 50", res)
	}
}

func TestPartTwo(t *testing.T) {
	input := testInput[:len(testInput)-1]

	lines := strings.Split(input, "\n")

	coordinates := make([]coordinate, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		coordinates = append(coordinates, coordinate{X: x, Y: y})
	}

	res := secondPart(coordinates)

	if res != 24 {
		t.Errorf("Second part result is %d, expected 24", res)
	}
}
