package main

import (
	"os"

	"github.com/drafolin/advent-of-code/2025/day_01"
	"github.com/drafolin/advent-of-code/2025/day_02"
	"github.com/drafolin/advent-of-code/2025/day_03"
	"github.com/drafolin/advent-of-code/2025/day_04"
	"github.com/drafolin/advent-of-code/2025/day_05"
	"github.com/drafolin/advent-of-code/2025/day_06"
	"github.com/drafolin/advent-of-code/2025/day_07"
	"github.com/drafolin/advent-of-code/2025/day_08"
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
	case "day_06":
		day_06.Main()
	case "day_07":
		day_07.Main()
	case "day_08":
		day_08.Main()
	}
}
