package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
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
		bots[i].Pos = bots[i].Pos.Add(bots[i].Vel.MulInt(100)).Mod(utils.Coordinate{X: grid.Width(), Y: grid.Height()})

		quad, isInQuad := grid.GetQuadrant(bots[i].Pos)
		if isInQuad {
			quads[quad]++
		}
	}

	safetyFactor := 1

	for _, quad := range quads {
		safetyFactor *= quad
	}

	fmt.Println(safetyFactor)

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
