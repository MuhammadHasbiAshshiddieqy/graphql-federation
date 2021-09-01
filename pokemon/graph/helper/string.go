package helper

import (
	"strings"
)

func findObj(inpt string) string {
	lastIdx := strings.LastIndex(inpt, ".")
	if lastIdx == -1 {
		return inpt
	}
	return inpt[:lastIdx]
}
