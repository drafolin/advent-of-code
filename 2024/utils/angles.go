package utils

type Angle int

func NewAngle(val int) Angle {
	return Angle(val % 360)
}
