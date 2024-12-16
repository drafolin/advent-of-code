package main

import (
	"github.com/drafolin/advent-of-code/2024/utils"
	"math"
	"slices"
)

type Node struct {
	coordinate utils.Coordinate
	distance   int
	parent     *Node
}

func main() {
	//data := utils.StrToGrid("###############\n#.......#....E#\n#.#.###.#.###.#\n#.....#.#...#.#\n#.###.#####.#.#\n#.#.#.......#.#\n#.#.#####.###.#\n#...........#.#\n###.#.#####.#.#\n#...#.....#.#.#\n#.#.#.###.#.#.#\n#.....#...#.#.#\n#.###.#.#.#.#.#\n#S..#.....#...#\n###############")
	// data := utils.StrToGrid("#################\n#...#...#...#..E#\n#.#.#.#.#.#.#.#.#\n#.#.#.#...#...#.#\n#.#.#.#.###.#.#.#\n#...#.#.#.....#.#\n#.#.#.#.#.#####.#\n#.#...#.#.#.....#\n#.#.#####.#.###.#\n#.#.#.......#...#\n#.#.###.#####.###\n#.#.#...#.....#.#\n#.#.#.#####.###.#\n#.#.#.........#.#\n#.#.#.#########.#\n#S#.............#\n#################")
	data := utils.StrToGrid(utils.ReadInput("16th"))

	start, found := data.Index('S')
	if !found {
		panic("Start not found")
	}

	end, found := data.Index('E')
	if !found {
		panic("End not found")
	}

	unvisited := make([]Node, 0)

	for y, line := range data {
		for x, cell := range line {
			if cell != '#' {
				distance := math.MaxInt
				if x == start.X && y == start.Y {
					distance = 0
				}

				unvisited = append(unvisited, Node{coordinate: utils.Coordinate{X: x, Y: y}, distance: distance})
			}
		}
	}

	for {
		if len(unvisited) == 0 {
			break
		}

		current := slices.MinFunc(unvisited, func(i, j Node) int {
			return i.distance - j.distance
		})

		if current.distance == math.MaxInt {
			break
		}

		var previousDirection utils.Direction
		if current.parent != nil {
			previousDirection = utils.DirectionFromVector(current.coordinate.Diff(current.parent.coordinate))
		} else {
			previousDirection = utils.Right
		}

		for _, dir := range []utils.Direction{utils.Up, utils.Left, utils.Down, utils.Right} {
			neighbor := current.coordinate.MoveTowards(dir)
			if neighbor.IsInGrid(data) {
				if data.At(neighbor) != '#' {
					additionalDistance := 1

					if dir != previousDirection {
						additionalDistance = 1001
					}

					if dir == previousDirection.Opposite() {
						additionalDistance = 2001
					}
					neighborComparator := func(node Node) bool {
						return node.coordinate == neighbor
					}

					if !slices.ContainsFunc(unvisited, neighborComparator) {
						continue
					}
					neighborNode := &unvisited[slices.IndexFunc(unvisited, neighborComparator)]

					if current.distance+additionalDistance < neighborNode.distance {
						neighborNode.distance = current.distance + additionalDistance
						neighborNode.parent = &current
					}

					if neighbor == end {
						println("Distance:", neighborNode.distance)
						return
					}
				}
			}
		}

		index := slices.Index(unvisited, current)
		unvisited = append(unvisited[:index], unvisited[index+1:]...)
	}
}
