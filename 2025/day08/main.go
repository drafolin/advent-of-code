package day08

import (
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type coordinate3 struct {
	X int
	Y int
	Z int
}

func (c coordinate3) Distance(c2 coordinate3) float64 {
	return math.Sqrt(
		math.Pow(float64(c.X-c2.X), 2) +
			math.Pow(float64(c.Y-c2.Y), 2) +
			math.Pow(float64(c.Z-c2.Z), 2))
}

type jbox struct {
	Position coordinate3
}

func Main() {
	f, err := os.ReadFile("day_08/input")
	if err != nil {
		panic(err)
	}

	s := string(f)

	s = s[:len(s)-1]

	lines := strings.Split(s, "\n")

	jboxes := make([]*jbox, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		jboxes = append(jboxes, &jbox{Position: coordinate3{X: x, Y: y, Z: z}})
	}

	timeStart := time.Now()
	result := firstPart(jboxes, 1000)
	fmt.Println("First part result:", result)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))

	timeStart = time.Now()
	result = secondPart(jboxes)
	fmt.Println("Second part result:", result)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
}

type tuple[T any] struct {
	A, B T
}

func firstPart(jboxes []*jbox, count int) int {
	distances := make(map[tuple[*jbox]]float64)
	for i, box := range jboxes {
		for _, jbox2 := range jboxes[i+1:] {
			if box == jbox2 {
				continue
			}
			distances[tuple[*jbox]{A: box, B: jbox2}] = box.Position.Distance(jbox2.Position)
		}
	}

	distanceList := slices.Collect(maps.Keys(distances))
	slices.SortFunc(distanceList, func(a, b tuple[*jbox]) int {
		if distances[a] > distances[b] {
			return 1
		}
		if distances[a] < distances[b] {
			return -1
		}
		return 0
	})

	circuits := make([][]*jbox, 0)
	for i := range count {
		var matchingCircuits [][]*jbox
		var indices []int

		for index, c := range circuits {
			if slices.Contains(c, distanceList[i].A) || slices.Contains(c, distanceList[i].B) {
				matchingCircuits = append(matchingCircuits, c)
				indices = append(indices, index)
			}
		}

		if len(matchingCircuits) == 0 {
			matchingCircuits = [][]*jbox{{distanceList[i].A, distanceList[i].B}}
			indices = []int{len(circuits)}
		}

		if len(matchingCircuits) > 1 {
			newCircuit := make([]*jbox, 0)
			for _, circuit := range matchingCircuits {
				newCircuit = append(newCircuit, circuit...)
			}

			newCircuits := make([][]*jbox, 0)
			for i := range circuits {
				if !slices.Contains(indices, i) {
					newCircuits = append(newCircuits, circuits[i])
				}
			}
			circuits = append(newCircuits, newCircuit)
		} else {
			circuit := matchingCircuits[0]
			j := indices[0]

			if !slices.Contains(circuit, distanceList[i].A) {
				circuit = append(circuit, distanceList[i].A)
			}

			if !slices.Contains(circuit, distanceList[i].B) {
				circuit = append(circuit, distanceList[i].B)
			}

			if j >= len(circuits) {
				circuits = append(circuits, circuit)
			} else {
				circuits[j] = circuit
			}
		}
	}

	lengths := make([]int, 0)
	for _, circuit := range circuits {
		lengths = append(lengths, len(circuit))
	}

	slices.SortFunc(lengths, func(a, b int) int {
		return b - a
	})

	total := 1

	for _, length := range lengths[:3] {
		total *= length
	}

	return total
}

func secondPart(jboxes []*jbox) int {
	distances := make(map[tuple[*jbox]]float64)
	for i, box := range jboxes {
		for _, jbox2 := range jboxes[i+1:] {
			if box == jbox2 {
				continue
			}
			distances[tuple[*jbox]{A: box, B: jbox2}] = box.Position.Distance(jbox2.Position)
		}
	}

	distanceList := slices.Collect(maps.Keys(distances))
	slices.SortFunc(distanceList, func(a, b tuple[*jbox]) int {
		if distances[a] > distances[b] {
			return 1
		}
		if distances[a] < distances[b] {
			return -1
		}
		return 0
	})

	circuits := make([][]*jbox, 0)
	i := 0
	for {
		var matchingCircuits [][]*jbox
		var indices []int

		for index, c := range circuits {
			if slices.Contains(c, distanceList[i].A) || slices.Contains(c, distanceList[i].B) {
				matchingCircuits = append(matchingCircuits, c)
				indices = append(indices, index)
			}
		}

		if len(matchingCircuits) == 0 {
			matchingCircuits = [][]*jbox{{distanceList[i].A, distanceList[i].B}}
			indices = []int{len(circuits)}
		}

		if len(matchingCircuits) > 1 {
			newCircuit := make([]*jbox, 0)
			for _, circuit := range matchingCircuits {
				newCircuit = append(newCircuit, circuit...)
			}

			newCircuits := make([][]*jbox, 0)
			for i := range circuits {
				if !slices.Contains(indices, i) {
					newCircuits = append(newCircuits, circuits[i])
				}
			}
			circuits = append(newCircuits, newCircuit)
		} else {
			circuit := matchingCircuits[0]
			j := indices[0]

			if !slices.Contains(circuit, distanceList[i].A) {
				circuit = append(circuit, distanceList[i].A)
			}

			if !slices.Contains(circuit, distanceList[i].B) {
				circuit = append(circuit, distanceList[i].B)
			}

			if j >= len(circuits) {
				circuits = append(circuits, circuit)
			} else {
				circuits[j] = circuit
			}
		}
		if len(circuits) == 1 && len(circuits[0]) == len(jboxes) {
			break
		}
		i++
	}

	return distanceList[i].A.Position.X * distanceList[i].B.Position.X
}
