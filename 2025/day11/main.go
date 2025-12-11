package day11

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Main() {
	f, err := os.ReadFile("day11/input")
	if err != nil {
		panic(err)
	}

	s := string(f)
	s = s[:len(s)-1]

	lines := strings.Split(s, "\n")

	nodes := make([]node, 0)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		from := parts[0]
		to := strings.Split(parts[1], " ")
		nodes = append(nodes, node{name: from, children: to})
	}

	timeStart := time.Now()
	res := firstPart(nodes)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))
	fmt.Println("First part result: ", res)

	timeStart = time.Now()
	res = secondPart(nodes)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
	fmt.Println("Second part result: ", res)
}

type node struct {
	name     string
	children []string
}

var cache = map[string]int{
	"out": 1,
}

func dfsFrom(nodes []node, name string) int {
	if val, ok := cache[name]; ok {
		return val
	}

	total := 0
	for _, node := range nodes {
		if node.name == name {
			for _, child := range node.children {
				total += dfsFrom(nodes, child)
			}
		}
	}

	cache[name] = total
	return total
}

func firstPart(nodes []node) int {
	total := 0
	for _, node := range nodes {
		if node.name == "you" {
			for _, child := range node.children {
				total += dfsFrom(nodes, child)
			}
		}
	}

	return total
}

type dacFftValidation int8

const (
	ValidationNone dacFftValidation = 0
	ValidationDAC  dacFftValidation = 1 << iota
	ValidationFFT
)

var dfsWithValidationCache = map[string]map[dacFftValidation]int{
	"out": {
		ValidationDAC | ValidationFFT: 1,
		ValidationFFT:                 0,
		ValidationDAC:                 0,
		ValidationNone:                0,
	},
}

func dfsWithValidationFrom(nodes []node, name string, valid dacFftValidation) int {
	if val, ok := dfsWithValidationCache[name][valid]; ok {
		return val
	}

	switch name {
	case "dac":
		valid |= ValidationDAC
	case "fft":
		valid |= ValidationFFT
	}

	total := 0
	for _, node := range nodes {
		if node.name == name {
			for _, child := range node.children {
				total += dfsWithValidationFrom(nodes, child, valid)
			}
		}
	}

	if _, ok := dfsWithValidationCache[name]; !ok {
		dfsWithValidationCache[name] = make(map[dacFftValidation]int)
	}
	dfsWithValidationCache[name][valid] = total
	return total
}

func secondPart(nodes []node) int {
	total := 0
	for _, node := range nodes {
		if node.name == "svr" {
			for _, child := range node.children {
				total += dfsWithValidationFrom(nodes, child, 0)
			}
		}
	}

	return total
}
