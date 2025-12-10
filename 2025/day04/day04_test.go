package day04

import (
	"strings"
	"testing"
)

const s = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestPart1(t *testing.T) {
	lines := strings.Split(s, "\n")

	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for j, cell := range line {
			grid[i][j] = cell == '@'
		}
	}

	res := firstPart(grid)

	if res != 13 {
		t.Errorf("First part is %d, expected %d", res, 13)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(s, "\n")
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for j, cell := range line {
			grid[i][j] = cell == '@'
		}
	}
	res := secondPart(grid)

	if res != 43 {
		t.Errorf("Second part is %d, expected %d", res, 43)
	}
}
