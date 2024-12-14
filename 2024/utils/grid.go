package utils

type Grid [][]rune

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
