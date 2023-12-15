package day02

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay02(t *testing.T) {
	var bag Bag = Bag{red: 12, blue: 13, green: 14}

	testData := []struct {
		input string
		want  int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 0},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 0},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5},
	}

	expectedSum := 8
	actualSum := 0

	for _, td := range testData {
		expectedGameId := strings.Split(td.input, ":")[0]
		t.Run(fmt.Sprintf("GetValue_%s", expectedGameId), func(t *testing.T) {
			have := GetValue(bag, td.input)
			actualSum += have
			if have != td.want {
				t.Errorf("Expected have to be %d, but got %d", td.want, have)
			}
		})
	}
	if expectedSum != actualSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, actualSum)
	}
}
