package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"math"
	"slices"
	"strings"
)

type Bot struct {
	Pos, Vel utils.Coordinate
}

func main() {
	//data := utils.StrToLineList("p=0,4 v=3,-3\np=6,3 v=-1,-3\np=10,3 v=-1,2\np=2,0 v=2,-1\np=0,0 v=1,3\np=3,0 v=-2,-2\np=7,6 v=-1,-3\np=3,0 v=-1,-2\np=9,3 v=2,3\np=7,3 v=-1,2\np=2,4 v=2,-3\np=9,5 v=-3,-3")
	data := utils.ReadInputSplitted("14th")

	bots := make([]Bot, len(data))

	quads := map[utils.Direction]int{
		utils.UpLeft:    0,
		utils.UpRight:   0,
		utils.DownLeft:  0,
		utils.DownRight: 0,
	}

	grid := make(utils.Grid, 103)
	grid[0] = make([]rune, 101)

	for i, line := range data {
		bots[i] = parseBot(line)
	}

	for i := range bots {
		botsCpy := make([]Bot, len(bots))
		copy(botsCpy, bots)

		botsCpy[i].Pos = botsCpy[i].Pos.Add(botsCpy[i].Vel.MulInt(100)).Mod(utils.Coordinate{X: grid.Width(), Y: grid.Height()})
		quad, isInQuad := grid.GetQuadrant(botsCpy[i].Pos)
		if isInQuad {
			quads[quad]++
		}
	}

	safetyFactor := 1

	for _, quad := range quads {
		safetyFactor *= quad
	}

	fmt.Println(safetyFactor)

	// get the smallest Y variance
	var smallestYVariance float64
	var smallestYVarianceIndex int

	for i := 0; i < grid.Height(); i++ {
		var thisYVariance float64
		botsCpy := make([]Bot, len(bots))
		copy(botsCpy, bots)

		sumOfY := 0
		for j := range botsCpy {
			botsCpy[j].Pos = botsCpy[j].Pos.Add(botsCpy[j].Vel.MulInt(i)).Mod(utils.Coordinate{X: grid.Width(), Y: grid.Height()})
			sumOfY += botsCpy[j].Pos.Y
		}

		averageY := float64(sumOfY) / float64(len(botsCpy))

		for _, bot := range botsCpy {
			thisYVariance += math.Abs(float64(bot.Pos.Y) - averageY)
		}

		thisYVariance = thisYVariance / float64(len(botsCpy))

		if i == 0 || thisYVariance < smallestYVariance {
			smallestYVariance = thisYVariance
			smallestYVarianceIndex = i
		}
	}

	// get the smallest X variance
	var smallestXVariance float64
	var smallestXVarianceIndex int

	for i := 0; i < grid.Width(); i++ {
		var thisXVariance float64
		botsCpy := make([]Bot, len(bots))
		copy(botsCpy, bots)

		sumOfX := 0
		for j := range botsCpy {
			botsCpy[j].Pos = botsCpy[j].Pos.Add(botsCpy[j].Vel.MulInt(i)).Mod(utils.Coordinate{X: grid.Width(), Y: grid.Height()})
			sumOfX += botsCpy[j].Pos.X
		}

		averageX := float64(sumOfX) / float64(len(botsCpy))

		for _, bot := range botsCpy {
			thisXVariance += math.Abs(float64(bot.Pos.X) - averageX)
		}

		thisXVariance = thisXVariance / float64(len(botsCpy))

		if i == 0 || thisXVariance < smallestXVariance {
			smallestXVariance = thisXVariance
			smallestXVarianceIndex = i
		}
	}

	var w = func(x int) int {
		return x*grid.Width() + smallestXVarianceIndex
	}

	var h = func(y int) int {
		return y*grid.Height() + smallestYVarianceIndex
	}

	wResults := make([]int, 0)
	hResults := make([]int, 0)

	step2Result := 0

	for i := 0; ; i++ {
		wRes := w(i)
		hRes := h(i)

		if slices.Contains(hResults, wRes) {
			step2Result = wRes
			break
		}

		if slices.Contains(wResults, hRes) {
			step2Result = hRes
			break
		}

		wResults = append(wResults, wRes)
		hResults = append(hResults, hRes)
	}

	fmt.Println(step2Result)
}

func parseBot(line string) Bot {
	datas := strings.Split(line, " ")
	pString := strings.Split(datas[0], "=")[1]
	vString := strings.Split(datas[1], "=")[1]

	pX := strings.Split(pString, ",")[0]
	pY := strings.Split(pString, ",")[1]
	p := utils.Coordinate{
		X: utils.StringToInt(pX),
		Y: utils.StringToInt(pY),
	}

	vX := strings.Split(vString, ",")[0]
	vY := strings.Split(vString, ",")[1]
	v := utils.Coordinate{
		X: utils.StringToInt(vX),
		Y: utils.StringToInt(vY),
	}

	return Bot{
		Pos: p,
		Vel: v,
	}
}
