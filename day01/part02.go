package day01

import (
	"log"
	"regexp"
	"strconv"
)

var wordToNumberMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func GetPart02Value(input string) int {
	var result int = 0
	var matches []string
	pattern := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|zero|\d)`)

	for i := 0; ; {
		stringIndex := pattern.FindStringIndex(input[i:])
		if stringIndex == nil {
			break
		}

		start := stringIndex[0] + i
		end := stringIndex[1] + i

		matches = append(matches, input[start:end])

		i = start + 1
	}

	if matches != nil {
		// Check if any of the matches are spelled out numbers, if so, convert them to digits
		for i, match := range matches {
			if numericValue, ok := wordToNumberMap[match]; ok {
				matches[i] = numericValue
			}
		}

		if len(matches) == 1 {
			result, _ = strconv.Atoi(matches[0] + matches[0])
		} else {
			result, _ = strconv.Atoi(matches[0] + matches[len(matches)-1])
		}
	}

	log.Printf("Input: %s, Result: %d", input, result)

	return result
}
