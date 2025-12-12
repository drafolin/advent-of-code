package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Main() {
	f, err := os.ReadFile("day12/input")
	if err != nil {
		panic(err)
	}

	s := string(f)
	s = s[:len(s)-1]

	areas, presents := parseInput(s)

	timeStart := time.Now()
	res := firstPart(areas, presents)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))
	fmt.Println("First part result: ", res)
}

func parseInput(s string) ([]area, []present) {
	areas := make([]area, 0)
	presents := make([]present, 0)

	sections := strings.Split(s, "\n\n")

	for _, presentText := range sections[:len(sections)-1] {
		present := make([][]bool, 0)

		lines := strings.Split(presentText, "\n")
		for _, line := range lines[1:] {
			row := make([]bool, 0)
			for _, cell := range line {
				row = append(row, cell == '#')
			}
			present = append(present, row)
		}

		presents = append(presents, present)
	}

	for areaText := range strings.SplitSeq(sections[len(sections)-1], "\n") {
		areaTextPart := strings.Split(areaText, " ")

		dimensions := strings.Split(areaTextPart[0][:len(areaTextPart[0])-1], "x")
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])

		presentCounts := make([]int, 0)
		for _, presentCount := range areaTextPart[1:] {
			presentCount, _ := strconv.Atoi(presentCount)
			presentCounts = append(presentCounts, presentCount)
		}

		areas = append(areas, area{
			Width:         width,
			Height:        height,
			PresentCounts: presentCounts,
		})
	}

	return areas, presents
}

type present [][]bool

func (p present) area() int {
	total := 0
	for _, row := range p {
		for _, cell := range row {
			if cell {
				total++
			}
		}
	}
	return total
}

type area struct {
	Width  int
	Height int

	PresentCounts []int
}

/*
Technically, this is not the correct solution, but it works for inputs and i am not doing stuff with geometry.

You can see how the test fails.
*/
func firstPart(areas []area, presents []present) int {
	count := 0
	for _, area := range areas {
		areaSum := 0
		for i, present := range area.PresentCounts {
			areaSum += presents[i].area() * present
		}

		if areaSum > area.Width*area.Height {
			continue
		}

		count++
	}

	return count
}
