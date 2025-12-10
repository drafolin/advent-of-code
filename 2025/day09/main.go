package day09

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type coordinate struct {
	X, Y int
}

func coordinateFromString(s string) coordinate {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coordinate{X: x, Y: y}
}

func Main() {
	f, err := os.ReadFile("day_09/input")
	if err != nil {
		panic(err)
	}

	s := string(f)

	s = s[:len(s)-1]

	coordinates := make([]coordinate, 0)
	for line := range strings.SplitSeq(s, "\n") {
		coordinates = append(coordinates, coordinateFromString(line))
	}

	timeStart := time.Now()
	res := firstPart(coordinates)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))
	fmt.Println("First part result:", res)

	timeStart = time.Now()
	res = secondPart(coordinates)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
	fmt.Println("Second part result:", res)
}

func firstPart(coordinates []coordinate) int {
	maxVal := 0.0
	for i, coordinate := range coordinates {
		for _, otherCoordinate := range coordinates[i+1:] {
			// add one because the border is included
			deltaX := math.Abs(float64(coordinate.X-otherCoordinate.X)) + 1
			deltaY := math.Abs(float64(coordinate.Y-otherCoordinate.Y)) + 1

			maxVal = max(maxVal, deltaX*deltaY)
		}
	}

	return int(maxVal)
}

func secondPart(coordinates []coordinate) int {
	maxVal := 0.0
	for i, coord := range coordinates {
	coordLoop:
		for _, otherCoordinate := range coordinates[i+1:] {
			minX := min(coord.X, otherCoordinate.X)
			maxX := max(coord.X, otherCoordinate.X)
			minY := min(coord.Y, otherCoordinate.Y)
			maxY := max(coord.Y, otherCoordinate.Y)

			i := len(coordinates) - 1
			for j := range coordinates {
				if coordinates[i] == coord || coordinates[i] == otherCoordinate ||
					coordinates[j] == coord || coordinates[j] == otherCoordinate {
					i = j
					continue
				}

				// check if edge (i, j) intersects the rectangle
				minEdgeX := min(coordinates[i].X, coordinates[j].X)
				maxEdgeX := max(coordinates[i].X, coordinates[j].X)
				minEdgeY := min(coordinates[i].Y, coordinates[j].Y)
				maxEdgeY := max(coordinates[i].Y, coordinates[j].Y)

				if minEdgeX < maxX && maxEdgeX > minX && minEdgeY < maxY && maxEdgeY > minY {
					continue coordLoop
				}

				i = j
			}

			// add one because the border is included
			deltaX := math.Abs(float64(coord.X-otherCoordinate.X)) + 1
			deltaY := math.Abs(float64(coord.Y-otherCoordinate.Y)) + 1

			maxVal = max(maxVal, deltaX*deltaY)
		}
	}

	return int(maxVal)
}
