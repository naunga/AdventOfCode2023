package main

import (
	"bufio"
	"log"
	"os"
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

func GetCalibrationValue(input string) int {
	var result int = 0
	pattern := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|zero|\d)`)

	matches := pattern.FindAllString(input, -1)

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

	return result
}

func main() {
	calibrationValuesSum := 0

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		calibrationValuesSum += GetCalibrationValue(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	log.Printf("Calibration values sum: %d", calibrationValuesSum)
}
