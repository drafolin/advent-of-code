package main

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"github.com/drafolin/advent-of-code/2024/utils"
	"math"
	"strings"
)

//var grid = utils.NewTypedGridFunc(7, 7, func(x, y int) *MemoryAdress { return &MemoryAdress{coord: utils.Coordinate{X: x, Y: y}} })

var grid = utils.NewTypedGridFunc(71, 71, func(x, y int) *MemoryAdress {
	return &MemoryAdress{
		coord:     utils.Coordinate{X: x, Y: y},
		corrupted: false,
	}
})

type MemoryAdress struct {
	coord     utils.Coordinate
	corrupted bool
}

func (ma *MemoryAdress) PathNeighbors() (neighbors []astar.Pather) {
	neighbors = make([]astar.Pather, 0)

	for _, dir := range []utils.Direction{utils.Up, utils.Left, utils.Down, utils.Right} {
		if grid.HasCoord(ma.coord.MoveTowards(dir)) && !grid.At(ma.coord.MoveTowards(dir)).corrupted {
			neighbors = append(neighbors, grid.At(ma.coord.MoveTowards(dir)))
		}
	}

	return
}

func (ma *MemoryAdress) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (ma *MemoryAdress) PathEstimatedCost(to astar.Pather) float64 {
	return math.Abs(float64(ma.coord.X-to.(*MemoryAdress).coord.X)) + math.Abs(float64(ma.coord.Y-to.(*MemoryAdress).coord.Y))
}

func main() {
	//input := utils.StrToLineList("5,4\n4,2\n4,5\n3,0\n2,1\n6,3\n2,4\n1,5\n0,6\n3,3\n2,6\n5,1\n1,2\n5,5\n2,5\n6,5\n1,4\n0,4\n6,4\n1,1\n6,1\n1,0\n0,5\n1,6\n2,0")
	input := utils.ReadInputSplitted("18th")

	//const bytesToRead = 12
	const bytesToRead = 1024

	for _, coordStr := range input[:bytesToRead] {
		splittedCoord := strings.Split(coordStr, ",")
		coord := utils.Coordinate{X: utils.StringToInt(splittedCoord[0]), Y: utils.StringToInt(splittedCoord[1])}
		grid[coord.Y][coord.X].corrupted = true
	}

	input = input[bytesToRead-1:]

	firstRun := true
	var lastCoord string

	for {
		_, distance, found := astar.Path(grid[0][0], grid[grid.Height()-1][grid.Width()-1])
		if !found {
			break
		}

		if firstRun {
			fmt.Println(distance)
			firstRun = false
		}

		lastCoord = input[0]
		input = input[1:]
		splittedCoord := strings.Split(lastCoord, ",")
		grid[utils.StringToInt(splittedCoord[1])][utils.StringToInt(splittedCoord[0])].corrupted = true
	}

	fmt.Println(lastCoord)
}
