package utils

import (
	"strings"
)

func StrToLineList(in string) []string {
	return strings.Split(in, "\n")
}
