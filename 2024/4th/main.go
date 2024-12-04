package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/coords"
	"github.com/drafolin/advent-of-code/2024/coords/directions"
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
				coord := coords.Coordinate{X: x, Y: y}
				diagCnt := 0
				for _, diag := range checkDiagsFor(coord, 'M', dataGrid) {
					sCoord := coord.MoveTowards(diag.Opposite())
					if v, _ := sCoord.IsInGrid(dataGrid); v && dataGrid[sCoord.Y][sCoord.X] == 'S' {
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

func checkDiagsFor(coord coords.Coordinate, letter rune, data [][]rune) []directions.Direction {
	res := make([]directions.Direction, 0)

	for _, direction := range []directions.Direction{directions.UpLeft, directions.UpRight, directions.DownLeft, directions.DownRight} {
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
