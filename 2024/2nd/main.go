package main

import (
	"github.com/drafolin/advent-of-code/2024/utils"
	"math"
	"strconv"
	"strings"
)

func main() {
	//data := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"
	data := utils.ReadInput("2nd")

	reportsStrs := strings.Split(data, "\n")
	reports := make([][]int, len(reportsStrs))

	for i, report := range reportsStrs {
		report := strings.Split(report, " ")
		reports[i] = make([]int, len(report))
		for j, val := range report {
			var err error
			reports[i][j], err = strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
		}
	}

	println(computeSafetyScore(reports))
}

func computeSafetyScore(val [][]int) int {
	c := make(chan int)

	for _, report := range val {
		go func() {
			res := make(chan bool)

			go isSafe(report, res)

			if <-res {
				c <- 1
				return
			}

			for i := range report {
				reportCpy := make([]int, len(report))
				copy(reportCpy, report)
				slicedReport := append(reportCpy[:i], reportCpy[i+1:]...)
				go isSafe(slicedReport, res)
			}

			for i := 0; i < len(report); i++ {
				if <-res {
					c <- 1
					return
				}
			}

			c <- 0
		}()
	}

	cnt := 0

	for v := 0; v < len(val); v++ {
		cnt += <-c
	}

	close(c)

	return cnt
}

func isSafe(report []int, res chan bool) {
	var increasing bool

	if report[0] > report[1] {
		increasing = false
	} else {
		increasing = true
	}

	for i, val := range report[1:] {
		if increasing && val < report[i] {
			res <- false
			return
		}

		if !increasing && val > report[i] {
			res <- false
			return
		}

		if val == report[i] {
			res <- false
			return
		}

		if math.Abs(float64(val-report[i])) > 3 {
			res <- false
			return
		}
	}

	res <- true
}
