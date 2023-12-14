package day01

import (
	"bufio"
	"log"
	"os"
)

func Run(fileName string) {
	part01Answer := 0

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
		part01Answer += GetPart01Value(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	log.Printf("Part01: Calibration values sum: %d", part01Answer)
}
