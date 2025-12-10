package day10

import (
	"strings"
	"testing"
)

const s = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

func TestPart1(t *testing.T) {
	lines := strings.Split(s, "\n")
	machines := make([]machine, 0)
	for _, line := range lines {
		machines = append(machines, machineFromString(line))
	}

	res := firstPart(machines)

	if res != 7 {
		t.Errorf("First part is %d, expected %d", res, 7)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(s, "\n")
	machines := make([]machine, 0)
	for _, line := range lines {
		machines = append(machines, machineFromString(line))
	}

	res := secondPart(machines)

	if res != 33 {
		t.Errorf("Second part is %d, expected %d", res, 33)
	}
}

func TestPart2BFS(t *testing.T) {
	lines := strings.Split(s, "\n")
	machines := make([]machine, 0)
	for _, line := range lines {
		machines = append(machines, machineFromString(line))
	}

	res := bfsSecondPart(machines)

	if res != 33 {
		t.Errorf("Second part is %d, expected %d", res, 33)
	}
}
