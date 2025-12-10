package day03

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Main() {
	f, err := os.ReadFile("day_03/input")
	if err != nil {
		panic(err)
	}

	s := string(f)

	s = s[:len(s)-1]

	lines := strings.Split(s, "\n")

	timeStart := time.Now()
	res := firstPart(lines)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))
	fmt.Println("First part result: ", res)

	timeStart = time.Now()
	res = secondPart(lines)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
	fmt.Println("Second part result: ", res)
}

func firstPart(lines []string) (total int) {
	for _, line := range lines {
		tenths := -1
		units := -1

		for i, n := range line {
			n, _ := strconv.Atoi(string(n))

			if n > tenths && i < len(line)-1 {
				tenths = n
				units = -1
				continue
			}

			if n > units {
				units = n
			}
		}

		total += tenths*10 + units
	}

	return
}

/* No parallelization, because benchmark shows that making the threads is slower than just doing it sequentially. */
func secondPart(lines []string) (total int) {
	const JoltageDigits = 12

	for _, bank := range lines {
		bankJoltages := make([]int, JoltageDigits)
		// The last index that would make a sequence long enough if the next batteries are all turned on.
		lastBatteryWithoutSkip := len(bank) - JoltageDigits

		for batteryIndex, thisBatteryJoltage := range bank {
			// The first index that would make the tail short enough if the next batteries are all turned on.
			// If negative, means the sequence is already long enough, so we start from 0.
			firstValidIndex := max(batteryIndex-lastBatteryWithoutSkip, 0)

			n, _ := strconv.Atoi(string(thisBatteryJoltage))

			for j := firstValidIndex; j < JoltageDigits; j++ {
				if n > bankJoltages[j] {
					bankJoltages[j] = n

					for k := j + 1; k < JoltageDigits; k++ {
						// Time optimization: if the next battery is already 0, then all of the following positions never got touched.
						if bankJoltages[k] == 0 {
							break
						}

						bankJoltages[k] = 0
					}

					break
				}
			}
		}

		bankJoltage := 0
		for j := range JoltageDigits {
			bankJoltage += bankJoltages[j] * int(math.Pow10(JoltageDigits-1-j))
		}
		total += bankJoltage
	}

	return
}
