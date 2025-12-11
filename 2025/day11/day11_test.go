package day11

import (
	"strings"
	"testing"
)

const s = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

func TestPart1(t *testing.T) {
	lines := strings.Split(s, "\n")

	nodes := make([]node, 0)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		from := parts[0]
		to := strings.Split(parts[1], " ")

		node := node{
			name:     from,
			children: to,
		}

		nodes = append(nodes, node)
	}

	res := firstPart(nodes)

	if res != 5 {
		t.Errorf("First part is %d, expected %d", res, 5)
	}
}

const s2 = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`

func TestPart2(t *testing.T) {
	lines := strings.Split(s2, "\n")

	nodes := make([]node, 0)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		from := parts[0]
		to := strings.Split(parts[1], " ")

		node := node{
			name:     from,
			children: to,
		}

		nodes = append(nodes, node)
	}

	res := secondPart(nodes)

	if res != 2 {
		t.Errorf("Second part is %d, expected %d", res, 2)
	}
}
