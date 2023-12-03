package day2

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func readFile(t *testing.T, filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Error reading file %s: %s", filename, err)
	}
	return string(content)
}

func TestDay2(t *testing.T) {
	for _, testCase := range []struct {
		name     string
		input    string
		function func(string) (int, error)
		expected int
	}{
		{"Part1_Example", "input_example_part1.txt", Part1, 8},
		{"Part1", "input.txt", Part1, 3035},
		{"Part2_Example", "input_example_part2.txt", Part2, 2286},
		{"Part2", "input.txt", Part2, 66027},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			sum, err := testCase.function(readFile(t, testCase.input))
			if err != nil {
				t.Errorf("Error occurred in function: %v", err)
			}
			assert.Equal(t, testCase.expected, sum)
		})
	}
}
