package day13

import (
	"aoc2021/utils"
	"strings"
)

func FirstStar() {
	utils.Star(13, 1, "visible dots after 1 fold", firstStar)
}

func SecondStar() {
	utils.Star(13, 2, "code", secondStar)
}

func firstStar(content string) int {
	dots, folds := extractDotsFolds(content)
	visible := execFolds(dots, folds, 1)
	return len(visible)
}

func secondStar(content string) string {
	dots, folds := extractDotsFolds(content)
	visible := execFolds(dots, folds, len(folds))
	return extractCode(visible)
}

type coord [2]int

type instr struct {
	left     bool
	position int
}

func extractDotsFolds(content string) (dots []coord, folds []instr) {
	blocks := strings.Split(content, "\n\n")
	for _, dot := range strings.Split(blocks[0], "\n") {
		pos := strings.Split(dot, ",")
		dots = append(dots, coord{
			utils.StringToInt(pos[0]),
			utils.StringToInt(pos[1]),
		})
	}
	for _, fold := range strings.Split(blocks[1], "\n") {
		toks := strings.Split(fold, "=")
		folds = append(folds, instr{
			left:     toks[0] == "fold along x",
			position: utils.StringToInt(toks[1]),
		})
	}
	return
}

func execFolds(dots []coord, folds []instr, count int) map[coord]bool {
	visible := make(map[coord]bool)
	for _, dot := range dots {
		for i := 0; i < count; i++ {
			axis := 1
			if folds[i].left {
				axis = 0
			}
			if dot[axis] > folds[i].position {
				dot[axis] = 2*folds[i].position - dot[axis]
			}
		}
		visible[dot] = true
	}
	return visible
}

func maxXY(dots map[coord]bool) (maxX int, maxY int) {
	for dot := range dots {
		if dot[0] > maxX {
			maxX = dot[0]
		}
		if dot[1] > maxY {
			maxY = dot[1]
		}
	}
	maxX++
	maxY++
	return
}

func extractCode(dots map[coord]bool) string {
	maxX, maxY := maxXY(dots)
	grid := make([]string, maxY)
	for dot := range dots {
		if grid[dot[1]] == "" {
			grid[dot[1]] = strings.Repeat(" ", maxX)
		}
		grid[dot[1]] = replaceAt(grid[dot[1]], "#", dot[0])
	}
	return strings.Join(grid, "\n")
}

func replaceAt(s string, r string, i int) string {
	return s[:i] + r + s[i+1:]
}
