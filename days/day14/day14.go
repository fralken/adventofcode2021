package day14

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	content := utils.ReadFile("./input/day14.txt")
	value := firstStar(content)
	fmt.Printf("day 14.1 - most minus least common elements after 10 steps: %d\n", value)
}

func SecondStar() {
	content := utils.ReadFile("./input/day14.txt")
	value := secondStar(content)
	fmt.Printf("day 14.2 - most minus least common elements after 40 steps: %d\n", value)
}

func firstStar(content string) int {
	polymer, pairs := extractPolymerPairs(content)
	return compute(polymer, pairs, 10)
}

func secondStar(content string) int {
	polymer, pairs := extractPolymerPairs(content)
	return compute(polymer, pairs, 40)
}

func extractPolymerPairs(content string) (polymer string, pairs map[string][2]string) {
	lines := strings.Split(content, "\n")
	polymer = lines[0]
	lines = lines[2:]
	pairs = make(map[string][2]string)
	for _, line := range lines {
		pair := strings.Split(line, " -> ")
		pairs[pair[0]] = [2]string{string(pair[0][0]) + pair[1], pair[1] + string(pair[0][1])}
	}
	return
}

func compute(polymer string, pairs map[string][2]string, steps int) int {
	count := make(map[string]int)
	for i := 0; i < len(polymer)-1; i++ {
		count[polymer[i:i+2]] = 1
	}
	for i := 0; i < steps; i++ {
		current := make(map[string]int)
		for k, v := range count {
			current[pairs[k][0]] += v
			current[pairs[k][1]] += v
		}
		count = current
	}
	elements := make(map[rune]int)
	for k, v := range count {
		elements[[]rune(k)[1]] += v
	}
	elements[[]rune(polymer)[0]]++
	min, max := findMinMax(elements)
	return max - min
}

func findMinMax(elements map[rune]int) (min int, max int) {
	for _, v := range elements {
		if min > v || min == 0 {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return
}
