package day02

import (
	"aoc2021/utils"
	"strconv"
	"strings"
)

func FirstStar() {
	utils.Star(2, 1, "final horizontal position * final depth", firstStar)
}

func SecondStar() {
	utils.Star(2, 2, "final horizontal position * final depth", secondStar)
}

func firstStar(content string) int {
	lines := strings.Split(content, "\n")
	horizontal := 0
	depth := 0
	for _, line := range lines {
		command := strings.Split(line, " ")
		value, _ := strconv.Atoi(command[1])
		if command[0] == "forward" {
			horizontal += value
		} else if command[0] == "down" {
			depth += value
		} else if command[0] == "up" {
			depth -= value
		}
	}
	return horizontal * depth
}

func secondStar(content string) int {
	lines := strings.Split(content, "\n")
	horizontal := 0
	aim := 0
	depth := 0
	for _, line := range lines {
		command := strings.Split(line, " ")
		value, _ := strconv.Atoi(command[1])
		if command[0] == "forward" {
			horizontal += value
			depth += aim * value
		} else if command[0] == "down" {
			aim += value
		} else if command[0] == "up" {
			aim -= value
		}
	}
	return horizontal * depth
}
