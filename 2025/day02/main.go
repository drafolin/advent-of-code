package day02

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Main() {
	f, err := os.ReadFile("day_02/input")
	if err != nil {
		panic(err)
	}

	s := string(f)

	trimmed := s[:len(s)-1]

	ranges := strings.Split(trimmed, ",")

	timeStart := time.Now()
	res := firstPart(ranges)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))
	fmt.Println("First part result: ", res)

	timeStart = time.Now()
	res = secondPart(ranges)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
	fmt.Println("Second part result: ", res)
}

func firstPart(ranges []string) (total int) {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for _, ids := range ranges {
		minText, maxText, _ := strings.Cut(ids, "-")

		min, _ := strconv.Atoi(minText)
		max, _ := strconv.Atoi(maxText)

		wg.Go(func() {
			for i := min; i <= max; i++ {
				digits := int(math.Log10(float64(i))) + 1
				if digits%2 == 1 {
					continue
				}

				divisor := int(math.Pow10(digits/2) + 1)

				if i%divisor == 0 {
					ch <- i
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		total += i
	}

	return
}

func secondPart(ranges []string) (total int) {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for _, ids := range ranges {
		minText, maxText, _ := strings.Cut(ids, "-")

		min, _ := strconv.Atoi(minText)
		max, _ := strconv.Atoi(maxText)

		wg.Go(func() {
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

					ch <- i
					break lengthsLoop
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		total += i
	}

	return
}
