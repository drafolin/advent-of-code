package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
	"strings"
)

type Bot struct {
	Position utils.Coordinate
}

type MapItem struct {
	PositionLeft utils.Coordinate
	Moveable     bool
}

func (mi *MapItem) CanMove(dir utils.Direction, items map[utils.Coordinate]*MapItem) bool {
	if !mi.Moveable {
		return false
	}

	newPos := mi.PositionLeft.MoveTowards(dir)
	newPosRight := newPos.MoveTowards(utils.Right)

	nextItem, ok := items[newPos]
	if ok && nextItem != mi {
		if !nextItem.CanMove(dir, items) {
			return false
		}
	}

	nextItemRight, ok := items[newPosRight]
	if ok && nextItemRight != mi {
		if !nextItemRight.CanMove(dir, items) {
			return false
		}
	}

	return true
}

func (mi *MapItem) Move(dir utils.Direction, items *map[utils.Coordinate]*MapItem) bool {
	if !mi.CanMove(dir, *items) {
		return false
	}

	newPos := mi.PositionLeft.MoveTowards(dir)
	newPosRight := newPos.MoveTowards(utils.Right)

	nextItem, ok := (*items)[newPos]
	if ok && nextItem != mi {
		nextItem.Move(dir, items)
	}

	nextItemRight, ok := (*items)[newPosRight]
	if ok && nextItemRight != mi {
		nextItemRight.Move(dir, items)
	}

	delete(*items, mi.PositionLeft)
	delete(*items, mi.PositionLeft.MoveTowards(utils.Right))
	mi.PositionLeft = newPos
	(*items)[newPos] = mi
	(*items)[newPosRight] = mi

	return true
}

func main() {
	//data := strings.Split("########\n#..O.O.#\n##@.O..#\n#...O..#\n#.#.O..#\n#...O..#\n#......#\n########\n\n<^^>>>vv<v>>v<<", "\n\n")
	//data := strings.Split("##########\n#..O..O.O#\n#......O.#\n#.OO..O.O#\n#..O@..O.#\n#O#..O...#\n#O..O..O.#\n#.OO.O.OO#\n#....O...#\n##########\n\n<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^\nvvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v\n><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<\n<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^\n^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><\n^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^\n>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^\n<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>\n^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>\nv^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^", "\n\n")
	data := strings.Split(utils.ReadInput("15th"), "\n\n")
	//data := strings.Split("#######\n#...#.#\n#.....#\n#..OO@#\n#..O..#\n#.....#\n#######\n\n<vv<<^^<<^^", "\n\n")
	mapStr := data[0]
	commandsStr := data[1]

	roomMap := utils.StrToGrid(mapStr)
	width := roomMap.Width()
	height := roomMap.Height()
	widerRoomMap := make(utils.Grid, height)
	for i := range roomMap {
		widerRoomMap[i] = make([]rune, width*2)
	}
	commands := []rune(commandsStr)

	var bot Bot
	items := map[utils.Coordinate]*MapItem{}

	for y, line := range roomMap {
		for x, char := range line {
			if char == '@' {
				bot = Bot{utils.Coordinate{X: 2 * x, Y: y}}
			}

			if char == 'O' {
				item := new(MapItem)
				item.PositionLeft = utils.Coordinate{X: 2 * x, Y: y}
				item.Moveable = true
				items[utils.Coordinate{X: 2 * x, Y: y}] = item
				items[utils.Coordinate{X: 2*x + 1, Y: y}] = item
			}

			if char == '#' {
				item := new(MapItem)
				item.PositionLeft = utils.Coordinate{X: 2 * x, Y: y}
				item.Moveable = false
				items[utils.Coordinate{X: 2 * x, Y: y}] = item
				items[utils.Coordinate{X: 2*x + 1, Y: y}] = item
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
	}

	sum := 0

	singleItems := make([]utils.Coordinate, 0)
	for _, item := range items {
		if slices.Contains(singleItems, item.PositionLeft) {
			continue
		}

		singleItems = append(singleItems, item.PositionLeft)

		if !item.Moveable {
			continue
		}

		sum += item.PositionLeft.Y*100 + item.PositionLeft.X
	}

	fmt.Println(sum)
}

func printMapItems(items map[utils.Coordinate]*MapItem, bot Bot, widerRoomMap utils.Grid) {
	printMap := make(map[utils.Coordinate]string)

	for pos, item := range items {
		if item.Moveable {
			if item.PositionLeft == pos {
				printMap[pos] = "["
			} else {
				printMap[pos] = "]"
			}
		} else {
			printMap[pos] = "#"
		}
	}

	printMap[bot.Position] = "@"

	widerRoomMap.Print(printMap)
}
