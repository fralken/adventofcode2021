package day07

import (
	"aoc2021/utils"
	"strings"
)

func FirstStar() {
	utils.Star(7, 1, "fuel spent to align", firstStar)
}

func SecondStar() {
	utils.Star(7, 2, "fuel spent to align", secondStar)
}

func firstStar(content string) int {
	crabs := utils.StringsToInts(strings.Split(content, ","))
	return fuel(crabs, distance)
}

func secondStar(content string) int {
	crabs := utils.StringsToInts(strings.Split(content, ","))
	return fuel(crabs, distance2)
}

func fuel(crabs []int, distance func([]int, int) int) int {
	center := 0
	for _, crab := range crabs {
		center += crab
	}
	center /= len(crabs)
	currDist := distance(crabs, center)
	for backDist, i := distance(crabs, center-1), 2; backDist <= currDist; backDist, i = distance(crabs, center-i), i+1 {
		currDist = backDist
	}
	for forwDist, i := distance(crabs, center+1), 2; forwDist <= currDist; forwDist, i = distance(crabs, center+i), i+1 {
		currDist = forwDist
	}
	return currDist
}

func distance(crabs []int, center int) int {
	distance := 0
	for _, crab := range crabs {
		if crab > center {
			distance += (crab - center)
		} else {
			distance += (center - crab)
		}
	}
	return distance
}

func distance2(crabs []int, center int) int {
	sum := func(n int) int {
		return (n*n + n) / 2
	}
	distance := 0
	for _, crab := range crabs {
		if crab > center {
			distance += sum(crab - center)
		} else {
			distance += sum(center - crab)
		}
	}
	return distance
}
