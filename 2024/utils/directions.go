package utils

type Direction int

const (
	Up Direction = iota
	UpLeft
	Left
	DownLeft
	Down
	DownRight
	Right
	UpRight
)

func (dir Direction) Opposite() Direction {
	switch dir {
	case Up:
		return Down
	case UpLeft:
		return DownRight
	case Left:
		return Right
	case DownLeft:
		return UpRight
	case Down:
		return Up
	case DownRight:
		return UpLeft
	case Right:
		return Left
	case UpRight:
		return DownLeft
	}

	return -1
}

func (dir Direction) GetMatrix() [2]int {
	switch dir {
	case Up:
		return [2]int{0, -1}
	case UpLeft:
		return [2]int{-1, -1}
	case Left:
		return [2]int{-1, 0}
	case DownLeft:
		return [2]int{-1, 1}
	case Down:
		return [2]int{0, 1}
	case DownRight:
		return [2]int{1, 1}
	case Right:
		return [2]int{1, 0}
	case UpRight:
		return [2]int{1, -1}
	}

	return [2]int{}
}

func (dir Direction) Rotate(ang Angle) Direction {
	if ang%45 != 0 {
		panic("Angle must be a multiple of 45.")
	}

	if ang < 0 {
		ang += 360
	}

	return Direction((int(dir) + int(ang)/45) % 8)
}

func DirectionFromRune(r rune) Direction {
	switch r {
	case '^':
		return Up
	case 'v', 'V':
		return Down
	case '<':
		return Left
	case '>':
		return Right
	}

	return -1
}
