package main

import (
	"os"
	"strconv"

	"github.com/drafolin/advent-of-code/2025/day_01"
	"github.com/drafolin/advent-of-code/2025/day_02"
)

func main() {
	day, _ := strconv.Atoi(os.Args[1])

	switch day {
	case 1:
		day_01.Main()
	case 2:
		day_02.Main()
	}
}
