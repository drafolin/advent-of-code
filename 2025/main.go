package main

import (
	"os"

	"github.com/drafolin/advent-of-code/2025/day01"
	"github.com/drafolin/advent-of-code/2025/day02"
	"github.com/drafolin/advent-of-code/2025/day03"
	"github.com/drafolin/advent-of-code/2025/day04"
	"github.com/drafolin/advent-of-code/2025/day05"
	"github.com/drafolin/advent-of-code/2025/day06"
	"github.com/drafolin/advent-of-code/2025/day07"
	"github.com/drafolin/advent-of-code/2025/day08"
	"github.com/drafolin/advent-of-code/2025/day09"
	"github.com/drafolin/advent-of-code/2025/day10"
	"github.com/drafolin/advent-of-code/2025/day11"
	"github.com/drafolin/advent-of-code/2025/day12"
)

func main() {
	day := os.Args[1]

	switch day {
	case "day01":
		day01.Main()
	case "day02":
		day02.Main()
	case "day03":
		day03.Main()
	case "day04":
		day04.Main()
	case "day05":
		day05.Main()
	case "day06":
		day06.Main()
	case "day07":
		day07.Main()
	case "day08":
		day08.Main()
	case "day09":
		day09.Main()
	case "day10":
		day10.Main()
	case "day11":
		day11.Main()
	case "day12":
		day12.Main()
	}
}
