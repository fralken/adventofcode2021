package day09

import (
	"aoc2021/utils"
	"fmt"
	"sort"
	"strings"
)

func FirstStar() {
	content := utils.ReadFile("./input/day09.txt")
	value := firstStar(content)
	fmt.Printf("day  9.1 - sum of risk levels: %d\n", value)
}

func SecondStar() {
	content := utils.ReadFile("./input/day09.txt")
	value := secondStar(content)
	fmt.Printf("day  9.2 - product of sizes of three largest basins: %d\n", value)
}

func firstStar(content string) int {
	grid := makeGrid(content)
	count := 0
	for i, row := range grid {
		for j, c := range row {
			if isLower(grid, i, j) {
				count += c + 1
			}
		}
	}
	return count
}

func secondStar(content string) int {
	grid := makeGrid(content)
	var lowPoints []coord
	for i, row := range grid {
		for j := range row {
			if isLower(grid, i, j) {
				lowPoints = append(lowPoints, coord{i, j})
			}
		}
	}
	var basins []int
	for _, point := range lowPoints {
		basins = append(basins, findBasin(grid, point))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
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

func isLower(grid [][]int, i int, j int) bool {
	c := grid[i][j]
	row := grid[i]
	return !((i > 0 && grid[i-1][j] <= c) ||
		(i < len(grid)-1 && grid[i+1][j] <= c) ||
		(j > 0 && row[j-1] <= c) ||
		(j < len(row)-1 && row[j+1] <= c))
}

func findBasin(grid [][]int, point coord) int {
	basin := make(map[coord]bool)
	var visit []coord
	basin[point] = true
	visit = append(visit, point)
	for len(visit) > 0 {
		i, j := visit[0][0], visit[0][1]
		visit = visit[1:]
		if (i > 0 && grid[i-1][j] < 9 && !basin[coord{i - 1, j}]) {
			basin[coord{i - 1, j}] = true
			visit = append(visit, coord{i - 1, j})
		}
		if (i < len(grid)-1 && grid[i+1][j] < 9 && !basin[coord{i + 1, j}]) {
			basin[coord{i + 1, j}] = true
			visit = append(visit, coord{i + 1, j})
		}
		if (j > 0 && grid[i][j-1] < 9 && !basin[coord{i, j - 1}]) {
			basin[coord{i, j - 1}] = true
			visit = append(visit, coord{i, j - 1})
		}
		if (j < len(grid[i])-1 && grid[i][j+1] < 9 && !basin[coord{i, j + 1}]) {
			basin[coord{i, j + 1}] = true
			visit = append(visit, coord{i, j + 1})
		}
	}
	return len(basin)
}
