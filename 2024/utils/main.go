package utils

import (
	"os"
)

func ReadInput(day string) string {
	data, err := os.ReadFile(day + "/input")
	dataString := string(data)

	if err != nil {
		panic(err)
	}

	return dataString
}
