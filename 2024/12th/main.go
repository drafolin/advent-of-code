package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
)

type area struct {
	plots     []utils.Coordinate
	perimeter int
}

func main() {
	dataStr := []string{
		"AAAA\nBBCD\nBBCC\nEEEC",
		"OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO",
		"RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE",
	}
	/*dataStr := []string{
		utils.ReadInput("12th"),
	}*/

	for i, dataStr := range dataStr {
		data := utils.StrToGrid(dataStr)
		fmt.Println(i, scanField(data))
	}
}

func scanField(data utils.Grid) int {
	res := make([]*area, 0)

	for y, row := range data {
		for x := range row {
			if slices.ContainsFunc(res, func(a *area) bool {
				return slices.Contains(a.plots, utils.Coordinate{X: x, Y: y})
			}) {
				continue
			}
			scanNewArea(data, utils.Coordinate{X: x, Y: y}, &res)
		}
	}

	sum := 0
	for _, group := range res {
		sum += group.perimeter * len(group.plots)
	}
	return sum
}

func scanNewArea(data utils.Grid, coordinate utils.Coordinate, res *[]*area) {

	i := len(*res)
	*res = append(*res, new(area))

	scanArea(data, coordinate, res, i)
}

func scanArea(data utils.Grid, coord utils.Coordinate, res *[]*area, i int) {
	plot := data.At(coord)
	currentArea := (*res)[i]
	if slices.Contains(currentArea.plots, coord) {
		return
	}

	(*res)[i].plots = append((*res)[i].plots, coord)

	for _, dir := range []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right} {
		newCoord := coord.MoveTowards(dir)
		if !newCoord.IsInGrid(data) {
			currentArea.perimeter++
			continue
		}

		if data.At(newCoord) != plot {
			currentArea.perimeter++
		} else {
			scanArea(data, newCoord, res, i)
		}
	}
}
