package day02

import (
	"strings"
	"testing"
)

const testString = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {
	ranges := strings.Split(testString, ",")
	res := firstPart(ranges)
	if res != 1227775554 {
		t.Errorf("First part is %d, expected %d", res, 1227775554)
	}
}

func TestPart2(t *testing.T) {
	ranges := strings.Split(testString, ",")
	res := secondPart(ranges)
	if res != 4174379265 {
		t.Errorf("First part is %d, expected %d", res, 4174379265)
	}
}
