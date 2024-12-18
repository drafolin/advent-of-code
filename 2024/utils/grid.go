package utils

import "fmt"

type Grid [][]rune
type TypedGrid[E any] [][]E

func NewGrid(width, height int) Grid {
	res := make(Grid, height)

	for i := range res {
		res[i] = make([]rune, width)
	}

	return res
}

func NewTypedGrid[E any](width, height int, defaultValue E) (res TypedGrid[E]) {
	res = make(TypedGrid[E], height)

	for i := range res {
		res[i] = make([]E, width)
		for j := range res[i] {
			res[i][j] = defaultValue
		}
	}

	return
}

func NewTypedGridFunc[E any](width, height int, f func(x, y int) E) (res TypedGrid[E]) {
	res = make(TypedGrid[E], height)

	for i := range res {
		res[i] = make([]E, width)
		for j := range res[i] {
			res[i][j] = f(j, i)
		}
	}

	return
}

func StrToGrid(in string) Grid {
	splitted := StrToLineList(in)
	res := make(Grid, len(splitted))

	for i, line := range splitted {
		res[i] = []rune(line)
	}

	return res
}

func CopyGrid(in Grid) Grid {
	res := make(Grid, len(in))

	for i, line := range in {
		res[i] = make([]rune, len(line))
		copy(res[i], line)
	}

	return res
}

func (g Grid) At(c Coordinate) rune {
	return g[c.Y][c.X]
}

func (g TypedGrid[E]) HasCoord(c Coordinate) bool {
	return c.Y >= 0 && c.Y < len(g) && c.X >= 0 && c.X < len(g[0])
}

func (g TypedGrid[E]) At(c Coordinate) E {
	return g[c.Y][c.X]
}

func (g TypedGrid[E]) PrintFunc(f func(x, y int) string) {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			fmt.Print(f(x, y))
		}
		fmt.Print("\n")
	}
}

func (g TypedGrid[E]) Width() int {
	return len(g[0])
}

func (g TypedGrid[E]) Height() int {
	return len(g)
}

func (g Grid) GetQuadrant(c Coordinate) (Direction, bool) {
	if c.X > len(g[0])/2 {
		if c.Y > len(g)/2 {
			return DownRight, true
		} else if c.Y < len(g)/2 {
			return UpRight, true
		} else {
			return Right, false
		}
	} else if c.X < len(g[0])/2 {
		if c.Y > len(g)/2 {
			return DownLeft, true
		} else if c.Y < len(g)/2 {
			return UpLeft, true
		} else {
			return Left, false
		}
	} else {
		if c.Y > len(g)/2 {
			return Down, false
		} else if c.Y < len(g)/2 {
			return Up, false
		} else {
			return -1, false
		}
	}
}

func (g Grid) Width() int {
	return len(g[0])
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Print(strs map[Coordinate]string) {
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if val, ok := strs[Coordinate{X: x, Y: y}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (g Grid) Index(target rune) (Coordinate, bool) {
	for y, line := range g {
		for x, val := range line {
			if val == target {
				return Coordinate{X: x, Y: y}, true
			}
		}
	}

	return Coordinate{}, false
}
