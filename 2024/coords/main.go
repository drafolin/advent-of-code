package coords

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/coords/directions"
	"reflect"
)

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) MoveTowards(dir directions.Direction) Coordinate {
	switch dir {
	case directions.Up:
		c.Y--
	case directions.UpLeft:
		c.Y--
		c.X--
	case directions.Left:
		c.X--
	case directions.DownLeft:
		c.X--
		c.Y++
	case directions.Down:
		c.Y++
	case directions.DownRight:
		c.Y++
		c.X++
	case directions.Right:
		c.X++
	case directions.UpRight:
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
