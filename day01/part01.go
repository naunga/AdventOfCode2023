package day01

import (
	"regexp"
	"strconv"
)

func GetPart01Value(input string) int {
	var result int = 0
	pattern := regexp.MustCompile(`(\d)`)

	matches := pattern.FindAllString(input, -1)

	if matches != nil {
		if len(matches) == 1 {
			result, _ = strconv.Atoi(matches[0] + matches[0])
		} else {
			result, _ = strconv.Atoi(matches[0] + matches[len(matches)-1])
		}
	}

	return result
}
