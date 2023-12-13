package main

import (
	"fmt"
	"testing"
)

func TestDay01(t *testing.T) {
	testData := []struct {
		input string
		want  int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"four82nine74", 84},
		{"hlpqrdh3", 33},
		{"two1nine", 29},
	}

	for _, td := range testData {
		t.Run(fmt.Sprintf("GetCalibrationValue_%s", td.input), func(t *testing.T) {
			have := GetCalibrationValue(td.input)
			if have != td.want {
				t.Errorf("Expected have to be %d, got %d", td.want, have)
			}
		})
	}
}
