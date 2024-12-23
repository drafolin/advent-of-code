package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
	"sync"
)

type Guard struct {
	Pos utils.Coordinate
	Dir utils.Direction
}

func main() {
	//data := "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...\n"
	data := utils.ReadInput("6th")

	grid := utils.StrToGrid(data)

	var guard Guard

	for y, line := range grid {
		for x, char := range line {
			if char == '^' || char == 'V' || char == '<' || char == '>' {
				guard = Guard{Pos: utils.Coordinate{X: x, Y: y}, Dir: getDirectionFromCharacter(char)}
			}
		}
	}

	type Position struct {
		Coord utils.Coordinate
		Dir   utils.Direction
	}
	visitedPos := []Position{{Coord: guard.Pos, Dir: guard.Dir}}

	for {
		pos, dir, ok := guard.GetNextPos(grid)
		if !ok {
			break
		}

		guard.Pos = pos
		guard.Dir = dir

		if !slices.ContainsFunc(visitedPos, func(pos Position) bool {
			return pos.Coord == guard.Pos
		}) {
			visitedPos = append(visitedPos, Position{pos, dir})
		}
	}

	// Part 1
	fmt.Println(len(visitedPos))

	workingObstacles := 0
	wg := sync.WaitGroup{}

	for i, pos := range visitedPos[1:] {
		wg.Add(1)
		go func(i int, pos Position, workingObstagles *int) {
			defer wg.Done()
			newGrid := utils.CopyGrid(grid)
			newGrid[pos.Coord.Y][pos.Coord.X] = '#'

			guard := Guard{Pos: visitedPos[i].Coord, Dir: visitedPos[i].Dir}
			visitedPosWithObstacle := make([]Position, len(visitedPos))
			copy(visitedPosWithObstacle, visitedPos[:i])

			for {
				pos, dir, ok := guard.GetNextPos(newGrid)
				if !ok {
					break
				}

				guard.Pos = pos
				guard.Dir = dir

				if !slices.Contains(visitedPosWithObstacle, Position{Coord: guard.Pos, Dir: guard.Dir}) {
					visitedPosWithObstacle = append(visitedPosWithObstacle, Position{pos, dir})
				} else {
					workingObstacles++
					break
				}
			}
		}(i, pos, &workingObstacles)
	}

	wg.Wait()

	// Part 2
	fmt.Println(workingObstacles)
}

func getDirectionFromCharacter(char rune) utils.Direction {
	switch char {
	case '^':
		return utils.Up
	case 'V':
		return utils.Down
	case '<':
		return utils.Left
	case '>':
		return utils.Right
	}

	return -1
}

func (guard Guard) GetNextPos(grid utils.Grid) (utils.Coordinate, utils.Direction, bool) {
	nextPos := guard.Pos.MoveTowards(guard.Dir)

	if v, err := nextPos.IsInAnyGrid(grid); !v {
		if err != nil {
			panic(err)
		}

		return nextPos, guard.Dir, false
	}

	if grid[nextPos.Y][nextPos.X] == '#' {
		guard.Dir = guard.Dir.Rotate(-90)
		return guard.GetNextPos(grid)
	}

	return nextPos, guard.Dir, true
}
