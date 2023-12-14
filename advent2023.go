package main

import (
	"AdventOfCode/day01"
	"fmt"
	"os"
)

var days = map[string]func(string){
	"day01": day01.Run,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a day to run.")
		os.Exit(1)
	}

	dayToRun := os.Args[1]

	fileName := fmt.Sprintf("%s/input.txt", dayToRun)

	if day, ok := days[dayToRun]; ok {
		day(fileName)
	} else {
		fmt.Println("Please provide a valid day to run.")
		os.Exit(1)
	}
}
