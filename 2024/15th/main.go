package main

import (
	"github.com/drafolin/advent-of-code/2024/utils"
	"strings"
)

type Bot struct {
	Position utils.Coordinate
}

type MapItem struct {
	Position utils.Coordinate
	Moveable bool
}

func (mi *MapItem) Move(dir utils.Direction, items *map[utils.Coordinate]MapItem) bool {
	if !mi.Moveable {
		return false
	}

	newPos := mi.Position.MoveTowards(dir)

	nextItem, ok := (*items)[newPos]
	if ok {
		if !nextItem.Move(dir, items) {
			return false
		}
	}

	delete(*items, mi.Position)
	mi.Position = newPos
	(*items)[newPos] = *mi

	return true
}

func main() {
	//data := strings.Split("########\n#..O.O.#\n##@.O..#\n#...O..#\n#.#.O..#\n#...O..#\n#......#\n########\n\n<^^>>>vv<v>>v<<", "\n\n")
	//data := strings.Split("##########\n#..O..O.O#\n#......O.#\n#.OO..O.O#\n#..O@..O.#\n#O#..O...#\n#O..O..O.#\n#.OO.O.OO#\n#....O...#\n##########\n\n<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^\nvvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v\n><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<\n<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^\n^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><\n^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^\n>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^\n<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>\n^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>\nv^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^", "\n\n")
	data := strings.Split(utils.ReadInput("15th"), "\n\n")
	mapStr := data[0]
	commandsStr := data[1]

	roomMap := utils.StrToGrid(mapStr)
	commands := []rune(commandsStr)

	var bot Bot
	items := map[utils.Coordinate]MapItem{}

	for y, line := range roomMap {
		for x, char := range line {
			if char == '@' {
				bot = Bot{utils.Coordinate{X: x, Y: y}}
			}

			if char == 'O' {
				items[utils.Coordinate{X: x, Y: y}] = MapItem{utils.Coordinate{X: x, Y: y}, true}
			}

			if char == '#' {
				items[utils.Coordinate{X: x, Y: y}] = MapItem{utils.Coordinate{X: x, Y: y}, false}
			}
		}
	}

	for _, command := range commands {
		dir := utils.DirectionFromRune(command)
		newPos := bot.Position.MoveTowards(dir)

		item, ok := items[newPos]
		canMove := true
		if ok {
			canMove = item.Move(dir, &items)
		}

		if canMove {
			bot.Position = newPos
		}

		/*
			printMap := map[utils.Coordinate]string{}
			for pos, item := range items {
				if item.Moveable {
					printMap[pos] = "O"
				} else {
					printMap[pos] = "#"
				}
			}

			printMap[bot.Position] = "@"

			roomMap.Print(printMap)
		*/
	}
	sum := 0

	for coord, item := range items {
		if !item.Moveable {
			continue
		}

		sum += coord.Y*100 + coord.X
	}

	println(sum)
}
