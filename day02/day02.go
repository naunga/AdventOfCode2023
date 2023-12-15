package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Day02 struct{}

type Bag struct {
	red   int
	blue  int
	green int
}

func GetValue(bag Bag, input string) int {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	pattern := regexp.MustCompile(`(\d+)\s+(red|green|blue)`)

	gameName, handfuls := strings.Split(input, ":")[0], strings.Split(input, ":")[1]

	gameID, err := strconv.Atoi(strings.Split(gameName, " ")[1])
	if err != nil {
		log.Fatalf("Could not convert '%s' to int: %s", gameName, err)
	}

	for _, handful := range strings.Split(handfuls, ";") {
		for _, revealed := range strings.Split(handful, ",") {
			matches := pattern.FindStringSubmatch(revealed)

			if matches != nil {
				count, err := strconv.Atoi(matches[1])
				if err != nil {
					log.Fatalf("Could not convert '%s' to int: %s", matches[1], err)
				}
				color := matches[2]

				switch color {
				case "red":
					if count > bag.red {
						return 0
					}
				case "green":
					if count > bag.green {
						return 0
					}
				case "blue":
					if count > bag.blue {
						return 0
					}
				}
			}

		}

	}

	return gameID
}

func (d Day02) Run(fileName string) {
	result := 0

	var bagContents = Bag{red: 12, green: 13, blue: 14}

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
		result += GetValue(bagContents, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part01: %d\n", result)
}
