package day03

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func FirstStar() {
	content := utils.ReadFile("./input/day03.txt")
	value := firstStar(content)
	fmt.Printf("day  3.1 - power consumption: %d\n", value)
}

func SecondStar() {
	content := utils.ReadFile("./input/day03.txt")
	value := secondStar(content)
	fmt.Printf("day  3.2 - life support rating: %d\n", value)
}

func firstStar(content string) int {
	lines := strings.Split(content, "\n")
	length := len(lines[0])
	ones, zeroes := countOnesAndZeroes(lines, length)
	gamma := 0
	epsilon := 0
	for i := 0; i < length; i++ {
		gamma *= 2
		epsilon *= 2
		if ones[i] > zeroes[i] {
			gamma++
		} else if ones[i] < zeroes[i] {
			epsilon++
		}
	}
	return gamma * epsilon
}

func secondStar(content string) int {
	lines := strings.Split(content, "\n")
	oxygen := extract(lines, mostCommon)
	co2 := extract(lines, leastCommon)
	return oxygen * co2
}

func countOnesAndZeroes(lines []string, length int) (ones []int, zeroes []int) {
	ones = make([]int, length)
	zeroes = make([]int, length)
	for _, line := range lines {
		for i := 0; i < length; i++ {
			if line[i] == '0' {
				zeroes[i]++
			} else if line[i] == '1' {
				ones[i]++
			}
		}
	}
	return
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeElements(s []string, digit byte, pos int) []string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i][pos] != digit {
			s = remove(s, i)
		}
	}
	return s
}

func mostCommon(one int, zero int) (digit byte) {
	digit = '0'
	if one >= zero {
		digit = '1'
	}
	return
}

func leastCommon(one int, zero int) (digit byte) {
	digit = '1'
	if zero <= one {
		digit = '0'
	}
	return
}

func extract(lines []string, f func(int, int) byte) int {
	s := make([]string, len(lines))
	copy(s, lines)
	length := len(s[0])
	pos := 0
	for len(s) > 1 && pos < length {
		ones, zeroes := countOnesAndZeroes(s, length)
		digit := f(ones[pos], zeroes[pos])
		s = removeElements(s, digit, pos)
		pos++
	}
	value, _ := strconv.ParseInt(s[0], 2, 32)
	return int(value)
}
