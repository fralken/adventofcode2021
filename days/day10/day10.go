package day10

import (
	"aoc2021/utils"
	"sort"
	"strings"
)

func FirstStar() {
	utils.Star(10, 1, "score of syntax errors", firstStar)
}

func SecondStar() {
	utils.Star(10, 2, "middle score", secondStar)
}

func firstStar(content string) int {
	lines := strings.Split(content, "\n")
	score := 0
	for _, line := range lines {
		score += parseForErrors(line)
	}
	return score
}

func secondStar(content string) int {
	lines := strings.Split(content, "\n")
	var scores []int
	for _, line := range lines {
		score := parseForValid(line)
		if score > 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

var open = map[rune]bool{
	'(': true,
	'[': true,
	'{': true,
	'<': true,
}

var pair = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var score = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var valid = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func parseForErrors(line string) int {
	var stack []rune
	for _, c := range line {
		if open[c] {
			stack = append(stack, c)
		} else if len(stack) == 0 || pair[c] != stack[len(stack)-1] {
			return score[c]
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return 0
}

func parseForValid(line string) int {
	var stack []rune
	for _, c := range line {
		if open[c] {
			stack = append(stack, c)
		} else if len(stack) == 0 || pair[c] != stack[len(stack)-1] {
			return 0
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		score = score*5 + valid[stack[i]]
	}
	return score
}
