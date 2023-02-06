package day22

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(22, 1, "cubes on in initialization area", firstStar)
}

func SecondStar() {
	utils.Star(22, 2, "cubes on", secondStar)
}

func firstStar(content string) int {
	cuboids := parseCuboids(content)
	var small []cuboid
	for _, c := range cuboids {
		if c.minx < -50 || c.minx > 50 {
			break
		}
		small = append(small, c)
	}
	small = *addCuboids(small)
	return countCubes(small)
}

func secondStar(content string) int {
	cuboids := parseCuboids(content)
	cuboids = *addCuboids(cuboids)
	return countCubes(cuboids)
}

type cuboid struct {
	action int
	minx int
	maxx int
	miny int
	maxy int
	minz int
	maxz int
}

func parseCuboids(content string) (cuboids []cuboid) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		var minx int
		var maxx int
		var miny int
		var maxy int
		var minz int
		var maxz int
		var command string
		var action int
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &command, &minx, &maxx, &miny, &maxy, &minz, &maxz)
		if command == "on" { action = 1 } else { action = 0 }
		cuboids = append(cuboids, cuboid{
			action,
			minx, maxx,
			miny, maxy,
			minz, maxz,
		})
	}
	return
}

func addCuboids(cuboids []cuboid) (cores *[]cuboid) {
	cores = &[]cuboid{}
	for _, c := range cuboids {
		var add []cuboid
		if c.action == 1 {
			add = append(add, c)
		}
		for _, core := range *cores {
			i := intersection(&c, &core)
			if i != nil {
				add = append(add, *i)
			}
		}
		*cores = append(*cores, add...)
	}
	return
}

func countCubes(cuboids []cuboid) int {
	count := 0
	for _, c := range cuboids {
		count += c.action * (c.maxx-c.minx+1) * (c.maxy-c.miny+1) * (c.maxz-c.minz+1) 
	}
	return count
}

func intersection(a *cuboid, b *cuboid) *cuboid {
	minx := max(a.minx, b.minx)
	maxx := min(a.maxx, b.maxx)
	miny := max(a.miny, b.miny)
	maxy := min(a.maxy, b.maxy)
	minz := max(a.minz, b.minz)
	maxz := min(a.maxz, b.maxz)
	if minx > maxx || miny > maxy || minz > maxz {
		return nil
	} else {
		return &cuboid{
			-b.action, // if cuboid is added, remove intersection, and vice versa
			minx, maxx,
			miny, maxy,
			minz, maxz,
		}
	}
}

func min(a, b int) int {
	if a < b { return a } else { return b }
}

func max(a, b int) int {
	if a > b { return a } else { return b }
}