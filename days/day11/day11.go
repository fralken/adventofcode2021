package day11

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	content := utils.ReadFile("./input/day11.txt")
	value := firstStar(content)
	fmt.Printf("day 11.1 - flashes after 100 steps: %d\n", value)
}

func SecondStar() {
	content := utils.ReadFile("./input/day11.txt")
	value := secondStar(content)
	fmt.Printf("day 11.2 - step when flashing simultaneously: %d\n", value)
}

func firstStar(content string) int {
	grid := makeGrid(content)
	sum := 0
	for i := 0; i < 100; i++ {
		sum += step(grid)
	}
	return sum
}

func secondStar(content string) int {
	grid := makeGrid(content)
	when := 1
	gridSize := len(grid) * len(grid[0])
	for step(grid) != gridSize {
		when++
	}
	return when
}

func makeGrid(content string) [][]int {
	lines := strings.Split(content, "\n")
	grid := make([][]int, len(lines))
	for l, line := range lines {
		grid[l] = utils.StringsToInts(strings.Split(line, ""))
	}
	return grid
}

type coord [2]int

func step(grid [][]int) int {
	flashes := make(map[coord]bool)
	var visit []coord
	for i := range grid {
		for j := range grid[i] {
			grid[i][j]++
			if grid[i][j] == 10 {
				flashes[coord{i, j}] = true
				visit = append(visit, coord{i, j})
			}
		}
	}
	for len(visit) > 0 {
		center := visit[0]
		visit = visit[1:]
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				var y, x = center[0], center[1]
				if y+i >= 0 && x+j >= 0 && y+i < len(grid) && x+j < len(grid[y+i]) {
					grid[y+i][x+j]++
					if grid[y+i][x+j] == 10 {
						flashes[coord{y + i, x + j}] = true
						visit = append(visit, coord{y + i, x + j})
					}
				}
			}
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] >= 10 {
				grid[i][j] = 0
			}
		}
	}
	return len(flashes)
}
