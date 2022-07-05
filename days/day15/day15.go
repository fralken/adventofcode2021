package day15

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	content := utils.ReadFile("./input/day15.txt")
	value := firstStar(content)
	fmt.Printf("day 15.1 - lowest total risk: %d\n", value)
}

func SecondStar() {
	content := utils.ReadFile("./input/day15.txt")
	value := secondStar(content)
	fmt.Printf("day 15.2 - lowest total risk: %d\n", value)
}

func firstStar(content string) int {
	risk, h, w := createRiskFunc1(content)
	return findPath(risk, h, w)
}

func secondStar(content string) int {
	risk, h, w := createRiskFunc2(content)
	return findPath(risk, h, w)
}

type coord [2]int

func makeCave(content string) [][]int {
	lines := strings.Split(content, "\n")
	cave := make([][]int, len(lines))
	for i, line := range lines {
		cave[i] = utils.StringsToInts(strings.Split(line, ""))
	}
	return cave
}

func findPath(risk func(int, int) int, h int, w int) int {
	paths := make(map[coord]int)
	visited := make(map[coord]bool)
	visited[coord{0, 0}] = true
	neighbours := []coord{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	i, j, v := 0, 0, 0
	for i != h || j != w {
		for _, n := range neighbours {
			y, x := i+n[0], j+n[1]
			if y >= 0 && y <= h && x >= 0 && x <= w {
				k := coord{y, x}
				if !visited[k] {
					visited[k] = true
					w := v + risk(y, x)
					if paths[k] == 0 || paths[k] > w {
						paths[k] = w
					}
				}
			}
		}
		i, j, v = lowest(paths)
	}
	return v
}

func lowest(paths map[coord]int) (i int, j int, min int) {
	var c coord
	for k, v := range paths {
		if min == 0 || min > v {
			c = k
			min = v
		}
	}
	delete(paths, c)
	i, j = c[0], c[1]
	return
}

func createRiskFunc1(content string) (risk func(int, int) int, h int, w int) {
	cave := makeCave(content)
	risk = func(i int, j int) int {
		return cave[i][j]
	}
	h, w = len(cave)-1, len(cave[0])-1
	return
}

func createRiskFunc2(content string) (risk func(int, int) int, h int, w int) {
	cave := makeCave(content)
	ch, cw := len(cave), len(cave[0])
	risk = func(i int, j int) int {
		d := i/ch + j/cw
		i, j = i%ch, j%cw
		return (cave[i][j]+d-1)%9 + 1
	}
	h, w = (5*ch)-1, (5*cw)-1
	return
}
