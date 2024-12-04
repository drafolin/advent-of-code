package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/coords"
	"github.com/drafolin/advent-of-code/2024/utils"
)

func main() {
	/*
			data := `..X...
		.SAMX.
		.A..A.
		XMAS.S
		.X....`*/
	/*
	   		data := `MMMSXXMASM
	   MSAMXMSMSA
	   AMXSXMAAMM
	   MSAMASMSMX
	   XMASAMXAMM
	   XXAMMXXAMA
	   SMSMSASXSS
	   SAXAMASAAA
	   MAMMMXMMMM
	   MXMXAXMASX`*/

	data := utils.ReadInput("4th")

	dataGrid := utils.StrToGrid(data)
	cnt := 0
	for y, line := range dataGrid {
		for x, char := range line {
			if char == 'X' {
				coord := coords.Coordinate{X: x, Y: y}
				cnt += checkForMasAround(coord, dataGrid)
			}
		}
	}
	fmt.Println(cnt)
}

func checkForMasAround(coord coords.Coordinate, data [][]rune) int {
	cnt := 0
	if directions := checkAroundLetterFor(coord, 'M', data); len(directions) > 0 {
		for _, direction := range directions {
			aCoord := coord.MoveTowards(direction).MoveTowards(direction)
			sCoord := aCoord.MoveTowards(direction)

			if v, _ := sCoord.IsInGrid(data); !v {
				continue
			}

			if data[aCoord.Y][aCoord.X] == 'A' &&
				data[sCoord.Y][sCoord.X] == 'S' {
				cnt++
			}
		}
	}

	return cnt
}

func checkAroundLetterFor(coord coords.Coordinate, letter rune, data [][]rune) []coords.Direction {
	res := make([]coords.Direction, 0)
	for direction := coords.Direction(0); direction <= coords.Direction(7); direction++ {
		newCoord := coord.MoveTowards(direction)
		if v, err := newCoord.IsInGrid(data); !v {
			if err != nil {
				panic(err)
			}
			continue
		}

		if data[newCoord.Y][newCoord.X] == letter {
			res = append(res, direction)
		}
	}

	return res
}
