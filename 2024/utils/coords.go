package utils

import (
	"fmt"
	"reflect"
)

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) MoveTowards(dir Direction) Coordinate {
	switch dir {
	case Up:
		c.Y--
	case UpLeft:
		c.Y--
		c.X--
	case Left:
		c.X--
	case DownLeft:
		c.X--
		c.Y++
	case Down:
		c.Y++
	case DownRight:
		c.Y++
		c.X++
	case Right:
		c.X++
	case UpRight:
		c.X++
		c.Y--
	}

	return c
}

func (c Coordinate) IsInGrid(g interface{}) (bool, error) {
	val := reflect.ValueOf(g)
	if val.Kind() != reflect.Slice {
		return false, fmt.Errorf("The provided argument is not a slice.")
	}

	if c.Y >= val.Len() || c.Y < 0 {
		return false, nil
	}

	line := val.Index(c.Y)
	if line.Kind() != reflect.Slice {
		return false, fmt.Errorf("The provided argument is not a grid.")
	}

	if c.X >= line.Len() || c.X < 0 {
		return false, nil
	}

	return true, nil
}

func (c Coordinate) Diff(other Coordinate) Coordinate {
	return Coordinate{X: c.X - other.X, Y: c.Y - other.Y}
}

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{X: c.X + other.X, Y: c.Y + other.Y}
}
