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
	/*
		dataStr := []string{
			"AAAA\nBBCD\nBBCC\nEEEC",
			"OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO",
			"RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE",
		}
	*/

	dataStr := []string{
		utils.ReadInput("12th"),
	}

	for i, dataStr := range dataStr {
		data := utils.StrToGrid(dataStr)
		sumPerimeters, sumSides := scanField(data)
		fmt.Println(i, sumPerimeters, sumSides)
	}
}

func scanField(data utils.Grid) (int, int) {
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

	sumPerimeters := 0
	sumSides := 0
	for _, group := range res {
		sumPerimeters += group.perimeter * len(group.plots)
		sumSides += len(group.plots) * countCorners(*group, data)
	}
	return sumPerimeters, sumSides
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

func isSamePlant(coord1, coord2 utils.Coordinate, data utils.Grid) bool {
	if coord1.IsInGrid(data) != coord2.IsInGrid(data) {
		return false
	}

	if !coord1.IsInGrid(data) {
		return true
	}

	return data.At(coord1) == data.At(coord2)
}

func countCorners(group area, data utils.Grid) int {
	corners := 0
	for _, plot := range group.plots {
		for _, dir := range []utils.Direction{utils.UpLeft, utils.UpRight, utils.DownLeft, utils.DownRight} {
			if !slices.Contains(group.plots, plot.MoveTowards(dir.Rotate(45))) && !slices.Contains(group.plots, plot.MoveTowards(dir.Rotate(-45))) {
				corners++
				continue
			}

			if slices.Contains(group.plots, plot.MoveTowards(dir.Rotate(45))) &&
				slices.Contains(group.plots, plot.MoveTowards(dir.Rotate(-45))) &&
				!slices.Contains(group.plots, plot.MoveTowards(dir)) {
				corners++
				continue
			}
		}
	}

	return corners
}
