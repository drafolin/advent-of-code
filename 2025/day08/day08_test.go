package day08

import (
	"strconv"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

	jboxes := make([]*jbox, 0)
	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		jboxes = append(jboxes, &jbox{Position: coordinate3{X: x, Y: y, Z: z}})
	}

	result := firstPart(jboxes, 10)
	if result != 40 {
		t.Errorf("First part result is %d, expected 40", result)
	}
}

func TestPartTwo(t *testing.T) {
	input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

	jboxes := make([]*jbox, 0)
	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		jboxes = append(jboxes, &jbox{Position: coordinate3{X: x, Y: y, Z: z}})
	}

	result := secondPart(jboxes)
	if result != 25272 {
		t.Errorf("Second part result is %d, expected 25272", result)
	}
}
