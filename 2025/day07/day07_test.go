package day07

import (
	"strings"
	"testing"
)

const s = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func TestPart1(t *testing.T) {
	lines := strings.Split(s, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
		}
	}
	res := firstPart(grid)

	if res != 21 {
		t.Errorf("First part is %d, expected %d", res, 21)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(s, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
		}
	}
	res := secondPart(grid)

	if res != 40 {
		t.Errorf("Second part is %d, expected %d", res, 40)
	}
}
