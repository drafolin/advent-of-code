package day_04

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
 * Yeah, this solution is not very good, but I am not sure what would be the optimal way to do it.
 * Also, it works. And since it's go, it only takes a few milliseconds.
 */

type Grid [][]bool

func Main() {
	f, err := os.ReadFile("day_04/input")
	if err != nil {
		panic(err)
	}

	s := string(f)

	/*
			s = `..@@.@@@@.
		@@@.@.@.@@
		@@@@@.@.@@
		@.@@@@..@.
		@@.@@@@.@@
		.@@@@@@@.@
		.@.@.@.@@@
		@.@@@.@@@@
		.@@@@@@@@.
		@.@.@@@.@.
		`
	*/
	s = s[:len(s)-1]

	lines := strings.Split(s, "\n")
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for j, cell := range line {
			grid[i][j] = cell == '@'
		}
	}

	timeStart := time.Now()
	firstPart(grid)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))

	timeStart = time.Now()
	secondPart(grid)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
}

func firstPart(grid Grid) {
	totalAccessible := 0
	for y, row := range grid {
		for x := range row {
			if !grid[y][x] {
				continue
			}

			count := 0

			for _, dy := range []int{-1, 0, 1} {
				for _, dx := range []int{-1, 0, 1} {
					if dy == 0 && dx == 0 {
						continue
					}

					if y+dy < 0 || y+dy >= len(grid) || x+dx < 0 || x+dx >= len(row) {
						continue
					}

					if grid[y+dy][x+dx] {
						count++
					}
				}
			}

			if count < 4 {
				totalAccessible++
			}
		}
	}

	fmt.Println(totalAccessible)
}

func secondPart(grid Grid) {
	totalRemoved := 0
	for {
		toRemove := []struct {
			y int
			x int
		}{}

		for y, row := range grid {
			for x := range row {
				if !grid[y][x] {
					continue
				}

				count := 0

				for _, dy := range []int{-1, 0, 1} {
					for _, dx := range []int{-1, 0, 1} {
						if dy == 0 && dx == 0 {
							continue
						}

						if y+dy < 0 || y+dy >= len(grid) || x+dx < 0 || x+dx >= len(row) {
							continue
						}

						if grid[y+dy][x+dx] {
							count++
						}
					}
				}

				if count < 4 {
					toRemove = append(toRemove, struct {
						y int
						x int
					}{y, x})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, cell := range toRemove {
			grid[cell.y][cell.x] = false
			totalRemoved++
		}
	}

	fmt.Println(totalRemoved)
}
