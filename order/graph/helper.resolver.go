package graph

import (
	"regexp"
	"strings"
	"strconv"
)

// Key - ContextKey
type Key struct{}

func fixProfileID(profID string) (int) {
	findProfID := regexp.MustCompile(`"profile_id":?`)
	spltr := regexp.MustCompile(`[,}]`)
	splittedHeader := spltr.Split(profID, -1)

	for _, s := range(splittedHeader) {
		profIdx := findProfID.FindStringIndex(s)
		if (profIdx!=nil) {
			profileID := strings.TrimSpace(s[profIdx[1]:len(s)])
			profileIDInt, _ := strconv.Atoi(profileID)
			return profileIDInt
		}
	}

	return 0
}

func intToString(num int) (string) {
	return strconv.Itoa(num)
}
