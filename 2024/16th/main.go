package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"math"
	"slices"
)

type Node struct {
	coordinate utils.Coordinate
	direction  utils.Direction
	distance   int
	parents    []*Node
}

func main() {
	//data := utils.StrToGrid("###############\n#.......#....E#\n#.#.###.#.###.#\n#.....#.#...#.#\n#.###.#####.#.#\n#.#.#.......#.#\n#.#.#####.###.#\n#...........#.#\n###.#.#####.#.#\n#...#.....#.#.#\n#.#.#.###.#.#.#\n#.....#...#.#.#\n#.###.#.#.#.#.#\n#S..#.....#...#\n###############")
	//data := utils.StrToGrid("#################\n#...#...#...#..E#\n#.#.#.#.#.#.#.#.#\n#.#.#.#...#...#.#\n#.#.#.#.###.#.#.#\n#...#.#.#.....#.#\n#.#.#.#.#.#####.#\n#.#...#.#.#.....#\n#.#.#####.#.###.#\n#.#.#.......#...#\n#.#.###.#####.###\n#.#.#...#.....#.#\n#.#.#.#####.###.#\n#.#.#.........#.#\n#.#.#.#########.#\n#S#.............#\n#################")
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
				for _, dir := range []utils.Direction{utils.Up, utils.Left, utils.Down, utils.Right} {
					if x == start.X && y == start.Y && dir == utils.Right {
						distance = 0
					}

					unvisited = append(unvisited, Node{
						coordinate: utils.Coordinate{X: x, Y: y},
						distance:   distance,
						direction:  dir,
					})
				}
			}
		}
	}

djikstra:
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

		for _, dir := range []int{-90, 0, 90} {
			newNode := current
			var additionalDistance int
			if dir == 0 {
				newNode.coordinate = newNode.coordinate.MoveTowards(newNode.direction)
				additionalDistance = 1
			} else {
				newNode.direction = current.direction.Rotate(utils.NewAngle(dir))
				additionalDistance = 1000
			}

			if newNode.coordinate.IsInGrid(data) {
				if data.At(newNode.coordinate) != '#' {
					neighborComparator := func(node Node) bool {
						return node.coordinate == newNode.coordinate && node.direction == newNode.direction
					}

					var neighborIndex int
					if neighborIndex = slices.IndexFunc(unvisited, neighborComparator); neighborIndex == -1 {
						continue
					}
					neighborNode := &unvisited[neighborIndex]

					if current.distance+additionalDistance < neighborNode.distance {
						neighborNode.distance = current.distance + additionalDistance
						neighborNode.parents = append(neighborNode.parents, &current)
					} else if current.distance+additionalDistance == neighborNode.distance {
						neighborNode.parents = append(neighborNode.parents, &current)
					}

					if neighborNode.coordinate == end {
						fmt.Println("Distance:", neighborNode.distance)
						break djikstra
					}
				}
			}
		}

		index := slices.IndexFunc(unvisited, func(node Node) bool {
			return node.coordinate == current.coordinate && node.direction == current.direction
		})
		unvisited = append(unvisited[:index], unvisited[index+1:]...)
	}

	ancestorsQueue := []*Node{&unvisited[slices.IndexFunc(unvisited, func(node Node) bool {
		return node.coordinate == end
	})]}
	ancestors := make([]*Node, 0)

	for len(ancestorsQueue) > 0 {
		current := ancestorsQueue[0]
		ancestorsQueue = ancestorsQueue[1:]

		if !slices.ContainsFunc(ancestors, func(node *Node) bool {
			return node.coordinate == current.coordinate
		}) {
			ancestors = append(ancestors, current)
		}

		for _, parent := range current.parents {
			ancestorsQueue = append(ancestorsQueue, parent)
		}
	}

	fmt.Println("Ancestors:", len(ancestors))

}
