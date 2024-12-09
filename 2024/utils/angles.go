package utils

type Angle int

func New(val int) Angle {
	return Angle(val % 360)
}
