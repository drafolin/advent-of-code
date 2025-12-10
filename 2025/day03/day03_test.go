package day03

import (
	"strings"
	"testing"
)

const s = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	lines := strings.Split(s, "\n")
	res := firstPart(lines)

	if res != 357 {
		t.Errorf("First part is %d, expected %d", res, 357)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(s, "\n")
	res := secondPart(lines)

	if res != 3121910778619 {
		t.Errorf("Second part is %d, expected %d", res, 3121910778619)
	}
}
