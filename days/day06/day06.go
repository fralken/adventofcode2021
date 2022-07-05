package day06

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	content := utils.ReadFile("./input/day06.txt")
	value := firstStar(content)
	fmt.Printf("day  6.1 - lanternfishes after 80 days: %d\n", value)
}

func SecondStar() {
	content := utils.ReadFile("./input/day06.txt")
	value := secondStar(content)
	fmt.Printf("day  6.2 - lanternfishes after 256 days: %d\n", value)
}

func firstStar(content string) int {
	fishes := utils.StringsToInts(strings.Split(content, ","))
	return live(fishes, 80)
}

func secondStar(content string) int {
	fishes := utils.StringsToInts(strings.Split(content, ","))
	return live(fishes, 256)
}

func live(fishes []int, days int) int {
	living := make([]int, 9)
	for _, fish := range fishes {
		living[fish]++
	}
	for d := 0; d < days; d++ {
		living[(d+7)%9] += living[d%9]
	}
	sum := 0
	for _, l := range living {
		sum += l
	}
	return sum
}
