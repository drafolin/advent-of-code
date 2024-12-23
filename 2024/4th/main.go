package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
)

func main() {
	/*
			data := `..X...
		.SAMX.
		.A..A.
		XMAS.S
		.X....`*/

	/*data := `MMMSXXMASM
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
			if char == 'A' {
				coord := utils.Coordinate{X: x, Y: y}
				diagCnt := 0
				for _, diag := range checkDiagsFor(coord, 'M', dataGrid) {
					sCoord := coord.MoveTowards(diag.Opposite())
					if v, _ := sCoord.IsInAnyGrid(dataGrid); v && dataGrid[sCoord.Y][sCoord.X] == 'S' {
						diagCnt++
					}
				}
				if diagCnt == 2 {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}

func checkDiagsFor(coord utils.Coordinate, letter rune, data [][]rune) []utils.Direction {
	res := make([]utils.Direction, 0)

	for _, direction := range []utils.Direction{utils.UpLeft, utils.UpRight, utils.DownLeft, utils.DownRight} {
		newCoord := coord.MoveTowards(direction)
		if v, err := newCoord.IsInAnyGrid(data); !v {
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
