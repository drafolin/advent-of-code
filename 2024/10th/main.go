package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
	"strconv"
)

func main() {
	// data := "0123\n1234\n8765\n9876"
	data := "89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732"
	// data := utils.ReadInput("10th")
	dataGrid := utils.StrToGrid(data)
	scoreSum := 0

	for y, line := range dataGrid {
		for x, char := range line {
			if char != '0' {
				continue
			}

			res := recurseToEnd(dataGrid, 1, utils.Coordinate{X: x, Y: y})

			final := make([]utils.Coordinate, 0, len(res))
			for _, coord := range res {
				if !slices.Contains(final, coord) {
					final = append(final, coord)
				}
			}

			score := len(final)

			scoreSum += score
		}
	}

	fmt.Println(scoreSum)
}

func recurseToEnd(grid utils.Grid, nextNr int, currentCoord utils.Coordinate) []utils.Coordinate {
	res := make([]utils.Coordinate, 0)
	for _, dir := range []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right} {
		newCoord := currentCoord.MoveTowards(dir)

		if !newCoord.IsInGrid(grid) {
			continue
		}

		if v, _ := strconv.Atoi(string(grid.At(newCoord))); v == nextNr {
			if nextNr == 9 {
				res = append(res, newCoord)
			} else {
				res = append(res, recurseToEnd(grid, nextNr+1, newCoord)...)
			}
		}
	}

	return res
}
