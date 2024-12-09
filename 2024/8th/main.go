package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
)

func main() {
	// data := "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............"
	data := utils.ReadInput("8th")
	dataGrid := utils.StrToGrid(data)

	antinodes := make([]utils.Coordinate, 0)
	antennaes := make(map[rune][]utils.Coordinate)

	for y, line := range dataGrid {
		for x, char := range line {
			if char == '.' {
				continue
			}

			antennaes[char] = append(antennaes[char], utils.Coordinate{X: x, Y: y})
		}
	}

	for _, anntenaes := range antennaes {
		for _, antenna1 := range anntenaes {
			for _, antenna2 := range anntenaes {
				if antenna1 == antenna2 {
					continue
				}

				diff := antenna1.Diff(antenna2)
				antinode := antenna1.Add(diff)

				if v, _ := antinode.IsInGrid(dataGrid); v {
					if !slices.Contains(antinodes, antinode) {
						antinodes = append(antinodes, antinode)
					}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}
