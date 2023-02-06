package day05

import (
	"aoc2021/utils"
	"strings"
)

func FirstStar() {
	utils.Star(5, 1, "number of points with overlaps", firstStar)
}

func SecondStar() {
	utils.Star(5, 2, "number of points with overlaps", secondStar)
}

func firstStar(content string) int {
	lines := strings.Split(content, "\n")
	vents := extractVents(lines)
	width, height := findGridSize(vents)
	grid := makeGrid(width, height)
	for _, vent := range vents {
		if vent.start[0] == vent.end[0] {
			x := vent.start[0]
			var inc int
			if vent.start[1] < vent.end[1] {
				inc = 1
			} else {
				inc = -1
			}
			for y := vent.start[1]; y != vent.end[1]; y += inc {
				grid[y][x]++
			}
			grid[vent.end[1]][vent.end[0]]++
		}
		if vent.start[1] == vent.end[1] {
			y := vent.start[1]
			var inc int
			if vent.start[0] < vent.end[0] {
				inc = 1
			} else {
				inc = -1
			}
			for x := vent.start[0]; x != vent.end[0]; x += inc {
				grid[y][x]++
			}
			grid[vent.end[1]][vent.end[0]]++
		}
	}
	return countOverlappingPoints(grid)
}

func secondStar(content string) int {
	lines := strings.Split(content, "\n")
	vents := extractVents(lines)
	width, height := findGridSize(vents)
	grid := makeGrid(width, height)
	for _, vent := range vents {
		var incx, incy int
		if vent.start[0] < vent.end[0] {
			incx = 1
		} else if vent.start[0] > vent.end[0] {
			incx = -1
		} else {
			incx = 0
		}
		if vent.start[1] < vent.end[1] {
			incy = 1
		} else if vent.start[1] > vent.end[1] {
			incy = -1
		} else {
			incy = 0
		}
		for x, y := vent.start[0], vent.start[1]; x != vent.end[0] || y != vent.end[1]; x, y = x+incx, y+incy {
			grid[y][x]++
		}
		grid[vent.end[1]][vent.end[0]]++
	}
	return countOverlappingPoints(grid)
}

type line struct {
	start []int
	end   []int
}

func extractVents(lines []string) (vents []line) {
	for _, l := range lines {
		points := strings.Split(l, " -> ")
		vent := line{
			start: utils.StringsToInts(strings.Split(points[0], ",")),
			end:   utils.StringsToInts(strings.Split(points[1], ",")),
		}
		vents = append(vents, vent)
	}
	return
}

func findGridSize(vents []line) (width int, height int) {
	for _, vent := range vents {
		if vent.start[0] > width {
			width = vent.start[0]
		}
		if vent.end[0] > width {
			width = vent.end[0]
		}
		if vent.start[1] > height {
			height = vent.start[0]
		}
		if vent.end[1] > height {
			height = vent.end[0]
		}
	}
	width++
	height++
	return
}

func makeGrid(width int, height int) [][]int {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	return grid
}

func countOverlappingPoints(grid [][]int) int {
	count := 0
	for y, row := range grid {
		for x := range row {
			if grid[y][x] > 1 {
				count++
			}
		}
	}
	return count
}
