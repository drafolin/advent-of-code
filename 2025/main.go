package main

import (
	"os"

	"github.com/drafolin/advent-of-code/2025/day_01"
	"github.com/drafolin/advent-of-code/2025/day_02"
	"github.com/drafolin/advent-of-code/2025/day_03"
	"github.com/drafolin/advent-of-code/2025/day_04"
	"github.com/drafolin/advent-of-code/2025/day_05"
)

func main() {
	day := os.Args[1]

	switch day {
	case "day_01":
		day_01.Main()
	case "day_02":
		day_02.Main()
	case "day_03":
		day_03.Main()
	case "day_04":
		day_04.Main()
	case "day_05":
		day_05.Main()
	}
}
