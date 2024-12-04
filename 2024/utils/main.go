package utils

import (
	"os"
	"strings"
)

func ReadInput(day string) string {
	data, err := os.ReadFile(day + "/input")
	dataString := string(data)

	if err != nil {
		panic(err)
	}

	if dataString[len(dataString)-1] == '\n' {
		dataString = dataString[:len(dataString)-1]
	}

	return dataString
}

func ReadInputSplitted(day string) []string {
	str := ReadInput(day)
	splitted := strings.Split(str, "\n")

	if splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}

	return splitted
}

func StrToLineList(in string) []string {
	return strings.Split(in, "\n")
}

func StrToGrid(in string) [][]rune {
	splitted := StrToLineList(in)
	res := make([][]rune, len(splitted))

	for i, line := range splitted {
		res[i] = []rune(line)
	}

	return res
}
