package day01

import (
	"aoc2021/utils"
	"strings"
)

func FirstStar() {
	utils.Star(1, 1, "measurements larger than the previous one", firstStar)
}

func SecondStar() {
	utils.Star(1, 2, "measurements of sliding windows larger than the previous one", secondStar)
}

func firstStar(content string) int {
	lines := strings.Split(content, "\n")
	values := utils.StringsToInts(lines)
	count := 0
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			count++
		}
	}
	return count
}

func secondStar(content string) int {
	lines := strings.Split(content, "\n")
	values := utils.StringsToInts(lines)
	count := 0
	previous := values[0] + values[1] + values[2]
	for i := 3; i < len(values); i++ {
		next := values[i] + values[i-1] + values[i-2]
		if next > previous {
			count++
		}
		previous = next
	}
	return count
}
