package utils

import "strings"

func StringIsEmpty(str string) bool {
	return len(str) == 0
}

func StringEqual(s, t string) bool {
	return strings.EqualFold(s, t)
}
