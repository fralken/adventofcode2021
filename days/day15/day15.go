package day15

import (
	"aoc2021/utils"
	"container/heap"
	"strings"
)

func FirstStar() {
	utils.Star(15, 1, "lowest total risk", firstStar)
}

func SecondStar() {
	utils.Star(15, 2, "lowest total risk", secondStar)
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

type path struct {
	pos coord
	risk int
}

type paths []*path

func makeCave(content string) [][]int {
	lines := strings.Split(content, "\n")
	cave := make([][]int, len(lines))
	for i, line := range lines {
		cave[i] = utils.StringsToInts(strings.Split(line, ""))
	}
	return cave
}

func findPath(risk func(int, int) int, h int, w int) int {
	p := paths{}
	heap.Init(&p)
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
					heap.Push(&p, &path{ k, v + risk(y, x) })
				}
			}
		}
		path := heap.Pop(&p).(*path)
		i, j, v = path.pos[0], path.pos[1], path.risk
	}
	return v
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

func (p paths) Len() int { return len(p) }

func (p paths) Less(i, j int) bool {
	return p[i].risk < p[j].risk
}

func (p paths) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *paths) Push(x any) {
	item := x.(*path)
	*p = append(*p, item)
}

func (p *paths) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*p = old[:n-1]
	return item
}