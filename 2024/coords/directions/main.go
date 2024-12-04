package directions

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

func (direction Direction) Opposite() Direction {
	switch direction {
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
