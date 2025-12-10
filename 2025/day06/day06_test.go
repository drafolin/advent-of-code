package day06

import (
	"strings"
	"testing"
)

const s = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +`

func TestPart1(t *testing.T) {
	lines := strings.Split(s, "\n")
	res := firstPart(lines)

	if res != 4277556 {
		t.Errorf("First part is %d, expected %d", res, 4277556)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(s, "\n")
	res := secondPart(lines)

	if res != 3263392 {
		t.Errorf("Second part is %d, expected %d", res, 3263392)
	}
}
