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
