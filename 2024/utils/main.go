package utils

import (
	"strconv"
	"strings"
)

func StrToLineList(in string) []string {
	return strings.Split(in, "\n")
}

func StrToNumberList(in string) ([]int, error) {
	res := make([]int, 0)
	for _, n := range strings.Split(in, " ") {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}

	return res, nil
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
