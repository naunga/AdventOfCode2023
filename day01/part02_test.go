package day01

import (
	"fmt"
	"testing"
)

func TestDay02(t *testing.T) {
	testData := []struct {
		input string
		want  int
	}{
		{"two1nine", 29},
		{"1six15ninebgnzhtbmlxpnrqoneightfhp", 18},
	}

	for _, td := range testData {
		t.Run(fmt.Sprintf("GetCalibrationValue_%s", td.input), func(t *testing.T) {
			have := GetPart02Value(td.input)
			if have != td.want {
				t.Errorf("Expected have to be %d, got %d", td.want, have)
			}
		})
	}
}
