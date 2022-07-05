package day02

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func FirstStar() {
	content := utils.ReadFile("./input/day02.txt")
	value := firstStar(content)
	fmt.Printf("day  2.1 - final horizontal position * final depth: %d\n", value)
}

func SecondStar() {
	content := utils.ReadFile("./input/day02.txt")
	value := secondStar(content)
	fmt.Printf("day  2.2 - final horizontal position * final depth: %d\n", value)
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
