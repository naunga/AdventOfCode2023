package main

import (
	"AdventOfCode/day01"
	"AdventOfCode/day02"
	"fmt"
	"os"
)

type Day interface {
	Run(fileName string)
}

var days = map[string]Day{
	"day01": day01.Day01{},
	"day02": day02.Day02{},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a day to run.")
		os.Exit(1)
	}

	dayToRun := os.Args[1]

	fileName := fmt.Sprintf("%s/input.txt", dayToRun)

	if day, ok := days[dayToRun]; ok {
		day.Run(fileName)
	} else {
		fmt.Println("Please provide a valid day to run.")
		os.Exit(1)
	}
}
