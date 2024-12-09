package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
)

func main() {
	//data := "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............"
	data := utils.ReadInput("8th")
	dataGrid := utils.StrToGrid(data)

	antennaes := scanForAntennaes(dataGrid)

	fmt.Println(len(getAntinodesWithDoubleSpace(dataGrid, antennaes)))
	fmt.Println(len(getAntinodesWithAnySpace(dataGrid, antennaes)))
}

func scanForAntennaes(dataGrid utils.Grid) map[rune][]utils.Coordinate {
	antennaes := make(map[rune][]utils.Coordinate)

	for y, line := range dataGrid {
		for x, char := range line {
			if char == '.' {
				continue
			}

			antennaes[char] = append(antennaes[char], utils.Coordinate{X: x, Y: y})
		}
	}

	return antennaes
}

func getAntinodesWithDoubleSpace(dataGrid utils.Grid, antennaes map[rune][]utils.Coordinate) []utils.Coordinate {
	antinodes := make([]utils.Coordinate, 0)

	for _, anntenaes := range antennaes {
		for _, antenna1 := range anntenaes {
			for _, antenna2 := range anntenaes {
				if antenna1 == antenna2 {
					continue
				}

				diff := antenna1.Diff(antenna2)
				antinode := antenna1.Add(diff)

				if v, _ := antinode.IsInAnyGrid(dataGrid); v {
					if !slices.Contains(antinodes, antinode) {
						antinodes = append(antinodes, antinode)
					}
				}
			}
		}
	}

	return antinodes
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func getAntinodesWithAnySpace(dataGrid utils.Grid, antennaes map[rune][]utils.Coordinate) []utils.Coordinate {
	antinodes := make([]utils.Coordinate, 0)

	for _, anntenaes := range antennaes {
		for _, antenna1 := range anntenaes {
			for _, antenna2 := range anntenaes {
				if antenna1 == antenna2 {
					continue
				}

				diff := antenna1.Diff(antenna2)
				gcd := gcd(diff.X, diff.Y)
				diff.X /= gcd
				diff.Y /= gcd

				for antinode := antenna1; antinode.IsInGrid(dataGrid); antinode = antinode.Add(diff) {
					if v, _ := antinode.IsInAnyGrid(dataGrid); v {
						if !slices.Contains(antinodes, antinode) {
							antinodes = append(antinodes, antinode)
						}
					}
				}

				for antinode := antenna1.Diff(diff); antinode.IsInGrid(dataGrid); antinode = antinode.Diff(diff) {
					if v, _ := antinode.IsInAnyGrid(dataGrid); v {
						if !slices.Contains(antinodes, antinode) {
							antinodes = append(antinodes, antinode)
						}
					}
				}
			}
		}
	}

	return antinodes
}
