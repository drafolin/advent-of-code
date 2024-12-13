package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"math"
	"strconv"
	"strings"
)

type Machine struct {
	ButtonA, ButtonB, Prize utils.Coordinate
}

func main() {
	// dataStr := "Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400\n\nButton A: X+26, Y+66\nButton B: X+67, Y+21\nPrize: X=12748, Y=12176\n\nButton A: X+17, Y+86\nButton B: X+84, Y+37\nPrize: X=7870, Y=6450\n\nButton A: X+69, Y+23\nButton B: X+27, Y+71\nPrize: X=18641, Y=10279"
	dataStr := utils.ReadInput("13th")

	machinesStrs := strings.Split(dataStr, "\n\n")

	res := 0
	resPart2 := 0

	for i, machineStr := range machinesStrs {
		machine := Machine{}
		machinePart2 := Machine{}
		lines := strings.Split(machineStr, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "Button A") {
				machine.ButtonA = parseCoordinate(line)
				machinePart2.ButtonA = parseCoordinate(line)
			} else if strings.HasPrefix(line, "Button B") {
				machine.ButtonB = parseCoordinate(line)
				machinePart2.ButtonB = parseCoordinate(line)
			} else if strings.HasPrefix(line, "Prize") {
				machine.Prize = parseCoordinate(line)
				machinePart2.Prize = parseCoordinate(line).Add(utils.Coordinate{X: 10000000000000, Y: 10000000000000})
			}
		}

		res += solveMachine(i, machine)
		resPart2 += solveMachine(i, machinePart2)
	}

	fmt.Println(res)
	fmt.Println(resPart2)
}

func solveMachine(i int, machine Machine) int {
	ratioAtoB := float64(machine.Prize.Y*machine.ButtonB.X-machine.Prize.X*machine.ButtonB.Y) / float64(machine.Prize.X*machine.ButtonA.Y-machine.Prize.Y*machine.ButtonA.X)

	numberOfB := math.Round(float64(machine.Prize.X) / (float64(machine.ButtonA.X)*ratioAtoB + float64(machine.ButtonB.X)))
	numberOfA := math.Round(ratioAtoB * numberOfB)

	newX := machine.ButtonA.X*int(numberOfA) + machine.ButtonB.X*int(numberOfB)
	newY := machine.ButtonA.Y*int(numberOfA) + machine.ButtonB.Y*int(numberOfB)

	if newX == machine.Prize.X && newY == machine.Prize.Y {
		fmt.Println("Machine", i, "solved with", int(numberOfA), "of A and", int(numberOfB), "of B")
		return int(numberOfA*3) + int(numberOfB)
	} else {
		fmt.Println("Machine", i, "unsolvable")
		return 0
	}
}

func parseCoordinate(line string) utils.Coordinate {
	usefulData := strings.Split(line, ": ")[1]

	coords := strings.Split(usefulData, ", ")
	x, err := strconv.Atoi(coords[0][2:])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coords[1][2:])
	if err != nil {
		panic(err)
	}

	return utils.Coordinate{X: x, Y: y}
}
