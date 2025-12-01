package main

import (
	"os"
	"strconv"

	"github.com/drafolin/advent-of-code/2025/day_01"
)

func main() {
	day, _ := strconv.Atoi(os.Args[1])

	switch day {
	case 1:
		day_01.Main()
	}
}
