package day_02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Main() {
	f, err := os.ReadFile("day_02/input")
	if err != nil {
		panic(err)
	}

	s := string(f)
	// 	s = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124
	// `

	trimmed := s[:len(s)-1]

	ranges := strings.Split(trimmed, ",")

	timeStart := time.Now()
	firstPart(ranges)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))

	timeStart = time.Now()
	secondPart(ranges)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
}

func firstPart(ranges []string) {
	total := 0

	for _, ids := range ranges {
		minText, maxText, _ := strings.Cut(ids, "-")

		min, _ := strconv.Atoi(minText)
		max, _ := strconv.Atoi(maxText)

		for i := min; i <= max; i++ {
			id := strconv.Itoa(i)

			if id[0:len(id)/2] == id[len(id)/2:] {
				total += i
			}
		}
	}

	fmt.Println(total)
}

func secondPart(ranges []string) {
	total := 0

	for _, ids := range ranges {
		minText, maxText, _ := strings.Cut(ids, "-")

		min, _ := strconv.Atoi(minText)
		max, _ := strconv.Atoi(maxText)

		for i := min; i <= max; i++ {
			id := strconv.Itoa(i)

		lengthsLoop:
			for length := 1; length <= len(id)/2; length++ {
				if (len(id) % (length)) != 0 {
					continue
				}

				substr := id[0:length]
				for index := length; index <= len(id)-length; index += length {
					if id[index:index+length] != substr {
						continue lengthsLoop
					}
				}

				total += i
				break lengthsLoop
			}
		}
	}

	fmt.Println(total)
}
