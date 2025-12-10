package day01

import (
	"strings"
	"testing"
)

const testString = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestPart1(t *testing.T) {
	lines := strings.Split(testString, "\n")
	res := firstPart(lines)

	if res != 3 {
		t.Errorf("First part is %d, expected %d", res, 3)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(testString, "\n")
	res := secondPart(lines)

	if res != 6 {
		t.Errorf("Second part is %d, expected %d", res, 6)
	}
}
