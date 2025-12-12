package day12

import "testing"

const s = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`

func TestPart1(t *testing.T) {
	areas, presents := parseInput(s)
	res := firstPart(areas, presents)
	if res != 2 {
		t.Errorf("First part is %d, expected %d", res, 2)
	}
}
