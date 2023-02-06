package day08

import (
	"aoc2021/utils"
	"reflect"
	"strings"
)

func FirstStar() {
	utils.Star(8, 1, "number of 1, 4, 7, 8 digits", firstStar)
}

func SecondStar() {
	utils.Star(8, 2, "sum of output values", secondStar)
}

func firstStar(content string) int {
	lines := strings.Split(content, "\n")
	count := 0
	for _, line := range lines {
		display := strings.Split(line, " | ")
		output := strings.Fields(display[1])
		for _, digit := range output {
			segs := len(digit)
			if segs == 2 || segs == 4 || segs == 3 || segs == 7 {
				count++
			}
		}
	}
	return count
}

func secondStar(content string) int {
	lines := strings.Split(content, "\n")
	count := 0
	for _, line := range lines {
		display := strings.Split(line, " | ")
		input := strings.Fields(display[0])
		digits := mapDigits(input)
		output := strings.Fields(display[1])
		value := 0
		for _, digit := range output {
			value = 10*value + findDigit(digits, digit)
		}
		count += value
	}
	return count
}

func stringToMap(s string) map[rune]bool {
	set := make(map[rune]bool)
	for _, c := range s {
		set[c] = true
	}
	return set
}

func contains(container map[rune]bool, contained map[rune]bool) bool {
	for r := range contained {
		if !container[r] {
			return false
		}
	}
	return true
}

func mapDigits(input []string) (digits [10]map[rune]bool) {
	for _, digit := range input {
		segs := len(digit)
		if segs == 2 {
			digits[1] = stringToMap(digit)
		} else if segs == 3 {
			digits[7] = stringToMap(digit)
		} else if segs == 4 {
			digits[4] = stringToMap(digit)
		} else if segs == 7 {
			digits[8] = stringToMap(digit)
		}
	}
	for _, digit := range input {
		if len(digit) == 6 {
			set := stringToMap(digit)
			if contains(set, digits[4]) {
				digits[9] = set
			} else if contains(set, digits[7]) {
				digits[0] = set
			} else {
				digits[6] = set
			}
		}
	}
	for _, digit := range input {
		if len(digit) == 5 {
			set := stringToMap(digit)
			if contains(digits[9], set) {
				if contains(set, digits[7]) {
					digits[3] = set
				} else {
					digits[5] = set
				}
			} else {
				digits[2] = set
			}
		}
	}
	return
}

func findDigit(digits [10]map[rune]bool, digit string) int {
	set := stringToMap(digit)
	for i := range digits {
		if reflect.DeepEqual(digits[i], set) {
			return i
		}
	}
	panic("digit not found")
}
