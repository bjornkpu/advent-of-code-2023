package day1

import (
	"regexp"
	"strconv"
	"strings"
)

var digitRegexp = regexp.MustCompile("\\d")
var wordRegexp = regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")
var wordRegexpReversed = regexp.MustCompile("\\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin")
var wordToNumber = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
var wordToNumberReversed = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"eno":   1,
	"owt":   2,
	"eerht": 3,
	"ruof":  4,
	"evif":  5,
	"xis":   6,
	"neves": 7,
	"thgie": 8,
	"enin":  9,
}

func Part1(input string) (sum int, err error) {
	return calculateSum(input, findFirstAndLastDigit)
}

func Part2(input string) (sum int, err error) {
	return calculateSum(input, findFirstAndLastWord)
}

func calculateSum(input string, findFirstAndLastFn func(string) (int, int, error)) (sum int, err error) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		first, last, err := findFirstAndLastFn(line)
		if err != nil || first == 0 || last == 0 {
			return 0, err
		}

		value, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		if err != nil {
			return 0, err
		}

		sum += value
	}

	return sum, nil
}

func findFirstAndLastDigit(line string) (first int, last int, err error) {
	return findFirstAndLast(line, digitRegexp, digitRegexp)
}

func findFirstAndLastWord(line string) (first int, last int, err error) {
	return findFirstAndLast(line, wordRegexp, wordRegexpReversed)
}

func findFirstAndLast(line string, findRegexp *regexp.Regexp, findRegexpRev *regexp.Regexp) (first int, last int, err error) {
	firstMatch := findRegexp.FindString(line)
	lastMatch := findRegexpRev.FindString(reverseString(line))

	first = wordToNumber[firstMatch]
	last = wordToNumberReversed[lastMatch]

	return first, last, nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
