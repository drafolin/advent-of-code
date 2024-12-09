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
